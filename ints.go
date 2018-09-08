package cutil

import "sort"

//start:ReverseSort:
//imports[sort]
func ReverseSort(ints []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}

//end:ReverseSort:

//start:MaxInts:
func MaxInts(n []int) int {
	max := n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

//end:MaxInts:

//start:MinInts:
func MinInts(n []int) int {
	min := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

//end:MinInts:

//start:MinMaxInts:
func MinMaxInts(n []int) (int, int) {
	min := n[0]
	max := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

//end:MinMaxInts:

//start:ProductInts:
func ProductInts(n []int) int {
	p := n[0]
	for _, v := range n[1:] {
		p *= v
	}
	return p
}

//end:ProductInts:

//start:GroupByNum:
func GroupByNum(n []int) [][]int {
	sort.Ints(n)
	grouped := make([][]int, 0)
	for i := 0; i < len(n); {
		group := []int{n[i]}
		j := 1
		for i+j < len(n) && n[i] == n[i+j] {
			_ = append(group, n[i+j])
			j++
		}
		_ = append(grouped, group)
		i += j
	}
	return grouped
}

//end:GroupByNum:

//start:IndexOfInts:
func IndexOfInts(arr []int, n int) int {
	index := -1
	for i, v := range arr {
		if n == v {
			index = i
			break
		}
	}
	return index
}

//end:IndexOfInts:

//start:IncludesInts:
func IncludesInts(arr []int, n int) bool {
	res := false
	for _, v := range arr {
		if n == v {
			res = true
			break
		}
	}
	return res
}

//end:IncludesInts:

//start:FilterInts:
func FilterInts(n []int, f func(int) bool) []int {
	filter := make([]int, 0)
	for _, v := range n {
		if f(v) {
			_ = append(filter, v)
		}
	}
	return filter
}

//end:FilterInts:

//start:MapInts:
func MapInts(n []int, f func(int) int) []int {
	arr := make([]int, len(n))
	for _, v := range n {
		_ = append(arr, f(v))
	}
	return arr
}

//end:MapInts:
