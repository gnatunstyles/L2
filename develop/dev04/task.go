package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func setString(s string) string {
	var setStr string
	var set []string
	tmp := make(map[string]bool, 1)
	for _, r := range s {
		tmp[string(r)] = true
	}
	for e := range tmp {
		set = append(set, e)
	}
	sort.Strings(set)
	for _, v := range set {
		setStr += v
	}
	return setStr
}

func findAnagram(s []string) map[string][]string {
	var flag bool
	res := make(map[string][]string)
	for _, elt := range s {
		if len(elt) > 1 {
			lowElt := strings.ToLower(elt)
			for k := range res {
				if setString(lowElt) == setString(k) {
					res[k] = append(res[k], lowElt)
					flag = true
				}
			}
			if !flag {
				res[lowElt] = []string{lowElt}
				sort.Strings(res[lowElt])
			}
			flag = false
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	findAnagram([]string{"пятак", "пятка", "тяпка", "тьф", "листок", "слиток", "столик", "фть", "аб", "ба"})
}
