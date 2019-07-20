package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//Slice 切片结构类型
type Slice struct {
	Data unsafe.Pointer //万能指针类型
	Len  int            //数据长度
	Cap  int            //切片容量
}

//Create 创建切片的方法(长度,容量,数据)
func (s *Slice) Create(l int, c int, data ...int) {
	//容错处理
	if s == nil || s.Data == nil || data == nil {
		return
	}
	if len(data) == 0 {
		return
	}
	if l < 0 || c < 0 || l > c || len(data) > l {
		return
	}

	//申请内存空间
	s.Data = C.malloc(C.size_t(c) * 8)

	//初始化长度和容量
	s.Len = l
	s.Cap = c

	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	//遍历data存入申请的内存空间
	for _, v := range data {
		//将数值指针转为地址指针后再获取内存并存值
		*(*int)(unsafe.Pointer(p)) = v
		p += 8
	}

	//释放内存
	//...
}

//Print 打印切片的方法
func (s *Slice) Print() {
	if s == nil || s.Data == nil {
		return
	}
	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)
	//循环打印切片元素
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += 8
	}
	fmt.Println("")
}

//Append 追加切片元素的方法
func (s *Slice) Append(data ...int) {
	if s == nil || s.Data == nil {
		return
	}

	//判断是否需要扩容
	for len(data)+s.Len > s.Cap {
		//将容量扩展为原来的2倍
		s.Data = C.realloc(s.Data, C.size_t(s.Cap)*2*8)
		s.Cap *= 2
	}

	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	// for i := 0; i < s.Len; i++ {
	// 	p += 8
	// }

	//便宜置身到末尾
	p += uintptr(s.Len) * 8

	//遍历data取出数据存入内存
	for _, v := range data {
		*(*int)(unsafe.Pointer(p)) = v
		p += 8
	}
	//修改Len
	s.Len += len(data)

}

//Get 根据下标获取切片元素的方法
func (s *Slice) Get(index int) int {
	if s == nil || s.Data == nil {
		return -1
	}
	if index < 0 || index > s.Len {
		return -1
	}
	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	//偏移指针到index的位置
	p += uintptr(index) * 8

	//取出数据值并返回
	return *(*int)(unsafe.Pointer(p))
}

//Search 查找元素的下标
func (s *Slice) Search(data int) int {
	if s == nil {
		return -1
	}

	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	//对比查找的数据并返回元素下标
	for i := 0; i < s.Len; i++ {
		if *(*int)(unsafe.Pointer(p)) == data {
			return i
		}
		p += 8
	}

	//没有找到数据则返回-1
	return -1
}

//Delete 根据下标删除切片元素的方法
func (s *Slice) Delete(index int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index >= s.Len {
		return
	}

	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	//将指针偏移到index的位置
	p += uintptr(index) * 8

	//记录p指向的后一个元素
	aftp := p

	//循环挪移，完成后一个元素给前一个元素赋值的操作
	for i := index; i < s.Len; i++ {
		aftp += 8 //获得后一个元素的指针
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(aftp))
		p += 8
	}

	//剪掉最后一个元素
	s.Len--

}

//Insert 通过下标插入切片元素的方法
func (s *Slice) Insert(index int, data int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index >= s.Len {
		return
	}
	//如果插入的下标在切片末尾
	if index == s.Len {
		s.Append(data)
		return
	}
	//如果插入位置在切片中间
	//将s.Data指针转换为可以计算的数据值指针
	p := uintptr(s.Data)

	//将指针偏移到index的位置
	p += uintptr(index) * 8

	//插入元素后的最后一个元素位置
	tmp := uintptr(s.Data) + uintptr(s.Len)*8

	//循环将index之后的元素依次后移，实现前一个元素给后一个元素赋值
	for i := s.Len; i < index; i-- {
		*(*int)(unsafe.Pointer(tmp)) = *(*int)(unsafe.Pointer(tmp - 8))
		tmp -= 8
	}
	//将插入的数据替换到p对应位置的值
	*(*int)(unsafe.Pointer(p)) = data

	s.Len++
}

//Destroy 销毁切片的方法
func (s *Slice) Destroy() {
	if s == nil || s.Data == nil {
		return
	}
	C.free(s.Data)
	s.Data = nil //驱使GC工作
	s.Len = 0
	s.Cap = 0
	s = nil
}
