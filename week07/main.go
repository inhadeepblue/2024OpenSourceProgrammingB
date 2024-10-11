package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)
}
