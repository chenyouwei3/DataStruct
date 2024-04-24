package main

import (
	"fmt"
	"reflect"
)

func main() {
	var mike interface{}
	fmt.Println(reflect.TypeOf(mike))
}
