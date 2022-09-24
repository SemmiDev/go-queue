package queue

type SliceQueue[T any] []T

func newSliceQueue[T any]() *SliceQueue[T] {
	q := SliceQueue[T](make([]T, 0))
	return &q
}

func (q *SliceQueue[T]) GetItems() []T {
	return *q
}

func (q *SliceQueue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var emptyItem T
		return emptyItem, ErrEmptyQueue
	}

	item := (*q)[0]
	*q = (*q)[1:]
	return item, nil
}

func (q *SliceQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var emptyItem T
		return emptyItem, ErrEmptyQueue
	}

	return (*q)[0], nil
}

func (q *SliceQueue[T]) Size() int {
	return len(*q)
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}
