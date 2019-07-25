package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world.")

	var x int
	//var y = 30

	x = 10
	//z := 20
	//defer fmt.Println(add(x, 22))

	for i := 0; i < 10; i++ {

		if w := add(x, i); w%3 == 0 {
			fmt.Println(w)
		}
	}
	fmt.Println("")

	var p *int
	i := 42
	p = &i

	fmt.Printf("Pは%v\r\n", *p)
	i = 21
	fmt.Println("iは", *p)

	pp := Point{1, 2}

	v := &Point{1, 2}

	fmt.Printf("ppは%v %d\r\n", pp.x, v.y)
	fmt.Printf("ppは%v %d\r\n", (*v).x, (*v).y)

	// 配列
	var 配列 [3]int
	配列[0] = 1
	配列[1] = 3
	配列[2] = 5
	fmt.Printf("配列は%v \r\n", 配列[0])

	//スライス
	var s []int = 配列[0:3]

	for cnt := 0; cnt < 3; cnt++ {
		fmt.Printf("要素%vは%d\r\n", cnt, s[cnt])
	}
	fmt.Printf("要素数は%v\r\n", len(s))
	fmt.Printf("容量は%v\r\n", cap(s))

	sl := make([]int, 5, 6)
	sl = append(sl, 1, 2, 3, 4, 5, 6)
	for j := 0; j < len(sl); j++ {
		fmt.Printf("要素%vは%d\r\n", j, sl[j])
	}
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	fmt.Println(sl)

	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3)

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3)

	// Extend its length.
	s3 = s3[:5]
	printSlice(s3)

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3)
	//fmt.Println(x)

	//map
	//map1 := make(map[string]int)

	var m = map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Printf("mapの１つ目の要素は%d\r\n", m["one"])
	fmt.Printf("mapの２つ目の要素は%d\r\n", m["two"])

	elem, ok := m["one"]

	fmt.Println("mapの２つ目の要素は\r\n", ok, elem)

	pp.アップ()
	fmt.Println("pp", pp.Add())
}

func (p *Point) アップ() {
	p.x += 1
}

func (p Point) Add() int {
	return p.x + p.y
}

/* 構造体 */
type Point struct {
	x int
	y int
}

func add(x int, y int) int {
	return x + y
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
