package main;

import "fmt";

func twoSum(nums []int, target int) []int {
  var res [2]int {0,};
  for i := 0; i < len(nums); i++ {
    for j := i+1; j < len(nums); j++ {
      if(nums[i] + nums[j] == target){
        fmt.Println("Hello, World!",i);
        res[0] = i;
        res[1] = j;
      }
    }
  }
  return res;
}

func main() {
   fmt.Println("Hello, World!")
}
