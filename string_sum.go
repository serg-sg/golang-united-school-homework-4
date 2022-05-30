package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
//используем эти ошибки соответствующим образом, оборачивая их функцией fmt.Errorf

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	// Использовать, когда ввод пуст, а ввод считается пустым, если строка содержит только пробелы
	// "ввод пуст"
	errorEmptyInput = errors.New("input is empty")

	// Use when the expression has number of operands not equal to two
	// Используется, когда число операндов выражения не равно двум
	// "ожидал два операнда, но получил больше или меньше"
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

// Реализовать функцию, которая вычисляет сумму двух целых чисел, записанных в виде строки
// Например, имея входную строку "3+5", он должен вернуть выходную строку "8" и нулевую ошибку
// Рассмотрим случаи, когда операнды отрицательные ("-3+5" или "-3-5") и когда входная строка содержит пробелы ("3 + 5")
//
// Для случаев, когда входное выражение неверно (содержит символы, не являющиеся числами, +, - или пробел)
// функция должна возвращать пустую строку и соответствующую ошибку из пакета strconv, завернутую в вашу собственную ошибку
// с функцией fmt.Errorf
//
// Используем ошибки, определенные выше, как описано, снова заворачивая в fmt.Errorf

func StringSum(input string) (output string, err error) {
	// Уберём пробельные символи слева и справа
	a := strings.TrimSpace(input)

	// Если длина input равна 0, то вернуть ошибку
	if len(a) == 0 {
		err = fmt.Errorf("%q", errorEmptyInput)
		return
	}

	// Замена пробелов внутри строки
	re := regexp.MustCompile(`\s`)
	a = re.ReplaceAllString(a, "")

	// Делаем обход строки с 1 индекса на поиск знаков + и - и добавляя в срез
	var x = []int{}
	y := 0
	var value int
	for i := 1; i < len(a); i++ {
		if strings.ContainsAny(string(a[i]), "+-") {
			value, err = strconv.Atoi(string(a[y:i]))
			if err == nil {
				x = append(x, value)
				y = i
			} else {
				err = fmt.Errorf("%q", err)
				return
			}
		}
	}

	// Добавляем в х хвост строки
	value, err = strconv.Atoi(string(a[y:len(a)]))
	if err != nil {
		err = fmt.Errorf("%q", err)
		return
	}
	x = append(x, value)

	// Проверяем длину среза х, если не равно 2, то это ошибка
	if len(x) != 2 {
		err = fmt.Errorf("%q", errorNotTwoOperands)
		return
	}

	// Суммируем два числа
	output = strconv.Itoa(int(x[0]) + int(x[1]))

	return output, err
}
