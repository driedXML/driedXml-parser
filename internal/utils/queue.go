package utils

import (
	"errors"
)

func NewQueue() Queue {
	return &queue{}
}

type Queue interface {
	IsEmpty() bool
	Size() int
	Peek() (any, error)
	Enqueue(item any)
	Dequeue() (any, error)
}

type queue struct {
	items []any
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Size() int {
	return len(q.items)
}

func (q *queue) Peek() (any, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.items[0], nil
}

func (q *queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}
