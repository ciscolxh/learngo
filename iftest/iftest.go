package iftest

import "fmt"

func IfTest(age int)  {
	if age>0&&age<18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("成年")
	}

}
