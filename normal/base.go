package main

import "fmt"

func main() {
	//写出输出结果
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}
