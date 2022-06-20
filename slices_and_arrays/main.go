package main

import (
  "fmt"
  "reflect"
)

func main() {
  // arrays
  var int_array [5]int
  int_array[0] = 101
  fmt.Printf("%v: %v\n",reflect.TypeOf(int_array), int_array)

  string_array := [2]string{"string_one", "string_two"}
  fmt.Printf("%v: %v\n",reflect.TypeOf(string_array), string_array)

  string_array = [...]string{"string_one", "string_three"}
  fmt.Printf("%v: %v\n",reflect.TypeOf(string_array), string_array)

  // slices
  int_slice := []int{1,2,3,4,5,6}
  fmt.Printf("%v: %v\n",reflect.TypeOf(int_slice), int_slice)

  int_slice2 := make([]int, 3, 3)
  fmt.Printf("%v: %v\n",reflect.TypeOf(int_slice2), int_slice2)
  fmt.Printf("type: %v, len: %v, cap: %v\n",
    reflect.TypeOf(int_slice2), len(int_slice2), cap(int_slice2))

  int_slice2 = make([]int, 6)
  fmt.Printf("%v: %v\n",reflect.TypeOf(int_slice2), int_slice2)
  fmt.Printf("type: %v, len: %v, cap: %v\n",
    reflect.TypeOf(int_slice2), len(int_slice2), cap(int_slice2))

  // slices when created from another slice point to the same memory location
  string_slice := []string{
    "I",
    "Like",
    "Pizza",
  }
  // making a slice with the last element of the original ["pizza"]
  another_string_slice := string_slice[len(string_slice)-1:]
  fmt.Println("another_string_slice: %v", another_string_slice)
  //updating ["pizza"] to equal ["Cake"] in the new slice
  another_string_slice[0] = "Cake"
  // now the original slice is modified as well...
  // Interesting, guess we are using pointers here behind the scenes
  fmt.Printf("original_strig_slice: %v\n", string_slice)



  myArray := [2]int{1, 2}
  mySliceOfArray := myArray[:]
  // append seems to make a new slice in memory
  myAppendedSlice := append(mySliceOfArray, 3, 4, 5)
  fmt.Println(myArray)
  fmt.Println(myAppendedSlice)

}
