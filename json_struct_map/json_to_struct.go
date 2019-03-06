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
var Ss studentSlice
var Sarr studentSlice

func JsonToStruct(str string) error  {
	err := json.Unmarshal([]byte(str), &Ss)
	return err
}

func StructToJson()  {
	Sarr.Student = append(Sarr.Student,Student{S_num:5,S_name:"martin",S_addr:"莱创讯"})
	Sarr.Student = append(Sarr.Student,Student{S_num:6,S_name:"allen",S_addr:"莱创讯allen"})

	jsonStr,err := json.Marshal(Sarr);
	if err != nil {
		fmt.Println("struct to json fail",err)
	}
	fmt.Println("struct to json",string(jsonStr))   //这里用string 强制转换类型，不然都是乱码
}

func main()  {

	StructToJson()

	// int 类型的不能加双引号 否则结果为0
	//str := `{"student":[{"S_num":2,"S_name":"martin","S_addr":"martin_addr"},{"S_num":1,"S_name":"allen","S_addr":"allen_addr"}]}`
	////str := `{"S_num":2,"S_name":"martin","S_addr":"martin_addr"}`
	//err := JsonToStruct(str)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println("json decode fail")
	//}
	//
	//fmt.Println(Ss);
}