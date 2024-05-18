package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var res int
var op1 int
var op2 int
var intCount, romanCount int

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}

var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	} else {
		return false
	}
}

func isRoman(s string) bool {
	if roman[s] != 0 {
		return true
	} else {
		return false
	}
}

func getoperand(op string) int {
	res, _ := strconv.Atoi(op)

	if (res < 1) || (res > 10) {
		panic(fmt.Sprintf("Число не может быть <1 или >10"))
	}
	return res
}

func intToRoman(romanResult int) string {
	var romanNum string
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	return romanNum
}
func main() {
	var a, opr, b string
	fmt.Println("Введите операцию")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.Fields(console)

		if len(s) == 1 {
			panic(fmt.Sprintf("Строка не является математической операцией"))
		}
		if len(s) > 3 {
			panic(fmt.Sprintf("Слишком много аргументов"))
		}

		a = s[0]
		b = s[2]
		opr = s[1]
		opsslice := [2]string{a, b}
		romanCount = 0
		intCount = 0

		for i := 0; i < len(opsslice); i++ {
			if isInt(opsslice[i]) {
				intCount++
			}

			if isRoman(opsslice[i]) {
				romanCount++
			}
		}
		if intCount == 2 {
			op1 = getoperand(a)
			op2 = getoperand(b)

		} else if romanCount == 2 {
			op1 = roman[a]
			op2 = roman[b]
		} else {
			panic(fmt.Sprintf("Неверное сочитание цифр"))
		}

		switch opr {
		case "+":
			res = op1 + op2

		case "-":
			res = op1 - op2

		case "*":
			res = op1 * op2

		case "/":
			res = op1 / op2

		default:
			panic(fmt.Sprintf("Неправильная операция"))
		}
		if (res < 1) && (romanCount == 2) {
			panic(fmt.Sprintf("Римкие цифры не могут быть отрицательными"))
		} else if romanCount == 2 {
			fmt.Printf("Ваш результат: %s\n", intToRoman(res))
		} else {
			fmt.Printf("Ваш результат: %v\n", res)
		}
	}
}
