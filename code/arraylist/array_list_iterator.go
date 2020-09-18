package arraylist

import (
	"errors"
)

// Iterator 迭代器的接口
type Iterator interface {
	HasNext() bool              // 是否有下一个元素
	Next() (interface{}, error) // 下一个元素
	Remove()                    // 删除
	GetIndex() int              // 得到当前索引
}

// Iterable 创建定义迭代器的接口
type Iterable interface {
	Iterator() Iterator // 构造初始化接口
}

// ArrayListIterator 构造指针访问数组
type ArrayListIterator struct {
	list  *ArrayList // 数组指针
	index int        // 当前访问的索引
}

// Iterator ArrayList实现了 Iterable 接口，可以迭代
func (list *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // 构造迭代器
	it.index = 0
	it.list = list
	return it
}

// HasNext 是否有下一个元素
func (it *ArrayListIterator) HasNext() bool {
	return it.index < it.list.size // 是否有下一个
}

// Next 下一个元素
func (it *ArrayListIterator) Next() (interface{}, error) {
	if !it.HasNext() {
		return nil, errors.New("已经没有下一个元素了")
	}
	value, err := it.list.Get(it.index)
	it.index++
	return value, err
}

// Remove 删除当前迭代器 index 位置的元素
func (it *ArrayListIterator) Remove() {
	it.index--
	it.list.Delete(it.index)
}

// GetIndex 得到当前索引
func (it *ArrayListIterator) GetIndex() int {
	return it.index
}
