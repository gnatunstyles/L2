package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

*/

func wget(client *http.Client, index int, url string) (int, error) {
	resp, err := client.Get("https://" + url)
	if err != nil {
		resp, err = client.Get("http://" + url)
		if err != nil {
			return -1, fmt.Errorf("error during getting to the site: %s", err)
		}
	}
	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("index_%d.html", index))
	if err != nil {
		return -1, fmt.Errorf("error during creating the template: %s", err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, fmt.Errorf("error during reading the body: %s", err)
	}
	_, err = f.Write(b)
	if err != nil {
		return -1, fmt.Errorf("error during writing data to the file: %s", err)
	}
	fmt.Printf("Content of the %s was downloaded successful.\n", url)
	index++
	return index, nil
}

func main() {
	var url string
	var err error
	c := http.Client{}
	index := 1
	fmt.Println("=== Утилита wget ===")
	for {
		fmt.Println("Enter the site:")
		fmt.Scanln(&url)
		index, err = wget(&c, index, url)
		if err != nil {
			log.Printf("error: cannto download site %s:%s", url, err)
		}
	}

}
