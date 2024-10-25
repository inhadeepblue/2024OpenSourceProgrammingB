package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1 ~ 6
	//fmt.Println(answer)

	var win bool
	win = false
	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		//guess, err := strconv.ParseInt(input, 10, 32)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작은 수 입니다. LOW")
		} else { // answer < guess
			fmt.Println("입력하신 값은 정답보다 큰 수 입니다. HIGH")
		}
	}

	if win {
		fmt.Println("당신이 이겼습니다!")
	} else {
		fmt.Printf("당신이 졌습니다. 정답은 %d입니다\n", answer)
	}
}
