package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'weightedMean' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY X
 *  2. INTEGER_ARRAY W
 */

func weightedMean(X []int32, W []int32) {
    // Write your code here
    var numerator int32
    var denominator int32
    for i := 0; i < len(X); i++ {
        numerator += (X[i] * W[i])
        denominator += W[i]
    }
    fmt.Println(fmt.Sprintf("%.1f", (float32 (numerator)/float32 (denominator))))
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

    weightsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var weights []int32

    for i := 0; i < int(n); i++ {
        weightsItemTemp, err := strconv.ParseInt(weightsTemp[i], 10, 64)
        checkError(err)
        weightsItem := int32(weightsItemTemp)
        weights = append(weights, weightsItem)
    }

    weightedMean(vals, weights)
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
