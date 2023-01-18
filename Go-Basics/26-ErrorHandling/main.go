package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var first, last string
	fmt.Println("Enter first name:")
	_, err := fmt.Scanf("%s", &first)
	if err != nil {
		fmt.Println("Error Scanning first name")
	}
	fmt.Println("Enter last name:")
	_, err = fmt.Scanf("%s", &last)
	if err != nil {
		log.Panic("Error Scanning first name", err)
	}
	logFile, err := os.Create("logfile")
	defer logFile.Close()
	if err != nil {
		fmt.Println("Error Creating New Logfile", err)
	}
	log.SetOutput(logFile)
	log.Println("My Name is ", first, last)

	f, err := os.Create("names.txt")
	defer f.Close()
	if err != nil {
		//log.Panic(err)
		panic(err)
	}

	r := strings.NewReader(first + " " + last)
	io.Copy(f, r)
	defer foo()
	log.Panicln("Returning....")
	log.Fatalln("Exiting....")
}

func foo() {
	fmt.Println("Foo called...")
}
