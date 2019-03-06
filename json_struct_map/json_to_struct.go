package main

import (
	"fmt"
	"encoding/json"
)

//注意json里面的key和struct里面的key要一致，struct中的key的首字母必须大写，而json中大小写都可以。
type Student struct {
	S_num int
	S_name string
	S_addr string
}

type studentSlice struct {
	Student []Student
}

//var s Student
var ss studentSlice

func JsonToStruct(str string) error  {
	err := json.Unmarshal([]byte(str), &ss)
	return err
}

func main()  {
	// int 类型的不能加双引号 否则结果为0
	str := `{"student":[{"S_num":2,"S_name":"martin","S_addr":"martin_addr"},{"S_num":1,"S_name":"allen","S_addr":"allen_addr"}]}`
	//str := `{"S_num":2,"S_name":"martin","S_addr":"martin_addr"}`
	err := JsonToStruct(str)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json decode fail")
	}

	fmt.Println(ss);
}