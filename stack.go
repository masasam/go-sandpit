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

func newList() *List {
	lst := new(List)
	lst.top = new(Cell)
	return lst
}

func (cp *Cell) nthCell(n int) *Cell {
	i := -1
	for cp != nil {
		if i == n {
			return cp
		}
		i++
		cp = cp.next
	}
	return nil
}

func (lst *List) nth(n int) (int, bool) {
	cp := lst.top.nthCell(n)
	if cp == nil {
		return 0, false
	}
	return cp.item, true
}

func (lst *List) insertNth(n, x int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil {
		return false
	}
	cp.next = newCell(x, cp.next)
	return true
}

func (lst *List) deleteNth(n int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil || cp.next == nil {
		return false
	}
	cp.next = cp.next.next
	return true
}

func (lst *List) isEmpty() bool {
	return lst.top.next == nil
}

func (lst *List) printList() {
	cp := lst.top.next
	for ; cp != nil; cp = cp.next {
		fmt.Print(cp.item, " ")
	}
	fmt.Println("")
}

type Stack struct {
	content *List
}

func newStack() *Stack {
	st := new(Stack)
	st.content = newList()
	return st
}

func (st *Stack) push(x int) {
	st.content.insertNth(0, x)
}

func (st *Stack) pop() (int, bool) {
	x, ok := st.content.nth(0)
	if ok {
		st.content.deleteNth(0)
	}
	return x, ok
}

func (st *Stack) top() (int, bool) {
	return st.content.nth(0)
}

func (st *Stack) isEmpty() bool {
	return st.content.isEmpty()
}

func main() {

	st := newStack()
	for i := 0; i < 8; i++ {
		st.push(i)
		fmt.Println(st.top())
	}
	for !st.isEmpty() {
		fmt.Println(st.pop())
	}
}
