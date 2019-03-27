package main

import (
	"reflect"
	"fmt"
)

type User_ struct {
	Name   string `user:"name"`
	Passwd string `user:"password"`
}

func main() {
	user := &User_{"ds", "sds"}
	sb := reflect.TypeOf(user).Elem()
	for i := 0; i < sb.NumField(); i++ {
		fmt.Println(sb.Field(i).Tag)
	}
}
