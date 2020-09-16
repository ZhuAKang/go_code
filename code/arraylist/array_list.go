package arraylist

// List （线性表）接口
type List interface {
	Size() int                                  // 数组的大小
	Get(index int) (interface{}, error)         // 获取第几个元素
	Set(index int, newval interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newval interface{}) error            // 追加数据
	Clear()                                     // 清空数组
	Delete(index int) error                     // 删除指定位置数据
	String() string                             // 返回数组字符串
}

// ArrayList 数组的数据结构
type ArrayList struct {
	// 一个泛型数组用来存储数组元素：字符串、整数、实数等
	dataStore []interface{}
	// 存放数组的大小
	size int
}
