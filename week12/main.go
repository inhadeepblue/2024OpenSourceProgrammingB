package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[0] = time.Unix(0, 0)
	// dates[2] = time.Unix(1708012346, 0)
	// fmt.Println(dates[0], dates[1], dates[2]) // 1970-01-01 00:00:00, zero value, 2024-02-16 ...

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2]) // 1970-01-01 00:00:00, +1, 2024-02-16 ...

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2]) // 1970-01-01 00:00:00, +1, 2024-02-16 ...

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1708012346, 0), // need comma
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Println(dates)         // array
	// fmt.Printf("%#v\n", dates) // array literal

	//for i := 0; i <= 7; i++ {
	// for i := 0; i < len(dates); i++ {
	// 	fmt.Println(i, dates[i])
	// }

	for i, date := range dates { // like python for in, SAFE!
		fmt.Println(i, date)
	}

	// for _, date := range dates { // like python for in, SAFE!
	// 	fmt.Println(date)
	// }
}

/*
import random
numbers = [random.randint(1, 100) for _ in range(10)]
print(numbers)
#for i in range(len(numbers)):
#for i in range(11):  # Runtime error
#	print(numbers[i], end=' ')
for number in numbers:  # safe
	print(number, end=' ')
*/
