package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	length := 10
	list := make([]int, 0, length)
	for i := 1; i <= length; i++ {
		list = append(list, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	for {
		start := time.Now()
		num, err := promptNumber("pick a math facts number: ")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, i := range list {
			quizAdd(num, i)
		}
		fmt.Printf("quiz took %.2f seconds\n", time.Since(start).Seconds())
	}
}

func quizAdd(num1, num2 int) bool {
	question := fmt.Sprintf("%d + %d = ", num1, num2)
	expectedNumber := num1 + num2

	num, err := promptNumber(question)
	if err != nil {
		fmt.Println("wrong")
		return false
	}

	if num == expectedNumber {
		fmt.Println("correct")
		return true
	}

	fmt.Printf("wrong - you typed %d but the answer is %d\n", num, expectedNumber)
	return false
}

func promptNumber(promptText string) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", promptText)
	text, _, err := reader.ReadLine()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(text))
}
