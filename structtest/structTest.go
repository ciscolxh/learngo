package structtest

import (
	"fmt"
)

type Xiaohei struct {
	Name string `json:"name"` //通过json的后缀可以指定转出的json的字段名
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}


//定义一个长方形
type Rectangle struct {
	Lengths float32
	Width   float32
}

//动物类
type Animal struct {
	Leg   int
	Head  int
	Mouth int
	Food  string
	string
}

//猫和动物结构体组合
type Cat struct {
	Tail int
	Animal
}

//鸟和动物结构体组合
type Bird struct {
	Wing int
	Leg string
	Animal
}

type Engine interface {
	Stop()
	Start()
}

type Car struct {
	Engine
}

//吃东西
func (a *Animal) Eat() {
	fmt.Println(a.string + " eat " + a.Food)
}

func CreateStruct() {
	var mine = new(Xiaohei)
	mine.Age = 16
	mine.Name = "xiaohei"
	mine.Sex = "man"
	fmt.Println(mine)

	luoxiaohei := Xiaohei{Name: "luoxiaohei", Age: 18, Sex: "man"}
	fmt.Println(luoxiaohei)

	xiaoxiaohei := Xiaohei{}
	xiaoxiaohei.Name = "xiaohei"
	xiaoxiaohei.Age = 16
	xiaoxiaohei.Sex = "man"
	fmt.Println(xiaoxiaohei)
}

//计算面积
func (c Rectangle) Area() float32 {
	return c.Lengths * c.Width
}

//计算周长
func (c Rectangle) Primeter() float32 {
	return 2 * (c.Width + c.Lengths)
}

func CarStart() {
	car := new (Car)
	car.Start()
	car.Stop()
}

func (c Car) Start() {
	fmt.Println("start")
}

func (c Car) Stop() {
	fmt.Println("stop")
}
