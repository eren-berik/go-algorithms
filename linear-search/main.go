package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LinearSearch(arr []int, num int) int {
	for index := range arr {
		if arr[index] == num {
			return index
		}
	}
	return -1
}

func convertStringToSlice(str string) ([]int, error) {
	strNumber := strings.Fields((str))

	var arr []int
	for _, strNum := range strNumber {
		num := 0
		_, err := fmt.Sscanf(strNum, "%d", &num)
		if err != nil {
			fmt.Println("Error while converting input: ", err)
			return nil, err
		}
		arr = append(arr, num)
	}
	return arr, nil
}

func main() {
	var inputString string
	fmt.Print("Enter a slice: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = scanner.Text()

	arr, _ := convertStringToSlice(inputString)

	var num int
	fmt.Print("Enter a number you want to search for: ")
	fmt.Scan(&num)

	fmt.Printf("Index of the %d is [%d]\n", num, LinearSearch(arr, num))
}
