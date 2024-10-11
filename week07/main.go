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
	var realScore float64
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	score = strings.TrimSpace(score)             // 줄바꿈, 띄어쓰기, 탭 등 제거 (python strip과 유사)
	realScore, _ = strconv.ParseFloat(score, 64) // 실수형 64비트 타입으로 형변환
	if realScore >= 90 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}
}
