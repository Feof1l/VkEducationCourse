package main

import (
	"fmt"
	"strings"
)

func main() {
	/*var s string
	fmt.Scanln(&s)
	fmt.Println(maxLiteral(s))
	*/
	var a, b string
	fmt.Scan(&a, &b)
	a = strings.ReplaceAll(a, `"`, ``)
	a = strings.ReplaceAll(a, `,`, ``)
	b = strings.ReplaceAll(b, `"`, ``)
	fmt.Println(anagramma(a, b))
}
func maxLiteral(s string) int {
	m := make(map[string]int)
	for _, v := range s {
		if _, ok := m[string(v)]; ok {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}
	max := -1
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	return max
}
func anagramma(a, b string) bool {
	m := make(map[string]int)
	for _, v := range a {
		if _, ok := m[string(v)]; ok {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}
	for _, v := range b {
		if _, ok := m[string(v)]; ok {
			m[string(v)]--
		} else {
			return false
		}
	}

	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}
