package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*var n int
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
	}*/
	var n int
	fmt.Scanln(&n)
	count := make([]int, n, n)
	name := make([]string, n, n)
	year := make([]int, n, n)
	matrix := make([][]string, n, n)
	for i, _ := range matrix {
		matrix[i] = make([]string, 3, 3)
		for j, _ := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			switch j {
			case 0:
				{
					elem, _ := strconv.Atoi(matrix[i][j])
					//count = append(count, elem)
					count[i] = elem
				}
			case 1:
				{
					//name = append(name, matrix[i][j])\
					name[i] = matrix[i][j]
				}
			case 2:
				{
					elem, _ := strconv.Atoi(matrix[i][j])
					//year = append(year, elem)
					year[i] = elem
				}
			}

		}
	}
	fmt.Println(name)
	fmt.Println(year)
	fmt.Println(count)
	fmt.Println()
	/////////////
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			fmt.Print(matrix[i][j], " ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()
	res := make([]int, n, n)
	BottomUpMergeSort(year, res, n)
	for i, _ := range res {
		fmt.Print(res[i], " i=", i, " ")
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
