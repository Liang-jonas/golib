package db

import (
	"time"
)

type Model struct {
	ID       int64      `json:"id"`
	CreateAt *time.Time `json:"create-at"`
	UpdateAt *time.Time `json:"update-at"`
	DeleteAt *time.Time `json:"delete-at"`
}
