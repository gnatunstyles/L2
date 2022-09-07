package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

// Args - store arguments from CLI
type args struct {
	f []int
	d string
	s bool

	file string
}

// getArgs - returns *Args struct with parsed flags
func getArgs() (*args, error) {
	f := flag.String("f", "", "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	d := flag.String("d", "\t", "use DELIM instead of TAB for field delimite")
	s := flag.Bool("s", false, "do not print lines not containing delimiters")

	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil, errors.New("you need to specify files")
	}

	if len(*f) < 1 {
		return nil, errors.New("you need to specify a field: e.g.: 1,3")
	}

	// parsing f
	tmp := strings.Split(*f, ",")

	fields := make([]int, len(tmp))

	for i := range tmp {
		num, err := strconv.Atoi(tmp[i])
		if err != nil || num == 0 {
			return nil, fmt.Errorf("cannot convert string to int: %v", err)
		}
		fields[i] = num
	}

	args := &args{
		f: fields,
		d: *d,
		s: *s,
	}

	args.file = flag.Args()[0]

	return args, nil
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

func fields()    {}
func delimeter() {}
func separated() {}

func main() {

}
