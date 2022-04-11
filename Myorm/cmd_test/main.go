package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func main() {

	var a struct {
		v int
		d string
	}
	a.v = 1
	a.d = "dasdas"
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.Indirect(reflect.ValueOf(a)))
	fmt.Println(reflect.TypeOf(a).Kind())
	fmt.Println(reflect.ValueOf(a).Type())

}
