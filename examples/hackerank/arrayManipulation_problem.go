package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
    var sortedSlice = make([]int64, n)
    var maxValue, maxTotal int64 {0,0}
    
    maxI = 0
    
    m := len(queries)
    for i:=0; i<m; i++ {
        maxTotal += queries[i][2]
        sortedSlice[queries[i][0]-1] += queries[i][2]
        sortedSlice[queries[i][1]-1] += queries[i][2]

        if maxValue < sortedSlice[queries[i][2]] {
            maxValue = sortedSlice[queries[i][0]]
        }

        if maxValue < sortedSlice[queries[i][1]] {
            maxValue = sortedSlice[queries[i][1]]
        }
        // for j:=queries[i][0]-1; j<=queries[i][1]-1; j++ {
        //     sortedSlice[j] += int64(queries[i][2])
        //     if maxValue < sortedSlice[j] {
        //         maxValue = sortedSlice[j]
        //     }
        // }
    }

    return maxValue

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    var queries [][]int32
    for i := 0; i < int(m); i++ {
        queriesRowTemp := strings.Split(readLine(reader), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != int(3) {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := arrayManipulation(n, queries)

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
