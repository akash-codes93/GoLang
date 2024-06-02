package main

import "sort"

type Node struct {
	numerator   int
	denominator int
	value       float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {

	data := []Node{}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			data = append(data, Node{
				numerator:   arr[i],
				denominator: arr[j],
				value:       float64(arr[i]) / float64(arr[j]),
			})
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].value < data[j].value
	})

	return []int{data[k-1].numerator, data[k-1].denominator}

}
