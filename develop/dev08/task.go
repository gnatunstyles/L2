package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	proc "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*/

// readLine - reads line from std input
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// remove '\n'
	input = input[:len(input)-1]

	return input, nil
}
func cd(path string) error {
	if path == "" {
		os.Chdir(os.Getenv("HOME"))
		return nil
	}
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("error during changing the directory:%s", err)
	}
	wd, _ := pwd()
	fmt.Printf("current directory:%s", wd)
	return nil

}

func pwd() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error during getting working directory:%s", err)
	}
	fmt.Println(wd)
	return wd, nil
}
func echo(s string) {
	fmt.Println(s)
}

func kill(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("error during killing the process:%s", err)
	}
	err = proc.Kill()
	if err != nil {
		return fmt.Errorf("error during killing the process:%s", err)
	}
	fmt.Printf("Process (PID%d) was killed successfilly", pid)
	return nil
}
func ps() error {
	processes, err := proc.Processes()
	if err != nil {
		return fmt.Errorf("error during getting the process:%s", err)
	}
	for _, elt := range processes {
		log.Printf("%d\t%s\n", elt.Pid(), elt.Executable())
	}
	return nil
}
func helpMsg() {
	fmt.Println("Invalid command!\n- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)\n- pwd - показать путь до текущего каталога\n- echo <args> - вывод аргумента в STDOUT\n- kill <args> - убить процесс, переданный в качесте аргумента (пример: такой-то пример)\n- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*")
}

func main() {
	fmt.Println("=== Unix shell ===")
	for {
		line, err := readLine()
		if err != nil {
			log.Fatal(err)
		}
		// quit
		if line == "quit" {
			break
		}
		argLine := strings.Split(line, " ")
		switch argLine[0] {
		case "cd":
			if len(argLine) > 1 {
				cd(argLine[1])
			} else {
				os.Chdir(os.Getenv("HOME"))
			}
		case "pwd":
			pwd()
		case "echo":
			echo(argLine[1])
		case "kill":
			pid, err := strconv.Atoi(argLine[1])
			if err != nil {
				log.Printf("error input:%s", err)
			} else {
				kill(pid)
			}
		case "ps":
			ps()
		default:
			helpMsg()
		}
	}
}
