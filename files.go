package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// takes string parameter, confirms is a directory then lists content
// directories have + while files have -
func ListDir(dir string) {
	isDir, err := os.Stat(dir)
	if err != nil {
		log.Fatal(err)
	}
	if isDir.IsDir() == true {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if file.IsDir() == true {
				fmt.Println("+", file.Name())
			} else {
				fmt.Println("-", file.Name())
			}
		}
	}
}

func main() {
	currentDir := "."
	scanner := bufio.NewScanner(os.Stdin)
	ListDir(currentDir)
	fmt.Println("-> ")
	for scanner.Scan() {
		if scanner.Text() == "QUIT" {
			break
		}
		// check if given file is a directory
		if tmp, err := os.Stat(currentDir + "/" + scanner.Text()); err != nil {
			log.Fatal(err)
		} else if tmp.IsDir() == true {
			currentDir = currentDir + "/" + scanner.Text()
		}
		ListDir(currentDir)
		fmt.Println("-> ")
	}
}
