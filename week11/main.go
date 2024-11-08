package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// input
// 10
// 19

// output
// 11 13 17 19

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3
		for i*i <= n {
			if n%i == 0 {
				return false
			}
			//fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}

func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a)
	number, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func main() {
	fmt.Print("첫 번째 정수(시작 값) 입력 : ")
	n1 := getInteger()
	fmt.Print("두 번째 정수(끝 값) 입력 : ")
	n2 := getInteger()

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
