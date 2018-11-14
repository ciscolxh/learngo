package swichtest

import "fmt"

func SwitchTest(grade float32) {
	switch {
	case grade >= 0 && grade < 60:
		fmt.Println("E")
	case grade >= 60 && grade < 70:
		fmt.Println("D")
	case grade > 70 && grade < 80:
		fmt.Println("C")
	case grade >= 80 && grade < 90:
		fmt.Println("B")
	case grade >= 90 && grade <= 100:
		fmt.Println("A")
	default:
		fmt.Println("输入错误")
	}
}

func SwitchTests (a int,b int ,symbol string){
	switch symbol {
	case "+":
		fmt.Println(a+b)
	case "-":
		fmt.Println(a-b)
	case "*":
		fmt.Println(a*b)
	case "/":
		fmt.Println(float64(a) / float64(b))
	default:
		fmt.Println("输入有误")
	}
}
