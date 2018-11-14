package main

import "fmt"
import (
	"learngo/consttest"
)




var (
	a =1
	b ="hello world"
)


func main() {
	test()
	prt()
	testa()
	consttest.Consts()
	println(consttest.Orange)
}

func test()  {
	var  (
		name = "小黑"
		age = 18
	)
	fmt.Println(name,age)
}

func testa()  {
	var i = "i"
	var  e  = "e"
	fmt.Println(i,e)
}

func prt()  {
	fmt.Println(a,b)
}
