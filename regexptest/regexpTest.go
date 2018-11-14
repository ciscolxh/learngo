package regexptest

import (
	"regexp"
	"fmt"
)

func RegexpTest() {
	s := regexp.QuoteMeta("Hello world. (can you hear me?)")
	fmt.Println(s)


	ok,_:=regexp.Match("/[ 0-9,a-z/]",[]byte("123"))
	fmt.Println(ok)

	t,_ :=regexp.MatchString("/[ 0-9,a-z/]",",./")
	fmt.Println(t)

}
