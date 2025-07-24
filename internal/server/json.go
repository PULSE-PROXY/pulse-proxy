package server

import (
	"time"
)

type Metadata struct {
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

type Response struct {
	Status   int         `json:"status"`
	Detail   string      `json:"detail"`
	Content  interface{} `json:"content"`
	Error    interface{} `json:"error"`
	Metadata Metadata    `json:"metadata"`
}
