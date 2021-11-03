package jobModule

import (
	"FriendlyAlmond_backend/pkg/model"
	"time"
)

type RespJob struct {
	model.JSONResult
	Data Job `json:"data"`
}

type Job struct {
	JobId   int32     `json:"job_id" gorm:"not null;autoIncrement;primaryKey"`
	OrderId int32     `json:"order_id" gorm:"not null"`
	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time `gorm:"autoUpdateTime"`
}

func (Job) TableName() string {
	return "job_order"
}

type RespTask struct {
	model.JSONResult
	Data []Task `json:"data"`
}

type Task struct {
	TaskId      int32     `json:"task_id" gorm:"not null;autoIncrement;primaryKey"`
	JobId       int32     `json:"job_id" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	SectionId   int32     `json:"section_id" gorm:"not null"`
	ComponentId int32     `json:"component_id" gorm:"not null"`
	StartDate   string    `json:"start_date" gorm:"not null"`
	DueDate     string    `json:"due_date" gorm:"not null"`
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
	StaffId     int32     `json:"staff_id" gorm:"not null"`
}

func (Task) TableName() string {
	return "task"
}
