package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

/*
 * Complete the 'stdDev' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func stdDev(arr []int32) {
    sum := int32 (0)
    count := len(arr)
    // mean
    for i := 0; i < count; i++ {
        sum += arr[i]
    }
    
    mean := float64 (sum) / float64 (count)
    //fmt.Println(sum, mean)
    deviation := float64 (0)
    
    //standard deviation
    for i := 0; i < count; i++ {
        //fmt.Println(deviation)
        deviation += math.Pow((float64 (arr[i]) - mean),2)
    }
    
    deviation /= float64 (count)
    standard_dev := math.Sqrt(deviation)
    fmt.Println(fmt.Sprintf("%.1f", (standard_dev)))

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    valsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var vals []int32

    for i := 0; i < int(n); i++ {
        valsItemTemp, err := strconv.ParseInt(valsTemp[i], 10, 64)
        checkError(err)
        valsItem := int32(valsItemTemp)
        vals = append(vals, valsItem)
    }

    stdDev(vals)
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
