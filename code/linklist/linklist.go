package linklist

import (
	"errors"
	"strconv"
)

// LNode 单链表的结节结构体
type LNode struct {
	// 数据域
	element interface{}
	// 下一个节点的指针
	next *LNode
}

// LinkList 单链表结构体
type LinkList struct {
	// 单链表长度
	len int
	// 单链表头结点
	head *LNode
}

// NewLinkList 单链表初始化
func NewLinkList() *LinkList {
	list := new(LinkList)
	list.len = 0
	list.head = new(LNode)
	return list
}

// ----------- 实现 List 的接口 -------------

// Size 数组的大小
func (list *LinkList) Size() int {
	return list.len
}

// Get 获取第几个元素
func (list *LinkList) Get(index int) (interface{}, error) {
	if index > list.len {
		return nil, errors.New("访问越界")
	}
	p := list.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p.element, nil
}

// Set 修改指定位置的数据
func (list *LinkList) Set(index int, newval interface{}) error {
	if index > list.len {
		return errors.New("访问越界")
	}
	p := list.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	p.element = newval
	return nil
}

// Insert 插入数据
// 插入位置从 1 开始
func (list *LinkList) Insert(index int, newVal interface{}) error {
	if index > list.len || index <= 0 {
		return errors.New("访问越界")
	}
	// 插入位置的前一个结点：前驱
	p := list.head
	for i := 1; i < index; i++ {
		p = p.next
	}
	node := new(LNode)
	node.next = p.next
	node.element = newVal
	p.next = node
	list.len++
	return nil
}

// Append 追加数据
func (list *LinkList) Append(newval interface{}) {
	node := new(LNode)
	node.element = newval
	p := list.head
	for i := 1; i <= list.len; i++ {
		p = p.next
	}
	p.next = node
	list.len++
}

// Clear 清空链表
func (list *LinkList) Clear() {
	list.head.next = nil
	list.len = 0
}

// Delete 删除指定位置数据
func (list *LinkList) Delete(index int) error {
	if index > list.len || index <= 0 {
		return errors.New("访问越界")
	}
	// 删除位置的前一个结点：前驱
	p := list.head
	for i := 1; i < index; i++ {
		p = p.next
	}
	p.next = p.next.next
	list.len--
	return nil
}

// String 返回链表的字符串
func (list *LinkList) String() string {
	p := list.head
	str := ""
	for p.next != nil {
		p = p.next
		str = str + strconv.Itoa(p.element.(int))
	}
	return str
}
