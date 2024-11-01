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

	//counts := 0
	var isPrime bool = true // int -> bool, counts -> isPrime, memory
	if n <= 1 {
		//counts = -1
		isPrime = false // 가독성
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				//counts = counts + 1
				isPrime = false // +연산자 제거
			}
			i++
		}
	}

	//if counts == 0 {
	if isPrime { // == 비교연산자 제거
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다!", n)
	}
}
