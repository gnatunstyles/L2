package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита telnet ===

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт)
по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу
(через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

var defaultTimeout = 10 * time.Second

func messageWriter(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		if err == io.EOF {
			conn.Close()
			fmt.Println("\nConnection closed.")
			os.Exit(0)
		}
		fmt.Fprintf(conn, ": "+strings.Trim(text, " \r\n")+"\n")
	}
}

func messageReader(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			conn.Close()
			fmt.Println("\nConnection closed.")
			os.Exit(0)
		}
		fmt.Println(message)

	}
}
func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = input[:len(input)-1]

	return input, nil
}

func telnet(timeout time.Duration, host, port string) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		log.Printf("\nCaptured %v, exiting..", <-c)
		fmt.Fprintf(conn, " Client is leaving the server..."+"\n")
		os.Exit(1)
	}()

	//reading msgs from server
	go messageReader(conn)
	//sending msgs to the server
	messageWriter(conn)
	return nil
}

func main() {
	fmt.Println("=== Утилита telnet ===")
	s, err := readLine()
	if err != nil {
		log.Fatal(err)
	}
	args := strings.Split(s, " ")
	if args[0] == "go-telnet" {
		switch len(args) {
		case 3:
			host := args[1]
			port := args[2]
			err = telnet(defaultTimeout, host, port)
			if err != nil {
				log.Fatal(err)
			}
		case 4:
			if strings.Contains(args[1], "-timeout=") {
				t := strings.Split(args[1], "=")[1]
				measure := 0
				switch t[len(t)-1] {
				case []byte("m")[0]:
					measure = 60
				case []byte("s")[0]:
					measure = 1
				default:
					fmt.Println(t[len(t)-1])
					log.Fatal("Error. Wrong type of timeout data!")
				}
				timeInt, err := strconv.Atoi(t[:len(t)-1])
				if err != nil {
					log.Fatal(err)
				}
				host := args[2]
				port := args[3]
				err = telnet(time.Duration(timeInt*measure)*time.Second, host, port)
				if err != nil {
					log.Fatal(err)
				}
			}
		default:
			log.Fatal(`Error! Wrong type of incoming commands!`)
		}

	} else {
		log.Fatal(`Wrong type of incoming command! Try type "go-telnet" again!`)
	}
}
