package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabToRoman = []struct {
	Value   int
	Numeral string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	readString, _ := reader.ReadString('\n')
	input := strings.TrimSpace(readString)
	s := calculate(input)

	fmt.Printf("Результат: %v", s)
}

func calculate(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("на ввод подали не 3 знака")
	}
	a, b := parts[0], parts[2]
	operator := parts[1]

	aInt, isARoman := romanToArab[a]
	bInt, isBRoman := romanToArab[b]

	if isARoman && isBRoman {
		return calculateRoman(aInt, bInt, operator)
	}

	aInt, errA := strconv.Atoi(a)
	bInt, errB := strconv.Atoi(b)

	if errA == nil && errB == nil {
		if aInt < 1 || aInt > 10 || bInt < 1 || bInt > 10 {
			panic("только числа в диапозоне от 1 до 10")
		}
		return calculateArab(aInt, bInt, operator)
	}

	panic("используйте только арабские цифры или только римские цифры в пределах от 1(I) до 10(X)")
}

func calculateArab(a, b int, operator string) string {
	switch operator {
	case "+":
		return strconv.Itoa(a + b)
	case "-":
		return strconv.Itoa(a - b)
	case "*":
		return strconv.Itoa(a * b)
	case "/":
		return strconv.Itoa(a / b)
	default:
		panic("неверный математический оператор")
	}
}

func calculateRoman(a, b int, operator string) string {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("неверный математический оператор")
	}
	if result <= 0 {
		panic("результат работы меньше единицы вне допустимого диапазона для римских чисел")
	}

	return convertToRoman(result)
}

func convertToRoman(i int) string {
	var result strings.Builder
	for _, v := range arabToRoman {
		for i >= v.Value {
			result.WriteString(v.Numeral)
			i -= v.Value
		}
	}
	return result.String()
}
