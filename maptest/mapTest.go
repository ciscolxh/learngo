package maptest

import (
	"fmt"
	"sort"
)

/*
 * 创建Map的两种姿势
 */
func MapCreate() {
	//方法一

	fmt.Println("------------方法一---------")
	var map1 = make(map[string]string)
	map1["name"] = "xiaohei"
	map1["age"] = "26"
	map1["sex"] = "nan"

	for key, value := range map1 {
		fmt.Println(key + "为" + value)
	}

	//方法二
	fmt.Println("------------方法二---------")
	var map2 = map[string]string{"name": "xiaohei", "age": "26", "sex": "nan"}
	for key, value := range map2 {
		fmt.Println(key + "为" + value)
	}

}

/*
 * 判断Map是否存在某个键
 */
func MapJudge() {
	m := map[string]string{"name": "小红帽", "age": "7", "sex": "nv"}
	//第一个参数为键对应的值，第二个参数为，是否存在这个键
	value, have := m["age"]
	if have {
		fmt.Println("存在age字段,age的值为" + value)
	} else {
		fmt.Println("不存在age")
	}

	//简略写法
	if v, ok := m["name"]; ok {
		fmt.Println("存在name字段,name的值为" + v)
	} else {
		fmt.Println("不存在name字段")
	}

}

//删除某个Key
func MapDeleteKey() {
	m := map[string]string{"name": "小红帽", "age": "7", "sex": "nv"}
	//删除name字段
	delete(m, "name")
	if v, ok := m["name"]; ok {
		fmt.Println("存在name字段,name的值为" + v)
	} else {
		fmt.Println("不存在name字段")
	}
}

/*
 * map排序
 * 排序是新建一个切片，
 * 把key都放进去，
 * 然后对key进行排序
 * 通过key获取Value
 */
func MapSort() {
	m := map[string]int{"d": 6, "e": 5, "f": 4, "g": 3, "h": 2, "i": 1, "a": 9, "b": 8, "c": 7}

	arr := make([]string, len(m))
	i := 0
	for key := range m {
		arr[i] = key
		i++
	}

	sort.Strings(arr)

	for _, value := range arr {
		fmt.Println(value+"=", m[value])
	}
}

func Tests() {
	map1:=map[string] string {"name":"xiaohei","age":"26","sex":"man"}
	if _,ok:=map1["name"];ok{
		fmt.Println("存在name字段")
	}else{
		fmt.Println("不存在那么字段")
	}

	delete(map1,"name")
	if _,ok :=map1["name"] ;ok{
		fmt.Println("存在name字段")
	}else{
		fmt.Println("不存在name字段")
	}
}
