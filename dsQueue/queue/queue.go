package queue

import "errors"

type Queue struct {
	size int
	data []string
}

func New(size int) *Queue {
	return &Queue{
		size: size,
		data: make([]string, 0),
	}
}

func (q *Queue) EnQueue(ele string) error { //adds element in rear
	l := q.SizeOfQueue()
	if l < q.size {
		q.data = append(q.data, ele)
		return nil
	}
	return errors.New("queue is full")
}

func (q *Queue) DeQueue() (string, error) { //deletes element in front
	l := q.SizeOfQueue()
	if l == 0 {
		return "", errors.New("queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}

func (q *Queue) Front() (string, error) { //returns 1st element
	if q.SizeOfQueue() == 0 {
		return "", errors.New("queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) SizeOfQueue() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.SizeOfQueue() == 0
}
