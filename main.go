package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// slice := []int{1, 2, 3, 4}
	// newslice := slice[1:3] // [2, 3]
	// newslice = append(newslice, 6) // [2, 3, 6]
	// newslice = append(newslice, 4, 5) //[2 3 6 4 5]
	// fmt.Println(newslice)

	// arr1 := [4]int{1,2,3,4}
	// arr2 := arr1[0:3];
	// arr2[1] = 15
	// fmt.Println(arr1, arr2)

	// test := []int{}

	// for i:=0; i < 10; i++ {
	// 	test = append(test, i)
	// 	fmt.Println("slice, capacity", len(test), cap(test))
	// }

	// str := "Hello world"
	// isContain := strings.Contains(str, "He") // true
	// upperCase := strings.ToUpper(str)
	// replaseString := strings.Replace(str, "world", "hiren", -1) // n is a limitation of replace ment if it lessthan 0 than no limitation
	// splitString := strings.Split(str, " ") // [Hello world]
	// joinString := strings.Join([]string{"a","b","c"}, ".") //a.b.c
	// fmt.Println("isContain", joinString)


	//sorting and searching
	// arr := []int{2,4,5,3}
	// str := []string{"hello", "my", "love"}
	// sort.Strings(str)
	// fmt.Println(str)

	//time 
	// time.Now() // get current time 
	// time.Sleep(2 * time.Second) // stop exicution 2 second

	// json
	obj := map[string]interface{}{
		"name": "hiren",
		"age": 20,
	}
	jsonData, _ := json.Marshal(obj)
	fmt.Println(string(jsonData))

	newObj := map[string]interface{}{}
	json.Unmarshal([]byte(jsonData), &newObj)
	fmt.Println(newObj)
}
