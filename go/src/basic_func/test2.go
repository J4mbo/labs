/*
책: Go언어 웹프로그래밍 철저 입문
열거형 테스트1
*/
package main

import "fmt"

type Color int
const (
	RED Color = iota //0
	ORANGE 			 //1
	YELLOW			 //2
	GREEN            //3
)

type ByteSize int64
const (
	_ = iota                       // ignore
	KB ByteSize = 1 << (10 * iota) // 1 << (10 * 1)
	MB							   // 1 << (10 * 2)
	GB							   // 1 << (10 * 3)
	TB
	PB
	EB
)

const (
	DEFAULT_RATE = 5 + 0.3 * iota
	GREEN_RATE
	SILVER_RATE
	GOLD_REATE
)

func main() {
	type Rect struct {
		width  int //width
		height int //height
	}

	r := Rect{1, 2}
	fmt.Println(r.width*2 + r.height*2)
}
