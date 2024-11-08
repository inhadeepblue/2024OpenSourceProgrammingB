package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetInteger() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	a = strings.TrimSpace(a)
	number, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return number, nil
}
