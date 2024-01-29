package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	mas := make([]int, n, n)
	for i, _ := range mas {
		fmt.Scan(&mas[i])
	}
	var p int
	fmt.Scanln(&p)
	//fmt.Println(n, p, mas)
	fmt.Println(BinarySearch(mas, p))

}
func BinarySearch(mas []int, p int) bool {
	l := 0
	r := len(mas) - 1
	for l <= r {
		mid := l + (r-l)/2
		if mas[mid] == p {
			return true
		} else if mas[mid] > p {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
