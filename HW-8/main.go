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
	fmt.Println(Seach(mas, p))
}
func Seach(mas []int, p int) int {
	l := 0
	r := len(mas) - 1
	var mid int
	for l <= r {
		mid = (l + (r-l)/2)
		if mas[mid] == p {
			return mid
		} else if mas[mid] > p {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return mid

}
