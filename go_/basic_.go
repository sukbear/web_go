package main

import (

	"fmt"
)

func main() {
	info := make(map[string]string)
	info["sds"] = "sds"
	info["sd3s"] = "sds"
	fmt.Println(info["sds"])
	delete(info,"sds")
}
