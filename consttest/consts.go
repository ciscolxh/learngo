package consttest

import "fmt"

const (
	Red = iota
	Green
	Blue
	Black
	Orange
)

//常量的示例
func Consts () {

	/*在Go语言中一般常量也用常规命名，
	 *因为在Go中大写有其他含义，
	 *代表全局变量
	 */
	const p  = "3.14"
	fmt.Println(p)
	fmt.Println(Red,Green,Blue,Black,Orange)
}

