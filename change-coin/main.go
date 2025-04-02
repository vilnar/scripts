package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getExPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}

func getInputFilePath() string {
	return filepath.Join(getExPath(), "input.txt")
}

func readInputFile() ([]int, error) {
	fpath := getInputFilePath()
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	if len(nums) < 1 {
		return nil, fmt.Errorf("Помилка: не знайдені числа в файлі %s \n", fpath)
	}

	return nums, nil
}

func coinChange(coins []int, sum int) []int {
	result := make([]int, sum+1)
	cache := make([]int, sum+1)
	cache[0] = 0

	for i := 1; i <= sum; i++ {
		cache[i] = math.MaxInt32

		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				cacheValue := cache[i-coins[j]] + 1
				if cacheValue < cache[i] {
					cache[i] = cacheValue
					result[i] = j + 1
				}
			}
		}
	}

	return result
}

func printChange(coins []int, sum int) {
	res := coinChange(coins, sum)
	// fmt.Printf("res!!! \n%#v\n", res)

	csum := 0
	sumCount := sum
	isApproximateSum := false
	for i := 0; sumCount > 0; i++ {
		index := res[sumCount] - 1
		if index < 0 {
			sumCount = sumCount - 1
			isApproximateSum = true
			continue
		}
		csum += coins[index]
		fmt.Printf("%d\n", coins[index])
		sumCount = sumCount - coins[index]
	}

	if isApproximateSum {
		fmt.Printf("\nДана сума %d не може бути підібрана з тих комбінацій які є, тому знайдена приблизна сума %d\n", sum, csum)
	}
}

func main() {
	coins, err := readInputFile()
	if err != nil {
		log.Printf("%s\n", err)
		fmt.Scanln()
		os.Exit(1)
	}

	fmt.Printf("\nПриклади кілометрів взяті з файлу %s\n%v\n\n", getInputFilePath(), coins)

	fmt.Printf("Введіть потрібну суму кілометрів:\n")
	var sum int
	fmt.Scanln(&sum)

	fmt.Printf("\nРезультат:\n\n")
	printChange(coins, sum)
	fmt.Printf("\n\nЩоб закрити програму натисніть клавішу ENTER...\n")

	fmt.Scanln()
}
