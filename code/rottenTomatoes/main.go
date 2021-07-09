package main

import (
	"container/list"
	"fmt"
)

func rotAdjacentTomatoe(arr [][]int, index []int, m int, n int, rottenTomatoeQueue *list.List) [][]int {
	i, j := index[0], index[1]

	if i-1 >= 0 && arr[i-1][j] == 1 {
		arr[i-1][j] = 2
		rottenTomatoeQueue.PushBack([3]int{i - 1, j, 2})
	}
	if j-1 >= 0 && arr[i][j-1] == 1 {
		arr[i][j-1] = 2
		rottenTomatoeQueue.PushBack([3]int{i, j - 1, 2})
	}

	if i+1 < m && arr[i+1][j] == 1 {
		// fmt.Println(i+1, n)
		arr[i+1][j] = 2
		rottenTomatoeQueue.PushBack([3]int{i + 1, j, 2})
	}

	if j+1 < n && arr[i][j+1] == 1 {
		arr[i][j+1] = 2
		rottenTomatoeQueue.PushBack([3]int{i, j + 1, 2})
	}

	return arr
}

func checkRottenTomatoe(grid [][]int) bool {

	for _, arr := range grid {
		for _, value := range arr {

			if value == 1 {
				return false
			}
		}
	}
	return true
}

func main() {

	pl := fmt.Println
	// p := fmt.Printf

	// 0 empty cell
	// 1 good tomatoe
	// 2 bad tomatoe
	// 3 time delimeter

	grid := [][]int{
		{2, 1, 0, 2, 1},
		{1, 0, 1, 2, 1},
		{1, 0, 0, 2, 1},
	}

	rottenTomatoeQueue := list.New()

	for i, arr := range grid {
		for j, value := range arr {

			if value == 2 {
				rottenTomatoeQueue.PushBack([3]int{i, j, value})
			}
		}
	}

	// enqueue time delimeter
	rottenTomatoeQueue.PushBack([3]int{-1, -1, 3})

	queueLen := rottenTomatoeQueue.Len()
	timeFrame := 0
	// _ = timeFrame

	for queueLen > 0 {

		front := rottenTomatoeQueue.Front()
		rottenTomatoe := front.Value

		rottenTomatoe1 := rottenTomatoe.([3]int)

		if rottenTomatoe1[2] == 3 { // it is time delimeter
			if queueLen > 1 {
				timeFrame += 1
				rottenTomatoeQueue.PushBack([3]int{-1, -1, 3})
			} else {
				break
			}

		} else if rottenTomatoe1[2] == 2 {
			grid = rotAdjacentTomatoe(grid, rottenTomatoe1[:], 3, 5, rottenTomatoeQueue)
		}

		rottenTomatoeQueue.Remove(front)
		queueLen = rottenTomatoeQueue.Len()
	}

	if checkRottenTomatoe(grid) {
		pl("Total timeframe is: ", timeFrame)
	} else {
		pl("All tomatoe cannot be rotten.")
	}

}
