package main

import (
	"bufio"
	"fmt"
	"os"
)

var Verbose bool = false

func main() {
	if args := os.Args; len(args) > 1 {
		opt := args[1]
		menu(opt)
	} else {
		fmt.Println("Invalid number of arguments")
	}
}

func menu(opt string) {
	switch opt {
	case "-i":
		fmt.Println("Interactive mode")
		interactive()
	case "-h":
		fmt.Println("Help mode")
		help()
	case "-v":
		fmt.Println("Convert in verbose mode")
		verbose()
	case "-c":
		fmt.Println("filename in arguments")
		name1, name2 := getArgs()
		convert(name1, name2)

	default:
		fmt.Println("Invalid arguments provided. ")
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func interactive() {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter the file you wish to convert: ")
	reader.Scan()
	file1_str := reader.Text()

	fmt.Printf("\nEnter the file you wish to convert to: ")
	reader.Scan()
	file2_str := reader.Text()

	convert(file1_str, file2_str)

}

func convert(fileName1 string, fileName2 string) {
	bytes, error_1 := os.ReadFile(fileName1)

	check(error_1)
	if Verbose {
		fmt.Printf("Successfully converted %v to bytes.\n", fileName1)
	}
	error_2 := os.WriteFile(fileName2, bytes, 0777)
	check(error_2)
	if Verbose {
		fmt.Printf("Successfully converted bytes from %v to text in %v.\n", fileName1, fileName2)
	}
}

func help() {
	fmt.Println("./output [args: -v | -i | -h | -c ] [*.* & *.*]")
	fmt.Println("-v: verbose")
	fmt.Println("-i: interactive")
	fmt.Println("-h: help")
	fmt.Println("-c: converter")

}

func verbose() {
	Verbose = true
	name1, name2 := getArgs()
	convert(name1, name2)
}

func getArgs() (string, string) {

	fileName1 := os.Args[2]
	fileName2 := os.Args[3]

	return fileName1, fileName2

}
