package api

import "errors"

var (
	ErrDishesNotFound = errors.New("dishes not found")
	ErrDishesInsert   = errors.New("dishes insert error")
	ErrDishesUpdate   = errors.New("dishes update error")
	ErrDishesDelete   = errors.New("dishes delete error")
	ErrDishesList     = errors.New("dishes list error")
	ErrDishesDetail   = errors.New("dishes detail error")
)
