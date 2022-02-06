package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var number int
	tempNumbers := make([]int, 0, 0)
	number = 26712
	for number > 0 {
		tempNumbers = append(tempNumbers, number%10)
		number = number / 10
	}
	numbers := make([]int, len(tempNumbers), cap(tempNumbers))
	lentempNumbers := len(tempNumbers)
	for i, v := range tempNumbers {
		numbers[lentempNumbers-(i+1)] = v
	}
	lenNumbers := len(numbers)
	signs := make([]int64, lenNumbers-1, cap(numbers)-1)
	signsResult := make([]int64, lenNumbers-1, cap(numbers)-1)
	lenSigns := len(signs)
	isResult := false
	for i := 0; i < powInt(2, lenSigns); i++ {

		runs := []rune(strconv.FormatInt(int64(i), 2))
		offset := lenSigns - len(runs)
		for j := offset; j < lenSigns; j++ {
			signs[j], _ = strconv.ParseInt(string(runs[j-offset]), 10, 0)
		}
		sum := numbers[0]
		for j := 0; j < lenNumbers-1; j++ {
			sign := -1
			if signs[j] == 0 {
				sign = 1
			}
			sum = sum + numbers[j+1]*sign
		}
		if sum == 0 {
			isResult = true
			copy(signsResult, signs)
		}
	}
	if isResult {
		for ind, val := range numbers {
			if ind < len(numbers)-1 {
				sign := "-"
				if signsResult[ind] == 0 {
					sign = "+"
				}
				fmt.Print(val, sign)
			} else {
				fmt.Print(val)
			}
		}
		fmt.Print("=0")
	} else {
		fmt.Println("not possible")
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
