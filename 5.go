//https://www.hackerrank.com/challenges/s10-interquartile-range/problem
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
 * Complete the 'interQuartile' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY values
 *  2. INTEGER_ARRAY freqs
 */

func interQuartile(values []int32, freqs []int32) {
    // Print your answer to 1 decimal place within this function
    
    var morph_values []int32
    count := 0
    
    for i := 0; i < len(values); i++ {
        for j := int32 (0); j < freqs[i]; j++ {
            morph_values = append(morph_values, values[i])
            count += 1
        }
    }
    
    sort.Slice(morph_values, func(i, j int) bool { return morph_values[i] < morph_values[j] })
    
    pre_mid := count/2
    post_mid := pre_mid
    if count%2 != 0 {
        post_mid += 1
    }
    interquartile := getMedian(morph_values[post_mid:]) - getMedian(morph_values[:pre_mid])
    fmt.Println(fmt.Sprintf("%.1f", (interquartile)))
}

func getMedian(arr []int32) float32{
    
    count := len(arr)
    median := float32 (0)
    
    if count%2 == 0 {
        median = (float32 (arr[count/2]) + float32 (arr[(count/2) - 1]))/2
    } else {
        median = float32 (arr[count/2])
    }
    
    return median
    
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    valTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var val []int32

    for i := 0; i < int(n); i++ {
        valItemTemp, err := strconv.ParseInt(valTemp[i], 10, 64)
        checkError(err)
        valItem := int32(valItemTemp)
        val = append(val, valItem)
    }

    freqTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var freq []int32

    for i := 0; i < int(n); i++ {
        freqItemTemp, err := strconv.ParseInt(freqTemp[i], 10, 64)
        checkError(err)
        freqItem := int32(freqItemTemp)
        freq = append(freq, freqItem)
    }

    interQuartile(val, freq)
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
