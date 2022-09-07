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

// readFile - reads file line by line
func readFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// cutLines - cut lines by specified args
func cutLines(args *args, lines []string) []string {
	delimiter := "\t"

	if args.d != delimiter {
		delimiter = args.d
	}

	var result []string

	for _, line := range lines {
		if delimiter != "" && strings.Contains(line, delimiter) {
			words := strings.Split(line, delimiter)

			cutLine := strings.Builder{}

			for _, val := range args.f {
				if len(words) >= val {
					cutLine.WriteString(words[val-1])
					cutLine.WriteString(delimiter)
				}
			}

			// trim extra delimiter
			result = append(result, strings.TrimSuffix(cutLine.String(), delimiter))

		} else if !args.s {
			result = append(result, line)
		}
	}

	return result
}

// cut - works like linux cut with flags:
// -f -s -d
func cut() ([]string, error) {
	if len(os.Args) < 4 {
		return nil, errors.New("not enougth elements")
	}

	var result []string

	args, err := getArgs()
	if err != nil {
		return nil, err
	}

	lines, err := readFile(args.file)
	if err != nil {
		return nil, err
	}

	cutLines := cutLines(args, lines)
	result = append(result, cutLines...)

	return result, nil
}

func main() {

	lines, err := cut()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := range lines {
		fmt.Println(lines[i])
	}
}
