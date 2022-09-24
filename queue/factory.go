package queue

import "errors"

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type IQueue[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	GetItems() []T
}

type Type int

const (
	Slices Type = iota
	LinkList
)

func NewQueueFactory(queueType Type) IQueue[any] {
	switch queueType {
	case Slices:
		return newSliceQueue[any]()
	case LinkList:
		return newLinkListQueue[any]()
	default:
		return nil
	}
}
