// Package docs is an educational package used to learn how to document Go code
package docs

/* intall godoc - go get golang.org/x/tools/cmd/godoc - if you don't have it
to create html doc run `godoc -url "pkg/docs" > *.html` where * - is preferred name for webpage or
you can go doc -http=:port run localhost:port, go to pkg, find your package and extract it through browser*/

import (
	"sort"
)

// MinCoins is an example given at ex00
// which is should not be used anywhere.
func MinCoins(val int, coins []int) []int {
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

// SumOfCoins summarize integers of integer slice and returns the sum.
func SumOfCoins(coins []int) int {
	sum := 0
	for _, value := range coins {
		sum += value
	}
	return sum
}

// FindNumbers is a function that computes combinational sum recursively
// and return 2-dimensional slice of ints. Very slow.
func FindNumbers(arrays [][]int, coins []int, value int, index int, temp []int) [][]int {
	if value == 0 {
		temp2 := make([]int, len(temp))
		copy(temp2, temp)
		arrays = append(arrays, temp2)
		return arrays
	}
	for i := index; i < len(coins); i++ {
		if value-coins[i] >= 0 {
			temp = append(temp, coins[i])
			arrays = FindNumbers(arrays, coins, value-coins[i], i, temp)
			temp = temp[:len(temp)-1]
		}
	}
	return arrays
}

// findMinLentghArray is straight-forward function that iterates through
// 2-dimensional slice and returns first slice with the least amount of integers in it.
func FindMinLengthArray(arrays [][]int, val int) []int {
	var result []int
	if len(arrays) == 0 {
		return result
	}
	min := len(arrays[0])
	result = arrays[0]
	for i := 0; i < len(arrays); i++ {
		array := arrays[i]
		if len(array) < min && SumOfCoins(array) == val {
			min = len(array)
			result = array
		}
	}
	return result
}

// unique function creates and returns new integer slice
// without repeated elements.
func Unique(intSlice []int) []int {
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

// reverseIntSlice creates and returns reversed slice.
func ReverseIntSlice(slice []int) []int {
	var output []int
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}

// minCoins2 creates and returns the smallest slice
// constructed from coins which combinational sum is equal to val.
func minCoins2(val int, coins []int) []int {
	coins = Unique(coins)
	sort.Ints(coins)
	arrays := make([][]int, 0)
	temp := make([]int, 0)
	arrays = FindNumbers(arrays, coins, val, 0, temp)
	res := FindMinLengthArray(arrays, val)
	res = ReverseIntSlice(res)
	return res
}
