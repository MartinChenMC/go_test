package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

//Color 声明Color为byte的别名
type Color byte

//Box type box
type Box struct {
	len, wid, high float64
	c Color
}

//BoxList box集合，不确定个数，使用动态数组
type BoxList []Box

//Volume 计算box的容量
func (b Box) Volume() float64 {
	return b.len*b.wid*b.high
}

//SetColor 将box的颜色设为特定值
func (b Box) SetColor(c Color) {
	b.c = c
}

//BiggestColor 找boxlist中容量最大的box的颜色
func (boxlist BoxList) BiggestColor() (c Color) {
	v := 0.0
	c = WHITE

	//遍历boxlist，求容量最大的
	for _, b := range boxlist {
		if tmpv := b.Volume(); tmpv > v {
			v = tmpv
			c = b.c
		}
	}

	return
}

//SetBlack 将boxlist中的box颜色都设置为BLACK
func (boxlist BoxList) SetBlack() {
	for i := range boxlist {
		boxlist[i].c = BLACK
	}
}

//PrintBox 打印box
func (b Box) PrintBox() {
	fmt.Println(b.len, b.wid, b.high, b.Volume(), b.c.String())
}

//String 将颜色以字符串形式返回
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	//创建一个boxlist，存储若干box
	boxlist := BoxList{
		Box{1,2,3,RED},
		Box{1,4,2,WHITE},
		Box{4,2,5,BLACK},
		Box{2,4,3,BLUE},
	}

	//打印所有的box
	for i := range boxlist {
		boxlist[i].PrintBox()
	}

	//输出boxlist的大小
	fmt.Println("len of boxlist = ", len(boxlist))

	//输出第二个box的颜色
	fmt.Println("color of boxlist[1] = ", boxlist[1].c.String())

	//输出boxlist中容量最大的box的颜色
	fmt.Println("biggest color = ", boxlist.BiggestColor())

	//将boxlist中的所有box颜色都设置为黑色
	boxlist.SetBlack()
	fmt.Println("将所有box颜色设置为黑色")
	for i := range boxlist {
		boxlist[i].PrintBox()
	}
}