package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	mas := make([]int, n, n)
	for i, _ := range mas {
		fmt.Scan(&mas[i])
	}
	fmt.Println(hash(mas))
	/*var n int
	fmt.Scanln(&n)
	//
	mas := make([]int, n, n)
	for i, _ := range mas {
		fmt.Scan(&mas[i])
	}
	var p int
	fmt.Scanln(&p)
	res := ExpSearch(mas, p)

	fmt.Println(res-1, res+1)*/
}
func BinarySearch(mas []int, p int, l, r int) int {

	for l <= r {
		mid := l + (r-l)/2
		if mas[mid] == p {
			return mid
		} else if mas[mid] > p {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func ExpSearch(mas []int, p int) int {
	border := 1
	lastElem := len(mas) - 1
	for border < lastElem && mas[border] < p {
		border = border * 2
		if mas[border] == p {
			return border
		}
		if border > lastElem {
			border = lastElem
		}

	}
	//res := BinarySearch(mas, p, border/2, border)
	return border
}
func hash(mas []int) int {
	m := make(map[int]int)
	max := len(mas) / 2
	for _, v := range mas {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for key, value := range m {
		if value > max {
			return key
		}
	}
	return -1
}
