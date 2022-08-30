package main

import (
	"errors"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//метод распаковки строки
func unpack(s string) (string, error) {
	res := ""
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		_, err := strconv.Atoi(string(arr[i]))
		if err != nil {
			if i != len(arr)-1 {
				if count, err := strconv.Atoi(string(arr[i+1])); err != nil {
					res += string(arr[i])
				} else {
					for j := 0; j < count; j++ {
						res += string(arr[i])
					}
					i++
				}
			} else {
				res += string(arr[i])
			}
		} else {
			return "", errors.New("error: wrong input format")
		}
	}
	return res, nil

}

func main() {}
