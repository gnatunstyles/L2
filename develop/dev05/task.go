package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
*/
type args struct {
	a       int
	b       int
	context int
	count   bool
	i       bool
	v       bool
	f       bool
	n       bool
	file    string
	grep    string
}

func flagInit() *args {
	var file string
	var grep string
	after := flag.Int("A", -1, "print number strings after match")
	before := flag.Int("B", -1, "print number strings before match")
	context := flag.Int("C", -1, "print number strings around match")
	count := flag.Bool("c", false, "print number of matched strings")
	ignoreCase := flag.Bool("i", false, "ignore case")
	invert := flag.Bool("v", false, "invert matches, print non-matchable strings")
	fixed := flag.Bool("F", false, "full match with string, not pattern")
	lineNum := flag.Bool("n", false, "print number of first matching string")
	flag.Parse()

	file = flag.Args()[0]
	grep = flag.Args()[1]

	return &args{
		a:       *after,
		b:       *before,
		context: *context,
		count:   *count,
		i:       *ignoreCase,
		v:       *invert,
		f:       *fixed,
		n:       *lineNum,
		file:    file,
		grep:    grep,
	}

}

func fileInit(args *args) ([]string, error) {
	var res []string
	filename := args.file
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

func findFixed(lines []string, grep string) []string {
	var res []string
	for i := range lines {
		if lines[i] == grep {
			res = append(res, lines[i])
		}
	}
	return res
}

func ignoreCase(lines []string, grep string) []string {
	lowerGrep := strings.ToLower(grep)
	var res []string
	for i := range lines {
		if strings.Contains(strings.ToLower(lines[i]), lowerGrep) {
			res = append(res, lines[i])
		}
	}
	return res
}

func invert(lines []string, grep string) []string {
	var res []string
	for i := range lines {
		if !strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
		}
	}
	return res
}

func lineNum(lines []string, grep string) []string {
	var res []string
	for i := range lines {
		if strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
			fmt.Println(i + 1)
		}
	}
	return res
}

func after(lines []string, grep string, num int) []string {
	var res []string
	for i := range lines {
		if strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
			for j := 1; j <= num; j++ {
				if i+j < len(lines) {
					res = append(res, lines[i+j])
				} else {
					break
				}
			}
			break
		}
	}
	return res
}

func before(lines []string, grep string, num int) []string {
	var res []string
	for i := range lines {
		if strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					res = append(res, lines[i-j])
				} else {
					break
				}
			}
			break
		}
	}
	return res
}

func context(lines []string, grep string, num int) []string {
	var res []string
	for i := range lines {
		if strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
			for j := 1; j <= num; j++ {
				if i-j >= 0 {
					res = append(res, lines[i-j])
				} else {
					break
				}
			}
			for j := 1; j <= num; j++ {
				if i+j < len(lines) {
					res = append(res, lines[i+j])
				} else {
					break
				}
			}
			break
		}
	}
	return res
}

func count(lines []string, grep string) []string {
	var res []string
	var count int
	for i := range lines {
		if strings.Contains(lines[i], grep) {
			res = append(res, lines[i])
			count++
		}
	}
	fmt.Println(count)
	return res
}

func grep() {
	args := flagInit()
	lines, err := fileInit(args)
	if err != nil {
		log.Fatal(err)
	}
	if args.i {
		lines = ignoreCase(lines, args.grep)

	}
	if args.v {
		lines = invert(lines, args.grep)

	}
	if args.f {
		lines = findFixed(lines, args.grep)
	}
	if args.n {
		lines = lineNum(lines, args.grep)

	}
	if args.count {
		lines = count(lines, args.grep)
	}
	if args.a != -1 {
		lines = after(lines, args.grep, args.a)

	}
	if args.b != -1 {
		lines = before(lines, args.grep, args.b)

	}
	if args.context != -1 {
		lines = context(lines, args.grep, args.context)
	}
	for _, elt := range lines {
		fmt.Println(elt)
	}
}

func main() {
	fmt.Println("=== Утилита grep ===")
	grep()

}
