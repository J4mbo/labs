/*
책: Go언어 웹프로그래밍 철저 입문
if문 테스트
*/
package main

import "fmt"

func main() {
	if x, y:= -2, -4; x<0 && y<0 {
		fmt.Println(x, y, "모두 음수입니다.")
	} else if x>0 && y>0{
		fmt.Println(x, y, "모두 양수입니다.")
	} else {
		fmt.Println(x, y, "중 하나는 음수, 나머지 하나는 양수입니다.")
	}
}