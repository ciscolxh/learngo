package structtest

import (
	"testing"
	"fmt"
	"common"
)

var r =Rectangle{Lengths:2.5,Width:3.5}

var r1 = new (Rectangle)

func TestCreateStruct(t *testing.T) {
	CreateStruct()
}

func TestRectangle_Area(t *testing.T) {
	r1.Lengths = 3
	r1.Width = 5
	fmt.Println(r1.Area())
}

func TestRectangle_Primeter(t *testing.T) {
	fmt.Println(r.Primeter())
}

func TestAnimal_Eat(t *testing.T) {
	var cat = new(Cat)
	cat.Food = "Mouse"
	cat.string = "Cat"
	cat.Head = 1
	cat.Leg = 4
	cat.Mouth = 1
	cat.Tail = 1
	cat.Eat()

	fmt.Println(common.JsonEncode(cat))

	var a = new(Animal)
	a.Food = "insect"
	a.Head = 1
	a.Mouth = 1
	a.Leg = 2
	a.string = "bird"
	var b = new(Bird)
	b.Animal = *a
	b.Wing = 2
	b.Leg = "leg"
	b.Eat()
	fmt.Println(common.JsonEncode(b))
}

func TestCarStart(t *testing.T) {
	CarStart()
}







