package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number)
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				isPrime = false
				break // 1과 자기자신을 제외한 첫 번째 약수가 발견 되면 반복문 종료
			}
			fmt.Printf("%d ", i) // 반복 횟수 확인용 코드
			i++
		}
	}

	if isPrime {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!", n)
	}
}
