package main

import "fmt"

type Cell struct {
	item int
	next *Cell
}

type List struct {
	top *Cell
}

func newCell(x int, cp *Cell) *Cell {
	newcp := new(Cell)
	newcp.item, newcp.next = x, cp
	return newcp
}

type Queue struct {
	size        int
	front, rear *Cell
}

func newQueue() *Queue {
	return new(Queue)
}

func (q *Queue) enqueue(x int) {
	cp := newCell(x, nil)
	if q.size == 0 {
		q.front = cp
		q.rear = cp
	} else {
		q.rear.next = cp
		q.rear = cp
	}
	q.size++
}

func (q *Queue) dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	} else {
		x := q.front.item
		q.front = q.front.next
		q.size--
		if q.size == 0 {
			q.rear = nil
		}
		return x, true
	}
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func main() {
	que := newQueue()
	for i := 0; i < 8; i++ {
		que.enqueue(i)
		fmt.Println(i)
	}
	for !que.isEmpty() {
		fmt.Println(que.dequeue())
	}
}
