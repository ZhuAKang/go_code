package heap

import (
	"errors"
)

// Heap 堆接口，实现了就具有了堆的功能
type Heap interface {
	// ShiftUp 节点比它的父节点大（最大堆）或者小（最小堆）
	// 那么需要将它同父节点交换位置。这样是这个节点在数组的位置上升。
	ShiftUp(int)
	// ShiftDown 节点比它的子节点小（最大堆）或者大（最小堆），那么需要将它向下移动。
	ShiftDown(int)
	// Insert 在堆的尾部添加一个新的元素，然后使用 shiftUp 来修复堆。
	Insert(interface{}) error
	// Remove 移除并返回最大值（最大堆）或者最小值（最小堆）
	// 为了将这个节点删除后的空位填补上，需要将最后一个元素移到根节点的位置，然后使用 shiftDown 方法来修复堆。
	Remove() (interface{}, error)
	// RemoveIndex 和 remove() 一样，差别在于可以移除堆中任意节点，而不仅仅是根节点。
	// 当它与子节点比较位置不是无序时使用 shiftDown()，如果与父节点比较发现无序则使用 shiftUp()。
	RemoveIndex(int) (interface{}, error)
	// 将一个更小的值（最小堆）或者更大的值（最大堆）赋值给一个节点。
	// 由于这个操作破坏了堆属性，所以需要使用 shiftUp() 来修复堆属性。
	Replace(int, interface{}) error
	// Search 在堆中查找有无此元素
	Search(interface{}) (int, bool)
	// TopHeap 获取当前堆顶元素的值
	TopHeap() (interface{}, error)
}

// AHeap 使用数组实现的堆的数据结构
// 这里以最大堆为例（孩子节点的值一直小于等于父亲节点的值）
type AHeap struct {
	// 存储元素
	elements []int
	// 容量
	cap int
	// 堆中元素个数
	size int
}

// NewAHeap 新建一个堆
func NewAHeap(cap int) *AHeap {
	heap := new(AHeap)
	heap.cap = cap
	heap.size = 0
	heap.elements = make([]int, cap, cap)
	return heap
}

// ShiftUp 结点位置上提
// 节点位置从 1 开始
func (heap *AHeap) ShiftUp(index int) {
	// 当前节点是不是存在父节点，没有父节点就不用上移了
	// 有父节点，如果有父节点且父节点的值小于当前节点，则将此节点上移，父节点下移
	// 一直下去（递归）直到没有父节点或者值小于了父节点
	fa := (index + 1) / 2
	// 父节点存在且当前节点大于父节点
	for fa >= 1 && heap.elements[index-1] > heap.elements[fa-1] {
		heap.elements[index-1], heap.elements[fa-1] = heap.elements[fa-1], heap.elements[index-1]
		index = fa
		fa = (index + 1) / 2
	}
}

// ShiftDown 结点位置下移
// 节点位置从 1 开始
func (heap *AHeap) ShiftDown(index int) {
	// 如果没有孩子，退出
	// 有孩子：
	//        有两个孩子，元素值小于孩子节点的值，则选择两个孩子中较小的那一个与其交换
	//        只有一个左孩子，和左孩子交换
	// 继续循环

	child1 := 2 * index
	// index 的孩子存在
	for child1 < heap.size {
		// 有两个孩子
		if child1+1 <= heap.size {
			if heap.elements[child1-1] < heap.elements[child1] {
				// 和右孩子交换
				heap.elements[index-1], heap.elements[child1] = heap.elements[child1], heap.elements[index-1]
				index = child1 + 1
				child1 = 2 * index
			} else {
				// 和左孩子交换
				heap.elements[index-1], heap.elements[child1-1] = heap.elements[child1-1], heap.elements[index-1]
				index = child1
				child1 = 2 * index
			}
		} else {
			// 只有一个孩子，那就是 child1，交换之后直接可以退出了
			heap.elements[index-1], heap.elements[child1-1] = heap.elements[child1-1], heap.elements[index-1]
			return
		}
	}
}

// Insert 在堆的尾部添加一个新的元素，然后使用 shiftUp 来修复堆。
func (heap *AHeap) Insert(value int) error {
	if heap.size >= heap.cap {
		return errors.New("堆已满，不能继续插入")
	}
	heap.elements[heap.size] = value
	heap.size++
	heap.ShiftUp(heap.size)
	return nil
}

// Remove 移除并返回最大值（最大堆）或者最小值（最小堆）
func (heap *AHeap) Remove() (interface{}, error) {
	if heap.size == 0 {
		return nil, errors.New("堆为空，无法移除")
	}
	// 取出堆顶元素，将最后一个元素移到根节点的位置，然后使用 shiftDown 方法来修复堆
	value := heap.elements[0]
	heap.elements[0] = heap.elements[heap.size-1]
	heap.size--
	heap.ShiftDown(1)
	return value, nil
}

// RemoveIndex 和 remove() 一样，差别在于可以移除堆中任意节点，而不仅仅是根节点。
// 当它与子节点比较位置不是无序时使用 shiftDown()，如果与父节点比较发现无序则使用 shiftUp()。
// 节点位置从 1 开始
func (heap *AHeap) RemoveIndex(index int) (interface{}, error) {
	if index > heap.size {
		return nil, errors.New("超出堆索引")
	}
	// 删除的是堆顶
	if index == 1 {
		return heap.Remove()
	}
	// 删除的是堆的最后一个元素，那就直接删除
	if index == heap.size {
		value := heap.elements[index-1]
		heap.size--
		return value, nil
	}
	// 删除的是堆中间的元素，则要把堆的最后一个元素提到这儿来然后进行堆得调整
	value := heap.elements[index-1]
	heap.elements[index-1] = heap.elements[heap.size-1]
	heap.size--
	// 如果与父节点比较发现无序则使用 shiftUp()
	if heap.elements[index/2-1] < heap.elements[index-1] {
		heap.ShiftUp(index)
	}
	// 与子节点比较，无序则使用 shiftDown() ，这里也要分有一个还是有两个子节点
	// 有两个
	if index*2+1 <= heap.size {
		if heap.elements[2*index-1] > heap.elements[index-1] || heap.elements[2*index+1] > heap.elements[index-1] {
			heap.ShiftDown(index)
		}
		// 有一个
	} else if index*2 == heap.size {
		if heap.elements[2*index-1] > heap.elements[index-1] {
			heap.ShiftDown(index)
		}
	}
	// 没有孩子就不用再做别的了，直接返回
	return value, nil
}

// Replace 将一个更小的值（最小堆）或者更大的值（最大堆）赋值给一个节点。
// 由于这个操作破坏了堆属性，所以需要使用 shiftUp() 来修复堆属性。
func (heap *AHeap) Replace(index int, value int) error {
	if index > heap.size {
		return errors.New("超出堆索引")
	}
	// 开始替换
	if heap.elements[index-1] > value {
		// 用了个小的替换，则需要向下调整
		heap.ShiftDown(index)
	} else {
		// 用了个大的替换，则要向上调整
		heap.ShiftUp(index)
	}
	return nil
}

// Search 在堆中查找有无此元素，返回查找到的元素的地址和 true/false
func (heap *AHeap) Search(value int) (int, bool) {
	// 堆是空的或者查找的元素比最大堆的堆顶元素都大，那就查不到
	if heap.size == 0 || value > heap.elements[0] {
		return 0, false
	}
	// 就遍历一次搜索吧，不考虑太复杂的了
	for i := 0; i < heap.size; i++ {
		// 比最大的元素都大，那下面肯定没有了
		if heap.elements[i] == value {
			return i, true
		}
	}
	return 0, false
}

// TopHeap 获取当前堆顶元素的值
func (heap *AHeap) TopHeap() (int, error) {
	if heap.size == 0 {
		return 0, errors.New("堆为空，不存在堆顶元素")
	}
	return heap.elements[0], nil
}
