package models

type BookOrder struct {
	Side string `json:"side"`
	Price float64 `json:"price"`
}

// A Tree is a binary tree values.
type Tree struct {
	Left  *Tree
	Value BookOrder
	Right *Tree
}


