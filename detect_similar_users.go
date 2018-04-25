package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var s = bufio.NewScanner(os.Stdin)

type UserBookPair [2]int

func (ubp UserBookPair) userNum() int {
	return ubp[0]
}

func (ubp UserBookPair) bookNum() int {
	return ubp[1]
}

func nextLine() string {
	s.Scan()
	return s.Text()
}

func main() {
	inputValue := strings.Split(nextLine(), " ")
	ubp := UserBookPair{}

	// get the number of user and book
	for i, v := range inputValue {
		if num, err := strconv.Atoi(v); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			ubp[i] = num
		}
	}

	data := make([][]float64, ubp.userNum())
	result := make([][]float64, ubp.userNum()-1)

	// get the all evaluation
	for i := 0; i < ubp.userNum(); i += 1 {
		inputValue := strings.Split(nextLine(), " ")
		for _, v := range inputValue {
			if num, err := strconv.ParseFloat(v, 64); err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				data[i] = append(data[i], num)
			}
		}
	}

	// calculate degree of similarity
LOOP:
	for i, v := range data {
		var sum float64
		for j, w := range v {
			if i == 0 {
				continue LOOP
			}

			if data[0][j] == -1 {
				continue
			} else if w == -1 {
				continue
			} else {
				sum += math.Pow(math.Abs(data[0][j]-w), 2)
			}
		}
		result[i-1] = []float64{float64(i + 1), 1 / (math.Sqrt(sum) + 1)}
	}

	// sort
	sort.Slice(result, func(i, j int) bool { return result[i][1] > result[j][1] })

	// output results
	for _, v := range result {
		fmt.Println(v[0], v[1])
	}
}
