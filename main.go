package main

import (
	"custom_types/types"
	"fmt"
)

func main() {
	var z custom_types.Array
	z.Add("dsadas1") //0
	z.Add("dsadas2") //1
	z.Add(3) //2
	z.Add("dsadas4") //3
	z.Add(true) //4
	fmt.Println(z)
	fmt.Println(z.Contains("dsadas3"))
	fmt.Println(z.Contains(3))
	fmt.Println(z.Contains(true))

}