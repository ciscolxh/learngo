# 环境变量
搭建环境我们首先要把环境变量配置好

>> GOPATH存的是项目的工作路径，项目下有有这么几个子目录

* src存的是项目的源码文件，包括我们下载的导包都在这里
* bin目录 通过go install 编译出的二进制可执行文件放在这里
* pkg 放的是编译后的源文件（.a）


 GOROOT
>> GOROOT 指向的是官方的安装路径


# 基础语法

### 变量

1. 使用关键字var定义
2. 使用:=定义
3. 让编译器自动判断类型

###### 使用关键字var定义
```
func keywords(){
    var name string = "xiaohei"
    var age int = 18
    fmt.PrintLn(name,age)
}

func keywords(){
    var name  = "xiaohei"
    var age  = 18
    fmt.PrintLn(name,age)
}

func keywords(){
    var (
        name = "xiaohei"
        age = 18
        ）
    fmt.PrintLn(name,age)
}
```

###### 使用:=定义

```
func keywordToOmit(){
    name := "小黑"
    age := 18
    fmt.Println(name,age)
}
```

###### 编译器自动判断类型

```
func keywordType(){
   var name,age = "小黑",18
   fmt.Println(name,age)
}
```



```
package main
import "fmt"

funt main(){
    keywords()
    keywordToOmit()
    keywordType()
}
```

### 变量的作用域
1. 局部变量

> 局部变量写在函数内，作用域只是在函数内，出了函数就没有效果

2. 全局变量

> 全局变量写在函数外，作用在整个包内,相同包名可以随意使用

3. 形式参数

> 函数传入参数，作用在函数内部


### 常量

1. 用修饰词const修饰

```
    /*
      一般我们定义常量也和其他的变量名一至
      因为在Go语言中首字母大写代表全局变量
    */
    const p = 3.14
    fmt.Println(p)
```

### 枚举
* 枚举我们一般也用const来定义

```
    const {
        red = iota
        green
        black
        blue
    }
    fmt.Println(red,green,black,blue)
    //输入 0 1 2，3
```

### 条件语句

* 在Go语言中条件语句和我们其他的语言的写法比较相似，例如我用的是Java语言，只是取消掉了Java中if后面的括号来写

```
    func ifTest(age int){
        if age<18 {
            fmt.Println("未成年")
        }else{
            fmt.Println("成年")
        }
    }
```

### switch语句
* switch 可以像在Java中一样使用，但是Golang中的switch在执行完case之后会自动break，不需要我们手动加break
* 也可以在case里做比较操作

###### 用法一：

```
    func switchTest(a int,b int ,c string){
        switch c {
            case "+":
                fmt.Println(a+b)
            case "-":
                fmt.Println(a-b)
            case "*":
                fmt.Println(a*b)
            case "/":
                fmt.Println(a/b)
            default:
                fmt.Println("输入有误")
        }
    }
```

###### 用法二

```
func SwitchTest(grade float32) {
	switch {
	case grade >= 0 && grade < 60:
		fmt.Println("E")
	case grade >= 60 && grade < 70:
		fmt.Println("D")
	case grade > 70 && grade < 80:
		fmt.Println("C")
	case grade >= 80 && grade < 90:
		fmt.Println("B")
	case grade >= 90 && grade <= 100:
		fmt.Println("A")
	default:
		fmt.Println("输入错误")
	}
}
```

### 循环语句

* 在Go语言中循环语句和其他语言也是类似，只是在for后面少了一个括号

```
   func forTest()  {
	    for i:=0;i<10 ;i++  {
		    fmt.Println(i)
    	}
    }
```

* For循环的死循环写法

```
    func ForMore() {
	    for   {
		    fmt.Println("死循环")
	    }   
    }
```

* 在Go语言中没有while循环，可以用for替代while

```
    func forLoop(){
        i:=0;
        //满足条件执行循环内的语句
        for i<10{
            fmt.Println(i);
            i++
        }
    }
```
* for range 类似于Java中的增强for循环

```
    func arrayTest(){
        var array = [] int{1,2,3,4,5}
        //
        for key,value := range array{
            fmt.Println("索引为：",key)
            fmt.Println("值为：",value)
        }
    }
```

### 循环语句中的控制语句
 
>> 
1. continue 调出本次循环执行下一次循环
    ```
        for i:=0; i<10; i++{
            if i==3{
                //跳过本次循环
                continue
            }
            fmt.Print("i")
        }
        //输出 012456789
    ```
2. break 跳出for循环
    ```
    func test(){
        for i:=0;i<10;i++{
            if i==3 {
                //调出整个循环
                break
            }
            fmt.Print(i)
        }
        // break跳出循环会执行这里
        fmt.Print("执行结束")
    }
    ```
3. return 跳出函数
     ```
    func test(){
        for i:=0;i<10;i++{
            if i==3 {
                //跳出函数
                return
            }
            fmt.Print(i)
        }
        // return 跳出整个函数不会执行这里
        fmt.Print("执行结束")
    }
    ```

### 函数

* 函数基本的语法
    ```
        /*
            关键字func 后面加函数名 
            函数名首字母大写代表全局方法
            传入参数可以是多个用，隔开
            返回参数可以是多个用，隔开
        */
        func funcName (接收参数...) （返回参数...）{
            
        }
    ```
    
    例如：
    
    ```
        //求两个数的最大值
        func funcTest(a, b int) int {
        	if a > b {
	        	return a
        	} else {
	        	return b
    	    }
        }
    ```
    返回多个函数：
    ```
        func funcMoreReturn(a,b int) (int,int){
            return a,b
        }
    ```
    
###### 函数传递类型
1. 值传递
 
```
    func main(){
        a:= 10
        b:= 20
        //执行值传递方法
        swap(a,b)
        //传递后a,b 的值不变 
        fmt.Println(a,b)
    }
    
    func swap(a,b int){
        var c int 
        c = a
        a = b
        b = c
    }
```

2. 引用传递

```
     func main(){
        a:= 10
        b:= 20
        //执行引用传递方法
        swap(&a,&b)
        //传递后a,b 的值交换 
        fmt.Println(a,b)
    }
    
    func swap(a,b *int){
        var c int 
        c = *a
        *a = *b
        *b = c
    }
```

###### 函数用法
1. 函数作为值来使用
```
    func NumTest() {
    	//函数作为值
    	num := func(x int) int {
    		return x + 3
    	}
    	println(num(5))
    }
```

2. 闭包函数
    
```
    func Factory(types string) func(name string) {
	return func(name string) {
		fmt.Println(name + "." + types)
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
}
```
3.使用闭包函数调试

```
    where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
```

4. 方法

###### 格式
```
    //func (实体类) 方法名(){
        //这里执行函数方法
    }
```

示例

```
    //定义一个圆的实体类
    type Circle struct {
	    radius float64 `json:"radius"`
    }
    
    //给实体类做一个计算直径的方法
    func(c Circle) getDiam() float64{
        return 2 * c.radius
    }
    
    //使用实体类的方法
    func test(){
        //初始化一个圆
        var c Circle
        //给圆的半径赋值
        c.radius = 10
        //打印出圆的直径
        fmt.Println(c.getDiam())
    }
```

###### 传递长参数
长参数可以通过...TYPE转为数组类型
```
    func arrayType(nums ...int){
        var arr [] int = strings   
    }
```
数组可以通过ARRAY...转换为长参数类型

```
    func arrayTypeTest (){
        var arr = [] int {1,2,3,4}
        arrayType(arr...)
    }
```

### 数组
数组是一组相同类型的数据序列

特点：
1. 相同数据类型
2. 在初始化已经知道了长度
3. 可以通过索引存取

*注意：数组的数据类型包含长度[10]int 和 [5]int属于不同的数据类型*

###### 声明格式以及赋值

###### 声明
1. 先申明再赋值
```
    var arrayList [10] int
    
    for i:=0;i<len(arrayList);i++{
        arrayList[i] = i*i
    }
    
    for i:=0;i<len(arrayList);i++{
        fmt.Println(arrayList[i])
    }
```

2. new 一个数组然后赋值
```
    var arrayList = new ([10] int)
     for i:=0;i<len(arrayList);i++{
        arrayList[i] = i*i
    }
    
    for i:=0;i<len(arrayList);i++{
        fmt.Println(arrayList[i])
    }
```

3. 直接赋值
```
    var arrayList = [10] int{0,1,4,9,16,25,36,49,64,81}
    for _,value := range arrayList{
         fmt.Println(value)
    }
```

4. 赋值数组可以用...代替数组的指定长度
```
    var arrayList = [...] int {0,1,4,9,16,25,36,49,64,81}
    for _,value := range arrayList{
         fmt.Println(value)
    }
```

5. 赋值数组...也可以省略
```
    var arrayList = [] int {0,1,4,9,16,25,36,49,64,81}
    for _,value := range arrayList{
         fmt.Println(value)
    }
```

6. 初始化数组赋值部分

```
    var arrayList = [10] int {0,1,4}
    //数组的长度是10， 后面没有赋值的打印出默认为0
    for _,value := range arrayList{
         fmt.Println(value)
    }
```

7. 初始化数组并指定索引赋值
```
    //在索引为3的位置赋值9，在索引为5的位置赋值25
    var arrayList = [10] int {3:9,5:25}
    for _,value := range arrayList {
        fmt.Println(value)
    }
```

### 切片
切片是数组中连续的一部分，最短可以为0，最长为数字长度，切片可以理解为一个长度可变的数组

> 计算数组外的长度使用cap(s),也就是所在数组整个数组的长度

```
    0<=len(caps)<=cap(caps)    
```

map是一种key，value的数据结构，在其他语言中也是存在的

### Map

###### 创建map的方式

1. 姿势一先用make创建然后进行初始化

```
    map1 := make(map[string] string)
    map1["name"] = "xiaohei"
    map1["age"] = "26"
    map1["sex"] = "man"
```

2. 姿势二，创建并且初始化

```
    map2 := map[string] string{"name":"xiaohei","age":"26","sex":"man"}
```

###### 对map进行遍历

* map的遍历和数组比较相似  直接可以用range进行遍历，取出key和value

```
    map1 := map[string] string {"name":"xiaohei","age":"26","sex":"man"}
    for key,value := range map1{
        fmt.Println("键为:"+key)
        fmt.Println("值为:"+value)
    }
```

###### 判断map中是否存在某个值

* 通过map[key]取可以拿到两个值，第一个值为value，
  如果我们只是通过value是否为空不能够判断出value本身就是空串还是
  key不存在，因此我们要用第二个值判断这个key是否存在

```
    map1 := map[string] string {"name":"xiaohei","age":"26","sex":"man"}
    if value,ok := map1["name"];ok {
        fmt.Println("存在name的字段，值为"+value)
    } else{
        fmt.Println("不存在name的字段！")
    }
```

###### 删除map中的某一个key（字段）

```
    map1:=map[string] string {"name":"xiaohei","age":"26","sex":"man"}
    //判断是否存在name字段
    if _,ok:=map1["name"];ok{
        fmt.Println("存在name字段")//打印存在name字段
    }else{
        fmt.Println("不存在那么字段")
    }
    //删除name字段
    delete(map1,"name")
    //判断是否存在name字段
    if _,ok :=map1["name"] ;ok{
        fmt.Println("存在name字段")
    }else{
        fmt.Println("不存在name字段")//打印不存在name字段
    }
```

