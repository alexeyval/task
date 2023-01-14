package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//
	//str, _ := reader.ReadString()
	//fmt.Print(str)

	//reader := bufio.NewReader(os.Stdin)
	//str, _ := reader.

	//for s := range str {
	//	fmt.Printf("'%v'", string(s))
	//}

	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	fmt.Print(string(str))

	//duration("", longWord)
	//f()
}

func f2() {
	var strInput string
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.ToLower(str)
	re := regexp.MustCompile("[a-z0-9]+")
	strOutput := strings.Join(re.FindAllString(str, -1), "")
	for _, i := range strOutput {
		strInput = string(i) + strInput
	}
	if strInput == strOutput {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func f() {
	var N, counter int
	reader, file := getReader("")
	defer file.Close()

	fmt.Fscan(reader, &N)
	Ns, _ := bufio.NewReader(reader).ReadString('\n')
	N, _ = strconv.Atoi(Ns)
	text, _ := bufio.NewReader(reader).ReadString('\n')
	fmt.Printf("'%v'", text)
	textTrim := strings.Trim(text, "\n")
	textSplit := strings.Split(textTrim, " ")
	for _, word := range textSplit {
		if len(word) > counter {
			counter = len(word)
		}
	}
	fmt.Println(counter)
}

func getReader(fileName string) (reader *bufio.Reader, file *os.File) {
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	return
}

func longWord(fileName string) {
	reader := bufio.NewReader(os.Stdin)
	_, _, _ = reader.ReadLine()
	line, _, _ := reader.ReadLine()

	words := strings.Split(string(line), " ")
	maxWord := words[0]
	for _, word := range words {
		if len(word) > len(maxWord) {
			maxWord = word
		}
	}
	fmt.Printf("'%v'\n'%v'\n", maxWord, len(maxWord))
}

func duration(fileName string, f func(string)) {
	start := time.Now()
	fmt.Printf("------------- %v -------------"+
		"\nВвод:\n%v\n\nВывод:\n", fileName, strings.TrimSpace(readFileContents(fileName)))
	f(fileName)
	d := time.Since(start)
	fmt.Printf("\nВремя выполнения = %v\n\n", d)
}

func readFileContents(fileName string) string {
	if fileName != "" {
		bytes, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		return string(bytes)
	}
	return ""
}
