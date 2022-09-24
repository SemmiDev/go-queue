package queue

import "container/list"

type LinkListQueue[T any] struct {
	items *list.List
}

func newLinkListQueue[T any]() *LinkListQueue[T] {
	return &LinkListQueue[T]{
		items: list.New(),
	}
}

func (q *LinkListQueue[T]) Enqueue(item T) {
	q.items.PushBack(item)
}

func (q *LinkListQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var emptyItem T
		return emptyItem, ErrEmptyQueue
	}

	item := q.items.Front()
	q.items.Remove(item)
	return item.Value.(T), nil
}

func (q *LinkListQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var emptyItem T
		return emptyItem, ErrEmptyQueue
	}

	return q.items.Front().Value.(T), nil
}

func (q *LinkListQueue[T]) Size() int {
	return q.items.Len()
}

func (q *LinkListQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *LinkListQueue[T]) GetItems() []T {
	items := make([]T, 0)
	for item := q.items.Front(); item != nil; item = item.Next() {
		items = append(items, item.Value.(T))
	}

	return items
}
