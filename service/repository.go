package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository() *Repository {

    db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
    if err != nil {
        panic(err)
    }
    //defer db.Close()

    db.Exec("set search_path to demo_schema")

    db.AutoMigrate(&Widget{})

    return &Repository{
        db: db,
    }
}

func (r Repository) getWidgetById(id string) *Widget {

    var widget Widget
    if r.db.Find(&widget, id).RecordNotFound() {
        return nil
    }

    return &widget
}

func (r Repository) getWidgets() []Widget {

    var widgets[] Widget
    r.db.Find(&widgets)

    return widgets
}
