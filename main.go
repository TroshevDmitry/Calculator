package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	rom "github.com/brandenc40/romannumeral"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Произошла ошибка. Предварительный выход из программы. ", err)
		os.Exit(1)
	}

	fmt.Println("Результат: ", result)
}

func calculate(expression string) (string, error) {
	expression = strings.TrimSuffix(expression, "\r\n")

	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		return "", errors.New("Строка не является выражением")
	}

	var operand1, operand2 int
	var roman1, roman2 bool

	operand1, err := strconv.Atoi(tokens[0])
	if err != nil {
		operand1, err = rom.StringToInt(tokens[0])
		if err != nil {
			return "", errors.New("Неверно введен 1 операнд. Описание ошибки: " + err.Error())
		}

		roman1 = true
	}

	operator := tokens[1]

	operand2, err = strconv.Atoi(tokens[2])
	if err != nil {
		operand2, err = rom.StringToInt(tokens[2])
		if err != nil {
			return "", errors.New("Неверно введен 2 операнд. Описание ошибки: " + err.Error())
		}

		roman2 = true
	}

	if operand1 > 10 || operand2 > 10 {
		return "", errors.New("Калькулятор принимает на вход числа от 1 до 10 включительно")
	}

	if roman1 != roman2 {
		return "", errors.New("Используются одновременно разные системы счисления")
	}

	var result int
	switch operator {
	case "+":
		result = operand1 + operand2

	case "-":
		result = operand1 - operand2

	case "*":
		result = operand1 * operand2

	case "/":
		if operand2 == 0 {
			return "", errors.New("Деление на ноль")
		}

		result = operand1 / operand2

	default:
		return "", errors.New("Неверный оператор: " + operator)
	}

	string_result := strconv.Itoa(result)

	if roman1 && roman2 {
		roman_string_result, err := rom.IntToString(result)
		if err != nil {
			return string_result, errors.New("Ответ, отображаемый римскими числами, меньше 1. Описание ошибки: " + err.Error())
		}

		return roman_string_result, nil
	}

	return string_result, nil
}
