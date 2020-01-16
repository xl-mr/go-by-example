package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Empty struct{}

func main() {
	byt := []byte(`{"num":"6","strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"]
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(num.(string))

	strs := dat["strs"].([]interface{})
	fmt.Println(strs)

	res, ok := dat["nm"]
	if ok {
		fmt.Println("res: ", res)
	} else {
		fmt.Println("not exists")
	}

	path, _ := os.Getwd()
	fmt.Println(path)

	fmt.Println(reflect.TypeOf(Empty{}).PkgPath())
}
