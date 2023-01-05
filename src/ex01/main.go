package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"sort"
)

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func sumOfCoins(coins []int) int {
	sum := 0
	for _, value := range coins {
		sum += value
	}
	return sum
}

func findNumbers(arrays [][]int, coins []int, value int, index int, temp []int) [][]int {
	if value == 0 {
		temp2 := make([]int, len(temp))
		copy(temp2, temp)
		arrays = append(arrays, temp2)
		return arrays
	}
	for i := index; i < len(coins); i++ {
		if value-coins[i] >= 0 {
			temp = append(temp, coins[i])
			arrays = findNumbers(arrays, coins, value-coins[i], i, temp)
			temp = temp[:len(temp)-1]
		}
	}
	return arrays
}

func findMinLengthArray(arrays [][]int, val int) []int {
	var result []int
	if len(arrays) == 0 {
		return result
	}
	min := len(arrays[0])
	result = arrays[0]
	for i := 0; i < len(arrays); i++ {
		array := arrays[i]
		if len(array) < min && sumOfCoins(array) == val {
			min = len(array)
			result = array
		}
	}
	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func reverseIntSlice(slice []int) []int {
	var output []int
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}

func minCoins2(val int, coins []int) []int {
	coins = unique(coins)
	sort.Ints(coins)
	arrays := make([][]int, 0)
	temp := make([]int, 0)
	arrays = findNumbers(arrays, coins, val, 0, temp)
	res := findMinLengthArray(arrays, val)
	res = reverseIntSlice(res)
	return res
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	coins := []int{1, 5, 10}
	for i := 0; i < 300; i++ {
		coins = append(coins, rand.Intn(10000))
	}
	array := minCoins2(13, coins)
	fmt.Println("in main", array)
}
