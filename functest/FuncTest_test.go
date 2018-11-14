package functest

import (
	"testing"
	"fmt"
	"log"
	"runtime"
)

func TestFuncTest(t *testing.T) {
	fmt.Println(FuncTest(10, 20))
}

func TestFuncMoreReturn(t *testing.T) {
	fmt.Println(FuncMoreReturn("参数一","参数二"))
}

func TestSwapTest(t *testing.T) {
	SwapTest()
}

func TestNumTest(t *testing.T) {
	NumTest()
}

func TestAddTest(t *testing.T) {
	AddTest()
}

func TestMethodTest(t *testing.T) {
	MethodTest()
}

func TestMethodMore(t *testing.T) {
	MethodArray()
	//MethodMore(1,2,3,4,5)
}

func TestAddt(t *testing.T) {
	Addt()
}

//设置一个工厂类
func TestFactory(t *testing.T) {
	var pngFactory = Factory("png")
	var mp3Factory  = Factory("mp3")

	pngFactory("二手玫瑰")
	pngFactory("苏阳")

	mp3Factory("仙儿")
	mp3Factory("命运")
	mp3Factory("贤良")


}
