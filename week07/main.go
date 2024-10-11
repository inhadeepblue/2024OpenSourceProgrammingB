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
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	score = strings.TrimSpace(score)                // 줄바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
	realScore, _ := strconv.ParseInt(score, 10, 32) // 정수형 32비트 타입으로 형변환, 10진수 입력

	var grade string
	if realScore >= 90 {
		grade = "A"
	} else {
		grade = "BCDF"
	}
	fmt.Printf("%d점은 %s등급 입니다\n", realScore, grade)
}
