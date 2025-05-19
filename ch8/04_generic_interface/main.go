package main

import (
	"fmt"
	"math"
)

// Pair は同じ型の2つの値を保持する汎用コンテナです。
// fmt.Stringerを実装している型だけを指定する
type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

// Differ は任意の型 T における差を計算するためのインターフェースです。
// Diff メソッドは与えられた T 型の値との間の差を float64 型で返します。
// fmt.Stringer インターフェースを実装しており、文字列表現を取得可能です。
type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

// FindCloser は2つのPairのうち、Val1とVal2の差がより小さい方を返す汎用関数です。
func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	}
	return pair2
}

// Differインターフェイスを満たす型をいくつか定義

// Point2D は2次元平面上の点を表す構造体です。
// 各点は整数座標 X と Y を持ちます。
type Point2D struct {
	X, Y int
}

// String メソッドは Point2D 構造体を文字列形式で返します。
// 座標情報は "{X,Y}" の形式でフォーマットされます。
func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

// Diff メソッドは、現在の Point2D と指定された Point2D とのユークリッド距離を計算して返します。
func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

// Point3D は3次元空間上の点を表す構造体です。
type Point3D struct {
	X, Y, Z int
}

// String メソッドは Point3D 構造体を文字列形式で返します。
// 座標情報は "{X,Y,Z}" の形式でフォーマットされます。
func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}

// Diff メソッドは、現在の Point3D と指定された Point3D とのユークリッド距離を計算して返します。
func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func main() {
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer2 := FindCloser(pair3Da, pair3Db)
	fmt.Println(closer2)
}
