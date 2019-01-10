package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {

	var count_div, count_mod int64 = 0, 0
	// var index int64 = 0
	lengthString := int64(len(s))

	for _, value := range s {
		fmt.Println("a", string(value))
		if string(value) == "a" {
			count_div += 1
			fmt.Println("count_div", count_div)
		}
	}

	// Get division of
	tempValue := n / lengthString
	count_div *= tempValue

	for index := lengthString * tempValue; index < n; index++ {
		a := index % lengthString
		if string(s[a]) == "a" {
			count_mod += 1
			fmt.Println("count_mod", count_mod)
		}
	}

	return count_div + count_mod

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
