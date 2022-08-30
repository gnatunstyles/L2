package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками,
на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов

*/

type args struct {
	k     int
	n     bool
	r     bool
	u     bool
	files []string
}

func sortByColumn(lines []string, k int) []string {
	var stringTable [][]string
	for i := range lines {
		stringTable = append(stringTable,
			strings.Split(strings.Join(strings.Fields(lines[i]), " "), " "))
	}
	sort.SliceStable(stringTable, func(i, j int) bool {
		fmt.Println(stringTable[i][k], stringTable[j][k])
		return stringTable[i][k] < stringTable[j][k]
	})
	for i := range lines {
		lines[i] = strings.Join(stringTable[i], " ")
	}
	return lines
}

func sortByNum(lines []string) []string {
	var stringTable [][]string
	for i := range lines {
		stringTable = append(stringTable,
			strings.Split(strings.Join(strings.Fields(lines[i]), " "), " "))
	}
	sort.SliceStable(stringTable, func(i, j int) bool {
		return stringTable[i][4] < stringTable[j][4]
	})
	for i := range lines {
		lines[i] = strings.Join(stringTable[i], " ")
	}
	return lines
}

func reverse(lines []string) []string {
	i := 0
	j := len(lines) - 1
	for i < j {
		lines[i], lines[j] = lines[j], lines[i]
		i++
		j--
	}
	return lines
}
func unique(lines []string) []string {
	var newLines []string
	set := make(map[string]bool, 1)
	for _, elt := range lines {
		set[elt] = true
	}
	for k := range set {
		newLines = append(newLines, k)
	}
	return newLines
}

func flagInit() (*args, error) {
	var files []string
	column := flag.Int("k", 8, "number of sort column")
	if *column < 0 || *column > 8 {
		return nil, errors.New("error. wrong number of column")
	}
	number := flag.Bool("n", false, "number sort")
	reverse := flag.Bool("r", false, "reversed sort")
	unique := flag.Bool("u", false, "unique sort")
	flag.Parse()

	files = append(files, flag.Args()...)

	return &args{
		k:     *column,
		n:     *number,
		r:     *reverse,
		u:     *unique,
		files: files,
	}, nil

}

func fileInit(args *args) ([]string, error) {
	var res []string
	filename := args.files[0]
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		res = append(res, line)
	}
	return res, nil
}

func sortHandler() ([]string, error) {
	args, err := flagInit()
	if err != nil {
		return []string{}, err
	}
	lines, err := fileInit(args)
	if err != nil {
		return []string{}, err
	}
	if args.u {
		lines = unique(lines)
	}

	if args.k != 8 {
		sortByColumn(lines, args.k)
	}

	if args.n {
		lines = sortByNum(lines)
	}
	if args.r {
		lines = reverse(lines)
	}

	return lines, nil
}

func main() {
	fmt.Println("=== Утилита sort ===")
	sorted, err := sortHandler()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, elt := range sorted {
		_, err = f.WriteString(fmt.Sprintf("%s\n", elt))
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("=== OK ===")

}
