package main

import (
	"regexp"
	"fmt"
)

func isIp(ip string)(b bool){
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}
func main() {
	s:=isIp("127.0.0.1")
	fmt.Println(s)
}
