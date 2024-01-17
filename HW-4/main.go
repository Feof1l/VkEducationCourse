package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	m := make([]int, n, n)
	for i, _ := range m {
		fmt.Scan(&m[i])
	}

	res := make([]int, n, n)
	BottomUpMergeSort(m, res, n)
	//fmt.Println(res)
	for i, _ := range res {
		fmt.Print(res[i], " ")
	}
}

// функция слияния двух отсортированных массивов в один отсортированный
func Merge(a []int, iBegin, iMiddle, iEnd int, b []int) {
	i := iBegin
	j := iMiddle
	for k := iBegin; k < iEnd; k++ {
		if i < iMiddle && (j >= iEnd || a[i] <= a[j]) {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}
}
func copyArray(a []int, iBegin, iEnd int, b []int) {
	for k := iBegin; k < iEnd; k++ {
		b[k] = a[k]
	}
}
func TopDownSplitMerge(b []int, iBegin, iEnd int, a []int) {
	if iEnd-iBegin < 2 {
		return
	}
	iMiddle := (iEnd + iBegin) / 2
	TopDownSplitMerge(a, iBegin, iMiddle, b)
	TopDownSplitMerge(a, iMiddle, iEnd, b)
	Merge(b, iBegin, iMiddle, iEnd, a)
}
func TopDownMergeSort(a []int, b []int, n int) {
	copyArray(a, 0, n, b)
	TopDownSplitMerge(b, 0, n, a)
}

func BottomUpMergeSort(a []int, b []int, n int) {
	for width := 1; width < n; width = 2 * width {
		for i := 0; i < n; i = i + 2*width {
			Merge(a, i, min(i+width, n), min(i+2*width, n), b)
		}
		copyArray(b, 0, n, a)
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
