package structtest

import (
	"fmt"
)


type Xiaohei struct {
	Name string `json:"name"`//通过json的后缀可以指定转出的json的字段名
	Age int `json:"age"`
	Sex string `json:"sex"`
}

//定义一个长方形
type Rectangle struct {
	Lengths float32
	Width float32
}

func CreateStruct() {
	var mine = new(Xiaohei)
	mine.Age = 16
	mine.Name = "xiaohei"
	mine.Sex = "man"
	fmt.Println(mine)

	luoxiaohei := Xiaohei{Name:"luoxiaohei",Age:18,Sex:"man"}
	fmt.Println(luoxiaohei)

	xiaoxiaohei:= Xiaohei{}
	xiaoxiaohei.Name = "xiaohei"
	xiaoxiaohei.Age = 16
	xiaoxiaohei.Sex = "man"
	fmt.Println(xiaoxiaohei)
}

//计算面积
func (c Rectangle) Area() float32  {
	return c.Lengths * c.Width
}

//计算周长
func (c Rectangle) Primeter() float32  {
	return 2 *(c.Width + c.Lengths)
}



