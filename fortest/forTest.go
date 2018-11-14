package fortest

import "fmt"

//var a int

func ForTest()  {
	for i:=0;i<10 ;i++  {
		fmt.Println(i)
	}
}

func ForMore() {
	for   {
		fmt.Println("死循环")
	}
}

func ForLoop()  {
	a:=0
	for a < 5 {
		a++
		if a==3 {
			continue
		}else {
			fmt.Println(a)
		}
	}

	fmt.Println("执行结束")
}
