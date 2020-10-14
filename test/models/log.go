package models

import "time"

type TraceLogger struct {
	ReqTime   time.Time
	ReqUri    string
	ReqMethod string
	Proto     string
	UserAgent string
	Referer   string
	Length int64
}

func NewLogger() *TraceLogger {
	return &TraceLogger{}
}