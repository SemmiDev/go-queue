package queue

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	var tests = []struct {
		queueType    Type
		input        []any
		expectedSize int
	}{
		{Structures, []any{1, 2, 3, 4, 5}, 5},
		{Structures, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{Structures, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15},

		{Slices, []any{1, 2, 3, 4, 5}, 5},
		{Slices, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{Slices, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15},

		{LinkList, []any{1, 2, 3, 4, 5}, 5},
		{LinkList, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{LinkList, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15},
	}

	for _, test := range tests {
		q := NewQueueFactory(test.queueType)
		for _, item := range test.input {
			q.Enqueue(item)
		}

		if q.Size() != test.expectedSize {
			t.Errorf("expected size %d, got %d", test.expectedSize, q.Size())
		}
	}
}

func TestQueue_Dequeue(t *testing.T) {
	var tests = []struct {
		queueType    Type
		input        []any
		dequeueCount int
		expectedErr  []error
	}{
		{Structures, []any{1, 2, 3, 4, 5}, 5, []error{nil, nil, nil, nil, nil}},
		{Structures, []any{1, 2, 3, 4, 5}, 10, []error{nil, nil, nil, nil, nil, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue}},
		{Structures, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []error{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},

		{Slices, []any{1, 2, 3, 4, 5}, 5, []error{nil, nil, nil, nil, nil}},
		{Slices, []any{1, 2, 3, 4, 5}, 10, []error{nil, nil, nil, nil, nil, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue}},
		{Slices, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []error{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},

		{LinkList, []any{1, 2, 3, 4, 5}, 5, []error{nil, nil, nil, nil, nil}},
		{LinkList, []any{1, 2, 3, 4, 5}, 10, []error{nil, nil, nil, nil, nil, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue, ErrEmptyQueue}},
		{LinkList, []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []error{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
	}

	for _, test := range tests {
		q := NewQueueFactory(test.queueType)
		for _, item := range test.input {
			q.Enqueue(item)
		}

		for i := 0; i < test.dequeueCount; i++ {
			_, err := q.Dequeue()
			if err != test.expectedErr[i] {
				t.Errorf("expected error %v, got %v", test.expectedErr[i], err)
			}
		}
	}
}
