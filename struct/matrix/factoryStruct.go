package matrix

import (
	"fmt"
	"unsafe"
)

type File struct {
	fd   int
	name string
}

type matrix struct {
	aaa int
}

func main() {
	MyFile := NewFile(10194, "xiongyi.log")
	//%+v变量为结构体时展示字段名
	fmt.Printf("文件描述符：%v,文件名：%#v,Test：%+v\n", MyFile.fd, MyFile.name, *MyFile)
	//一个结构体实例占用了多少内存？
	fmt.Printf("此结构体实例占用的内存为：%v\n", unsafe.Sizeof(*MyFile))
	//wrong := new(matrix)
	//fmt.Println(wrong)
	//fmt.Println(Newmatrix())
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

// Newmatrix 强制使用工厂方法，将结构体变为私有的
func Newmatrix() *matrix {
	m := new(matrix)
	return m
}
