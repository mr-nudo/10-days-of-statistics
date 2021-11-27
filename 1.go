package main
import (
 "fmt"
 "sort"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 //fmt.Print("Enter text: ")
 var count int
 fmt.Scanln(&count)
 
 numberList := make(map[int]int)
 int_keys := make([]int, 0)
 sum := 0
 has_two_median := false
 median_pos := (count/2) + 1
 median := 0.0
 
 if count%2 == 0 {
  //divide by 2
  has_two_median = true   
 }
 
 
 for i := 0; i < count; i++ {
  var num int
  fmt.Scan(&num)
  if _, ok := numberList[num]; ok{
   numberList[num] += 1
  } else {
   numberList[num] = 1
   int_keys = append(int_keys, num)
  }
  sum += num
 }
 sort.Sort(sort.IntSlice(int_keys))
 //int_keys.Sort()
 //fmt.Println(int_keys)
 //fmt.Println(numberList)
 //mean
 fmt.Println(fmt.Sprintf("%.1f", (float32 (sum)/float32 (count))))

 //median & mean logic
 median_count := 0
 lowest_mean_count := -1
 lowest_mean_value := -1
 for _, key := range int_keys {
  //mean logic
  if lowest_mean_count == -1 {
   lowest_mean_count = numberList[key]
   lowest_mean_value = key
  }
  if (numberList[key] == lowest_mean_count) && (key < lowest_mean_value) {
   lowest_mean_value = key
  } else if numberList[key] > lowest_mean_count {
   lowest_mean_count = numberList[key]
   lowest_mean_value = key
  }
  //median logic
  if median_count < median_pos{
   for j := (median_count+1); j <= (median_count + numberList[key]); j++ {
    if (j == median_pos) || (has_two_median && (j == (median_pos - 1))){
        //fmt.Println(j, key)
        median += float64 (key)
    }
   }
  }
  median_count += numberList[key]
 }
 
 //median
 if has_two_median {
  fmt.Println(median/2)
 } else {
  fmt.Println(median)   
 }
 
 fmt.Println(lowest_mean_value)
}
