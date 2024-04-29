package api

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInsert   = errors.New("insert error")
	ErrUpdate   = errors.New("update error")
	ErrDelete   = errors.New("delete error")
	ErrList     = errors.New("list error")
	ErrDetail   = errors.New("detail error")
)
