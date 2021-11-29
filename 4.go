//https://www.hackerrank.com/challenges/s10-quartiles/problem
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'quartiles' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func quartiles(arr []int32) []int32 {
    // Write your code here
    var result []int32
    
    //sort array
    //sort.Sort(sort.IntSlice(arr))
    //fmt.Println(arr)
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    //fmt.Println(arr)
    count := len(arr)
    pre_mid := count/2
    post_mid := pre_mid
    if count%2 != 0 {
        post_mid += 1
    }
    println(pre_mid, post_mid)
    //median := int32 (0)
    //median_2 := int32 (0)
    //var median_3 int32
    result = append(result, getMedian(arr[:pre_mid]), getMedian(arr), getMedian(arr[post_mid:]))
    
    return result
}

func getMedian(arr []int32) int32{
    
    count := len(arr)
    median := int32 (0)
    
    if count%2 == 0 {
        median = (arr[count/2] + arr[(count/2) - 1])/2
    } else {
        median = arr[count/2]
    }
    
    return median
    
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    dataTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var data []int32

    for i := 0; i < int(n); i++ {
        dataItemTemp, err := strconv.ParseInt(dataTemp[i], 10, 64)
        checkError(err)
        dataItem := int32(dataItemTemp)
        data = append(data, dataItem)
    }

    res := quartiles(data)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
