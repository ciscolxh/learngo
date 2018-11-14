package arraytest

import "fmt"

func ArrayCreate1() {
	var arrayList [10] int

	for i:=0;i<len(arrayList);i++{
		arrayList[i] = i*i
	}

	for i:=0;i<len(arrayList);i++{
		fmt.Println(arrayList[i])
	}

}

func ArrayCreate2() {
	var arrayList = new ([10] int)
	for i:=0;i<len(arrayList);i++{
		arrayList[i] = i*i
	}

	for i:=0;i<len(arrayList);i++{
		fmt.Println(arrayList[i])
	}
}


func ArrayCreate3() {
	var arrayList = [10] int{0,1,4,9,16,25,36,49,64,81}
	for _,value := range arrayList{
		fmt.Println(value)
	}
}

func ArrayCreate4() {
	var arrayList = [...] int{0,1,4,9,16,25,36,49,64,81}
	for _,value := range arrayList{
		fmt.Println(value)
	}
}

func ArrayCreate5() {
	var arrayList = [] int{0,1,4,9,16,25,36,49,64,81}
	for _,value := range arrayList{
		fmt.Println(value)
	}
}

//指定长度赋值按从左到右依次赋值，没有的全部为0
func ArrayCreate6() {
	var arrayList = [10] int{0,1,4}
	for _,value := range arrayList{
		fmt.Println(value)
	}
}

func ArrayCreate7()  {
	var arrayList = [10] int {3:4,5:16}
	for _,value := range arrayList{
		fmt.Println(value)
	}
}

