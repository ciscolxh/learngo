package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strconv"
)


func main() {


	file, err :=os.Open("/test")
	if err!=nil {
		fmt.Println("读取错误")
	}else {
		inputReader := bufio.NewReader(file)
		var i = 1
		for {
			inputString, readerError := inputReader.ReadString('\n')
			if readerError == io.EOF {
				return
			}
			fmt.Print(strconv.Itoa(i) +" ")
			fmt.Println( inputString)
			i++
		}
	}
	defer file.Close()

}
