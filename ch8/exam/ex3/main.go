package main

import "fmt"

// Write a generic singly linked list data type. Each element can hold a
// comparable value and has a pointer to the next element in the list.
// The methods to implement are:
//
//    // adds a new element to the end of the linked list
//    Add(T)
//    // adds an element at the specified position in the linked list
//    Insert(T, int)
//    // returns the position of the supplied value, -1 if it's not present
//    Index (T) int

// Node は単方向リストの各要素を表す構造体です。
// ジェネリック型Tを使用しており、Tはcomparableである必要があります。
// Valueはノードが保持する値を表し、Nextは次のノードへのポインタを指します。
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// List は単方向リストを表す構造体です。
type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

// Add は単方向リストの末尾に新しい要素を追加します。
func (l *List[T]) Add(t T) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

// Insert は単方向リストの指定された位置に新しい要素を追加します。
func (l *List[T]) Insert(t T, pos int) {
	n := &Node[T]{
		Value: t,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if pos <= 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	curNode := l.Head
	for i := 1; i < pos; i++ {
		if curNode.Next == nil {
			curNode.Next = n
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}
	n.Next = curNode.Next
	curNode.Next = n
	if l.Tail == curNode {
		l.Tail = n
	}
	return
}

// Index は指定された値を持つ要素の位置を取得します。該当する要素がない場合は -1 を返します。
func (l *List[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}
	return -1
}

func main() {
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
