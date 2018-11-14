package functest

import "fmt"

/*定义结构体*/
type Circle struct {
	radius float64 `json:"radius"`
}

/**
  求两个数最大的一个
 */
func FuncTest(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func FuncMoreReturn(a, b string) (string, string) {
	return a, b
}

//值传递的方法
func swap(a, b int) {
	var c int
	c = a
	a = b
	b = c
}

func SwapTest() {
	a := 10
	b := 20
	swap(a, b)
	println(a, b)
	swaps(&a, &b)
	println(a, b)
}

//引用传递的方法
func swaps(a, b *int) {
	var c int
	c = *a
	*a = *b
	*b = c
}

func NumTest() {
	//函数作为值
	num := func(x int) int {
		return x + 3
	}

	println(nums(1) + num(5))
}

func nums(x int) int {
	return x + 5
}

func add(x, y int) func(m, n int) (int, int) {
	var i int
	return func(m, n int) (int, int) {
		i++
		return i + 2, i + 3
	}
}

func adds(x, y int) func(m, n int) (int, int) {
	var i int
	return func(m, n int) (int, int) {
		i++
		return x + m + i, y + n + i
	}
}

func AddTest() {
	add_s := adds(4, 5)
	fmt.Println(add_s(1, 2))
	fmt.Println(add_s(1, 2))
	fmt.Println(add_s(1, 2))
}

//求圆直径的方法
func (c Circle) getDiam() float64 {
	return 2 * c.radius
}

func MethodTest() {
	var c Circle
	c.radius = 10
	fmt.Println(c.getDiam())
}

func MethodMore(a ...int) {
	MethodArrays(a)
}

func MethodArray() {
	var arr = [] int{9, 8, 7, 6, 5, 4}
	//
	MethodMore(arr...)
}

func MethodArrays(a []int) {
	for k, v := range a {
		fmt.Println(k, "-----", v)
	}
}

func noNameMethad() func() {
	var x int
	return func() {
		x = x + 10
		fmt.Println(x)
	}
}

func Addt() {
	var a = noNameMethad()
	a()
	a()
}

func Factory(types string) func(name string) {
	return func(name string) {
		fmt.Println(name + "." + types)
	}
}
