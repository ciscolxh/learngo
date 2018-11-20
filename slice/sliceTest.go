package slice

import (
	"fmt"
	"bytes"
)


func SliceTest() {

	var arrayList = [10] int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}

	//切片包含起始的索引，不包含终止的索引，也就是索引为3，4，5，6指向的数据
	var sliceList = arrayList[3:7]

	//切片切出完整的数组
	var list = arrayList[:]

	for key, value := range list {
		fmt.Println(key, value)
	}

	fmt.Println("数组长度", len(arrayList))
	fmt.Println("数组索引0到切片终点", cap(sliceList))
	fmt.Println("切片长度", len(sliceList))

	fmt.Println("索引3的值为", arrayList[3])
	fmt.Println("索引7的值为", arrayList[7])
	fmt.Println("索引6的值为", arrayList[6])

	for key, value := range sliceList {
		fmt.Println(key, value)
	}

	var  buffer = bytes.Buffer{}
	buffer.WriteString("abc")
	buffer.WriteString("defg")
	buffer.WriteString("highlmn")
	fmt.Println(buffer.Bytes())

}

func BytesTest(){
	by:= []byte("test")
	by = append(by, ' ')
	by = append(by, []byte("Go Demo")...)


	m := make([]byte,18)
	for key,value:=range m{
		fmt.Println(key,value)
	}
	fmt.Println(string(m))
	slice := append(append(by[:4], '-'),  by[5:]...)
	fmt.Println(string(slice))

}
