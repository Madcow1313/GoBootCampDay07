package main

import "testing"

type Tests struct {
	coins, expected []int
	val             int
}

var tests = []Tests{
	{coins: []int{1, 5, 10}, expected: []int{10, 1, 1, 1}, val: 13},
	{coins: []int{1, 5, 10}, expected: []int{10}, val: 10},
	{coins: []int{3}, expected: []int{}, val: 10},
	{coins: []int{10, 1, 5}, expected: []int{10, 1, 1, 1}, val: 13},
	{coins: []int{5, 1, 2, 2}, expected: []int{2, 2}, val: 4},
	{coins: []int{}, expected: []int{}, val: 0},
	{coins: []int{99, 50}, expected: []int{50, 50}, val: 100},
	{coins: []int{50, 99}, expected: []int{50, 50}, val: 100},
	{coins: []int{99, 50}, expected: []int{99}, val: 99},
	{coins: []int{1, 1, 1, 10, 10}, expected: []int{10, 1, 1, 1}, val: 13},
	{coins: []int{1, 10, 1, 1, 10}, expected: []int{10, 1, 1, 1}, val: 13},
	{coins: []int{1, 5, 10, 50, 100, 500}, expected: []int{500, 100, 50, 10, 5, 1}, val: 666},
}

func compareSlice(first []int, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func Test(t *testing.T) {
	for n, test := range tests {
		if output := minCoins2(test.val, test.coins); !compareSlice(output, test.expected) {
			t.Errorf("test %d minCoins2 Output %d is not equal to expected %d\n", n+1, output, test.expected)
		}
		if output := minCoins(test.val, test.coins); !compareSlice(output, test.expected) {
			t.Errorf("test %d minCoins Output %d is not equal to expected %d\n", n+1, output, test.expected)
		}
	}
}
