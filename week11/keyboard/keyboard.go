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

func GetFloat() (float64, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	a = strings.TrimSpace(a)
	number, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
