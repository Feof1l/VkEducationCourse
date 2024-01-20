package main

import (
	"fmt"
)

func main() {
	/*var s string
	fmt.Scanln(&s)
	fmt.Println(string(deleteDouble(s)))*/
	var n int
	fmt.Scanln(&n)
	mas := make([]int, n, n)
	for i, _ := range mas {
		fmt.Scan(&mas[i])
	}
	fmt.Println(stack(n, mas))

}
func deleteDouble(s string) string {
	//stack:=make([]rune,0)
	r := []rune(s)

	for i := 0; i < len(r); i++ {
		//fmt.Println(len(r), i)
		if len(r) >= i+2 && r[i] == r[i+1] {

			r = append(r[:(i)], r[(i+2):]...)
			//fmt.Println(string(r))
			i = -1

		}
		if len(r) == 2 && r[0] == r[1] {
			return ""
		}
		if len(r) == 0 {
			return ""
		}
		//fmt.Println(i, len(s))
	}
	var str string
	str = string(r)
	return str
}

//Дана строка s. Строка состоит из английских букв в нижнем регистре.Необходимо из строки удалить все рядом стоящие повторяющиеся буквы.
//Например, в строке xyyx мы должны удалить yy, а после и оставшиеся xx и того после должна получиться пустая строка. Или же в строке fqffqzzsd после удаления остануться только fsd. Первыми удаляться ff,
//являющимися третьими и четвертыми символами, затем qq и после уже zz.

func stack(n int, mas []int) int {
	s := make([]int, 0)
	for _, v := range mas {
		if v%2 == 0 {
			s = append(s, v)
		}
	}
	if len(s) != 0 {
		return s[len(s)-1]
	} else {
		return -1
	}

}

//Дан массив не отсортированных целых чисел. Написать функцию, которая вернет первое с конца четное число.
//При написании кода используйте принцип стека. Если массив не содержит четного числа возвращать -1.
