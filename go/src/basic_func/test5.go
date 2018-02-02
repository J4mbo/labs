/*
책: Go언어 웹프로그래밍 철저 입문
switch문 테스트
*/
package main

import "fmt"

func main(){
	x := 17
	switch {
	case x > 0:
		fmt.Println(x, "는 양수입니다.")
		fallthrough //조건을 만족했더라도 switch문을 종료하지 않고 다음 case로 넘어간다.
	case x > 5:
		fmt.Println(x, "는 5보다 큰 수 입니다.")
		//fallthrough를 사용하지 않았으므로 조건을 만족했을 때 위의 문장을 print하고 switch문을 종료한다.
	case x > 10:
		fmt.Println(x, "는 10보다 큰 수 입니다.")// 이 구문은 x가 17인 경우에서 수행되지 않는다.
	}

}