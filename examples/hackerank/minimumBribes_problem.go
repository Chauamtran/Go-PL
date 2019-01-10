package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	lengthQueue := len(q)

	var sortedSlice = make([]int32, lengthQueue)
	count_total := 0
	var a int32 = 0

	for i := 0; i < lengthQueue; i++ {
		sortedSlice[i] = int32(i) + 1
	}

	for i := 0; i < lengthQueue; i++ {

		// Find the right a position
		if i+2 < lengthQueue {
			if sortedSlice[i+2] == q[i] {
				a = int32(i + 2)
			} else if sortedSlice[i+1] == q[i] {
				a = int32(i + 1)
			} else if sortedSlice[i] == q[i] {
				continue
			} else {
				fmt.Println("Too chaotic")
				return
			}
		} else if i+1 < lengthQueue {
			if sortedSlice[i+1] == q[i] {
				a = int32(i + 1)
			} else if sortedSlice[i] == q[i] {
				continue
			}
		}

		for q[i] != sortedSlice[i] {
			sortedSlice[a], sortedSlice[a-1] = sortedSlice[a-1], sortedSlice[a]
			count_total += 1
			a -= 1
		}
	}

	fmt.Println(count_total)

	return
}

func main() {

	// q := []int32{2, 5, 1, 3, 4}
	// Results:

	// 119814
	// 120175
	// 119810
	// 119827
	// Too chaotic

	q := []int32{3, 2, 5, 1, 7, 8, 6, 10, 4, 9}
	fmt.Println("array: ", q)
	minimumBribes(q)
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// t := int32(tTemp)

	// for tItr := 0; tItr < int(t); tItr++ {
	// 	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// 	checkError(err)
	// 	n := int32(nTemp)

	// 	qTemp := strings.Split(readLine(reader), " ")

	// 	// var q []int32

	// 	// for i := 0; i < int(n); i++ {
	// 	// 	qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
	// 	// 	checkError(err)
	// 	// 	qItem := int32(qItemTemp)
	// 	// 	q = append(q, qItem)
	// 	// }

	// 	var q[8]int32{1, 2, 5, 3, 7, 8, 6, 4}

	// 	minimumBribes(q)
	// }
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
