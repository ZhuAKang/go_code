package arraylist

import (
	"errors"
	"fmt"
)

// List （线性表）接口
type List interface {
	Size() int                                  // 数组的大小
	Get(index int) (interface{}, error)         // 获取第几个元素
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newval interface{})                  // 追加数据
	Clear()                                     // 清空数组
	Delete(index int) error                     // 删除指定位置数据
	String() string                             // 返回数组字符串
}

// ArrayList 数组的数据结构
type ArrayList struct {
	// 一个泛型切片用来存储数组元素：字符串、整数、实数等
	dataStore []interface{}
	// 存放数组的大小
	size int
}

// NewArrayList 初始化一个数组
func NewArrayList() *ArrayList {
	list := new(ArrayList) // 初始化结构体
	// 这里是使用了 make 创建了一个大小为 0 ，容量为 10 的切片用作存储
	list.dataStore = make([]interface{}, 0, 10) // 开辟10个数组空间
	list.size = 0
	return list
}

// Size 数组的大小
func (list *ArrayList) Size() int {
	return list.size
}

// Get 获取第几个元素
func (list *ArrayList) Get(index int) (interface{}, error) {
	// 索引越界等情况需要返回 error
	if index >= list.size || index < 0 {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

// Set 修改数据
func (list *ArrayList) Set(index int, newval interface{}) error {
	if index >= list.size || index < 0 {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval
	return nil
}

// Insert 插入数据
func (list *ArrayList) Insert(index int, newVal interface{}) error {
	if index >= list.size || index < 0 {
		return errors.New("索引越界")
	}
	// 判断内存的使用 （使用了内置的 cap 函数）
	if list.size == cap(list.dataStore) {
		// 内存满了，申请双倍内存
		// (此时申请的时候一定要初始化新的切片的大小，要不然make申请的时候size为0就不行了，拷贝不进去的)
		new := make([]interface{}, list.size, 2*list.size)
		// 将原切片内容拷贝到新的内存空间内（使用内置的 copy 函数）
		copy(new, list.dataStore)
		// 上述的内容拷贝也可以使用append函数，将源切片的内容全部追加到新的切片中
		// new = append(new, list.dataStore[:]...)
		list.dataStore = new
	}
	// 插入指定位置
	// 使用切片的相关操作来完成，也可以使用 for 循环
	// (下面的这个操作是错误的，因为切片是对底层数组的引用，在别的切片改变了底层数组值的时候，也引用了相应数组部分的切片内容也会发生改变)
	// 所以只有将其拷贝出来而是不单纯地改切片，拷贝出来的话要 make 一个新切片来存，太耗费内存，还是用for吧
	// 先把插入位置 index 之后（包括 index）的所有元素复制出来
	// after := list.dataStore[index:]
	// // 再将 index 之后（包括 index）的元素从源中全部删除
	// list.dataStore = list.dataStore[:index]
	// // 向删除后的切片中追加 新元素
	// list.dataStore = append(list.dataStore, newVal)
	// // 再把之前删除的再放进来
	// list.dataStore = append(list.dataStore, after[:]...)

	// 在使用 for 循环之前，一定要先把切片的内存扩充一个位置，要不然会越界
	// 你可能在想，为什么这儿要是 list.size+1 呢，因为切片切的时候是半闭半开区间，左闭右开
	// 且对一个切片再次进行切片引用，他的引用长度是可以大于源切片长度的
	list.dataStore = list.dataStore[:list.size+1]
	for i := list.size; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newVal
	list.size++
	return nil
}

// Append 追加数据
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval) // 叠加数据
	list.size++
}

// Clear 清空数组
func (list *ArrayList) Clear() {
	// 内存重新分配就行，垃圾回收机制会把原来的内存回收的
	list.dataStore = make([]interface{}, 0, 10) // 开辟10个数组空间
	list.size = 0
}

// Delete 删除指定下标的数据
// index 从 0 开始算起
func (list *ArrayList) Delete(index int) error {
	if index >= list.size || index < 0 {
		return errors.New("索引越界")
	}
	// 删除，就把 index 之后的全部前移就行
	// 下面这个是比较老派的做法
	// for i := index; i < list.size-1; i++ {
	// 	list.dataStore[i] = list.dataStore[i+1]
	// }
	// 可以使用切片中 append 函数的性质
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.size--
	return nil
}

// String 返回数组字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
