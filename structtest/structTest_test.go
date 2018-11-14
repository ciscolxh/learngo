package structtest

import (
	"testing"
	"fmt"
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


