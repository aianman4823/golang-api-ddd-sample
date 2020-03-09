package model

import "time"

// Book : Book を表すドメインモデル
type Book struct {
	Id       int64
	Title    string
	Author   string
	IssuedAt time.Time
}
