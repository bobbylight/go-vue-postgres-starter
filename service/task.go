package main

import "time"

type Task struct {
    Id string `json:"id"`
    Label       string `json:"label"`
    Description string `json:"desc"`
    StatusCd    string `json:"status"`
    Index       *int `json:"index"`
    CreatedDttm *string `json:"createdAt"`
    UpdatedAtDttm *string `json:"updatedAt"`
    CompletedDttm *string `json:"completedAt"`
}

func (t *Task) BeforeCreate() (err error) {
    now := time.Now().Format(time.RFC3339)
    t.CreatedDttm = &now
    return
}

func (t *Task) BeforeUpdate() (err error) {
    now := time.Now().Format(time.RFC3339)
    t.UpdatedAtDttm = &now
    return
}
