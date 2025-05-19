// 汎用関数と汎用データ構造の組み合わせ
package main

import (
	"cmp"
	"fmt"
)

func main() {
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30}))
	fmt.Println(t2.Contains(Person{"Fred", 25}))

	t3 := NewTree(Person.Order)
	t3.Add(Person{"Bob", 30})
	t3.Add(Person{"Maria", 35})
	t3.Add(Person{"Bob", 50})
	fmt.Println(t3.Contains(Person{"Bob", 30}))
	fmt.Println(t3.Contains(Person{"Fred", 25}))
}

// OrderableFunc は、2つの値を比較して順序を判定する汎用関数型です。
// 戻り値は負の値、ゼロ、または正の値で、それぞれ t1 が t2 より小さい、等しい、大きいを示します。
type OrderableFunc[T any] func(t1, t2 T) int

// Tree は汎用的な二分木を表すデータ構造です。
// 型パラメータ T が任意の型を表し、ノードの値の型として使用されます。
// OrderableFunc はノード間の順序を判定する関数を提供します。
// ルートノードは内部的に root として管理されます。
type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

// Node は二分木におけるノードを表す構造体です。
type Node[T any] struct {
	val         T
	left, right *Node[T]
}

// NewTree は、OrderableFunc に基づいて要素の順序を管理する新しい Tree を生成します。
func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

// Add は Tree に値 v を追加します。
func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

// Contains は Tree 中に値 v が存在するかどうかを判定します。
func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

// Add は指定された値 v を二分木に追加します。OrderableFunc f を使用して値の順序を判定します。
func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

// Contains は指定された値 v が二分木に存在するかどうかを判定します。OrderableFunc f を使用して値の順序を判定します。
func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

// Person は、人の名前と年齢を表す構造体です。
type Person struct {
	Name string
	Age  int
}

// OrderPeople は、2つの Person 構造体を名前と年齢で比較し、順序を判定する関数です。
// 名前を優先して比較し、名前が一致する場合は年齢で比較します。
// 負の値は p1 が先、0 は等しい、正の値は p2 が先を示します。
func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

// Order は、2つの Person 構造体を比較し、順序を判定するメソッドです。
func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}
