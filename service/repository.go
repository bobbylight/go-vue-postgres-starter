package main

import (
    "errors"
    "fmt"
    "github.com/google/uuid"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repository struct {
    db *gorm.DB
}

const order = "index asc, label asc"
const reverseOrder = "index desc, label desc"

func NewRepository() *Repository {

    db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        panic(err)
    }
    //defer db.Close()

    db.Exec("set search_path to dwitd")

    db.AutoMigrate(&Task{})

    return &Repository{
        db: db,
    }
}

func (r Repository) createTask(label string, desc string) *Task {

    // TODO: Don't crash the entire app on failure
    uuid, err := uuid.NewRandom()
    if err != nil {
        panic("Unable to generate UUID: " + err.Error())
    }

    taskCount := int(r.getTaskCount())

    task := Task{
        Id:          uuid.String(),
        Label:       label,
        Description: desc,
        Index:       &taskCount,
        StatusCd:    "notStarted",
    }

    r.db.Create(&task)
    return &task
}

func (r Repository) deleteTaskById(id string) *[]Task {

    var task Task
    if r.db.Where("id = ?", id).Delete(&task).RecordNotFound() {
        return nil
    }

    r.db.Exec("update tasks set index = index + 1 where index > ?", task.Index)

    remainingTasks := r.getTasks(0, 10000)
    return &remainingTasks
}

func (r Repository) getTaskById(id string) *Task {

    var task Task
    if r.db.Find(&task, "id = ?", id).RecordNotFound() {
        return nil
    }
    return &task
}

func (r Repository) getTaskCount() uint32 {
    var count uint32
    r.db.Model(&Task{}).Count(&count)
    return count
}

func (r Repository) getTasks(offset uint32, limit uint32) []Task {

    var tasks []Task
    r.db.Order(order).Offset(offset).Limit(limit).Find(&tasks)

    return tasks
}

func (r Repository) reorderTask(taskId string, params *TaskReorderParams) (*[]Task, error) {

    // Get the server-side value for the task by ID
    task := r.getTaskById(taskId)
    if task == nil {
        return nil, fmt.Errorf("no such task: %v", taskId)
    }

    index := task.Index
    var swapWith Task

    if params.Forward {

        // Find the lowest index greater than this one
        if r.db.Where("index > ?", index).Order(order).
            First(&swapWith).RecordNotFound() {
                return nil, errors.New("no task comes after this one")
        }
    } else {

        // Find the lowest index greater than this one
        if r.db.Where("index < ?", index).Order(reverseOrder).
            First(&swapWith).RecordNotFound() {
                return nil, errors.New("no task comes before this one")
        }
    }

    if &swapWith != nil {

        tx := r.db.Begin()
        defer func() {
            if r := recover(); r != nil {
                tx.Rollback()
            }
        }()

        if err := tx.Error; err != nil {
            return nil, errors.New("error starting transaction")
        }

        if err := tx.Exec("update tasks set index = ? where id = ? ",
                index, swapWith.Id).Error; err != nil {
            tx.Rollback()
            return nil, fmt.Errorf("error updating first task: %v", err)
        }

        if err := tx.Exec("update tasks set index = ? where id = ? ",
                swapWith.Index, task.Id).Error; err != nil {
            tx.Rollback()
            return nil, fmt.Errorf("error updating second task: %v", err)
        }

        if err := tx.Commit().Error; err != nil {
            return nil, fmt.Errorf("error performing commit: %v", err)
        }
    }

    newTasks := r.getTasks(0, 10000)
    return &newTasks, nil
}

func (r Repository) updateTask(task *Task) *Task {

    if r.db.Save(&task).RecordNotFound() {
        return nil
    }

    // Re-query for updated fields such as last-modified
    if r.db.Find(&task, "id = ?", task.Id).RecordNotFound() {
        return nil
    }
    return task
}
