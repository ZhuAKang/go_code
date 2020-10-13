package sort

import (
	"code/heap"
	"fmt"
)

// SelectSort 这里就简单的模拟一下，注重算法本身
func SelectSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("选择排序前数组：", arr)
	min := arr[0]
	position := 0
	for j := 0; j < len(arr); j++ {
		min = arr[j]
		position = j
		// 找到剩余的里面最小的元素的位置
		for i := j; i < len(arr); i++ {
			if arr[i] <= min {
				min = arr[i]
				position = i
			}
		}
		// 如果剩下的找到了最小的，交换元素
		if j != position {
			arr[j], arr[position] = arr[position], arr[j]
		}

	}
	fmt.Println("选择排序后数组：", arr)
}

// InsertSort 插入排序
func InsertSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("插入排序前数组：", arr)
	// 当前待插入的元素
	position := arr[0]
	var j int
	// 外层循环
	for i := 1; i < len(arr); i++ {
		// 待插入到前面有序部分的元素
		position = arr[i]
		// 内层循环进行插入
		for j = i - 1; j >= 0 && arr[j] > position; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = position
	}
	fmt.Println("插入排序后数组：", arr)
}

// ShellSort 希尔排序
func ShellSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("希尔排序前数组：", arr)
	n := len(arr)
	// 步长的选取
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	// 开始排序，控制步长
	for h >= 1 {
		// 外层循环
		for i := h; i < n; i++ {
			// 内层循环进行插入
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j-h], arr[j] = arr[j], arr[j-h]
			}
		}
		h = h / 3
	}
	fmt.Println("希尔排序后数组：", arr)
}

// BubbleSort 冒泡算法
func BubbleSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("冒泡排序前数组：", arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("冒泡排序后数组：", arr)
}

// MergeSort 归并排序(考虑使用递归来实现)
func MergeSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("归并排序前数组：", arr)
	bufarr := make([]int, len(arr), len(arr))
	MSort(arr[:], 0, len(arr)-1, bufarr)
	fmt.Println("归并排序后数组：", arr)
}

// MSort 归并排序
func MSort(arr []int, start int, end int, buf []int) {
	// 如果元素数量超过2个，那就在继续二分
	if start+1 < end {
		// 二分
		mid := start + (end-start)/2
		// 递归
		MSort(arr, start, mid, buf)
		MSort(arr, mid+1, end, buf)
		// 合并两个有序的数组
		Merge(arr, start, end, mid, buf)
		// 2个或者1个，只有一种条件就是左面的比右边的大才需要交换，其他条件都不需要操作
	} else if arr[start] > arr[end] {
		arr[start], arr[end] = arr[end], arr[start]
	}
}

// Merge 合并有序数组
func Merge(data []int, start int, end int, mid int, buf []int) {
	// index是合并后数组的索引，left是左半部分有序数组的索引，right是右半部分有序数组的索引
	index, left, right := start, start, mid+1
	for left != start+1 && right != end+1 {
		if data[left] > data[right] {
			buf[index] = data[right]
			index++
			right++
		} else {
			buf[index] = data[left]
			left++
			index++
		}
	}
	// 把剩下的左半部分合并到缓冲中
	for left != mid+1 {
		buf[index] = data[left]
		left++
		index++
	}
	// 把剩下的右半部分合并到缓冲中
	for right != end+1 {
		buf[index] = data[right]
		right++
		index++
	}
	// 把排好序的缓冲中的数据拷贝到数据内存中
	for i := start; i <= end; i++ {
		data[i] = buf[i]
	}
}

// QSort 快速排序
func QSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("快速排序前数组：", arr)
	QuickSort(arr[:])
	fmt.Println("快速排序后数组：", arr)

}

// QuickSort 快排的函数（递归）入口
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	base := arr[0] // 初始基准元素设为第一个
	left := 0
	right := len(arr) - 1 // 右边界
	// 进入循环
	for left < right {
		// 由于左边的(第0个)被取走当成基准值，所以在左边就留下一个洞，用于存储比基准小的值
		// 所以先从右边找，找到第一个比基准值小的
		for left < right && arr[right] >= base {
			right--
		}
		// 找到了比基准值小的，那就把这个值填入左边的洞，做索引要++
		if left < right {
			arr[left] = arr[right]
			left++
		}
		// 因为上面的操作让右边留了一个洞，开始从左边找比基准值大的
		for left < right && arr[left] <= base {
			left++
		}
		// 找到比基准值大的再次把右边洞填上，并且在左边又留了一个洞
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}

	// 把基准值填入到洞中，这就是本应该他所在的位置
	arr[left] = base
	// 继续分两组递归排序，注意此时基准值已经不用参与排序了
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

// ODSort 奇偶排序
func ODSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("奇偶排序前数组：", arr)
	OddEvenSort(arr[:])
	fmt.Println("奇偶排序后数组：", arr)
}

// OddEvenSort 奇偶排序
func OddEvenSort(arr []int) {
	// 一个标志位，表示是否已经排序过了
	sorted := false
	// 开始循环
	for !sorted {
		sorted = true
		// 偶排序
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		// 奇排序
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}

// HeapSort 堆排序
func HeapSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("堆排序前数组：", arr)
	heap := heap.NewAHeap(12)
	for i := 0; i < len(arr); i++ {
		heap.Insert(arr[i])
	}
	for i := 0; i < len(arr); i++ {
		value, _ := heap.Remove()
		arr[i] = value.(int)
	}
	fmt.Println("堆排序后数组：", arr)
}

// TournamentSort 锦标赛排序
func TournamentSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("锦标赛排序前数组：", arr)
	// 调用函数进行排序
	newarr := tournamentSort(arr[:])
	fmt.Println("锦标赛排序后数组：", newarr)
}

// tournamentSort 锦标赛排序
func tournamentSort(arr []int) []int {
	// 计算存储胜者树所需要的数组大小
	length := 2
	for length < len(arr) {
		length *= 2
	}
	length *= 2
	length--
	result := make([]int, len(arr), len(arr))
	// 树结构存储数据
	data := make([]int, length, length)
	// 存储指示位，指示这个位置对应的数据的下标
	// 如果是叶子节点，那代表的就是在 data 数组中的下标
	// 如果是非叶子节点，代表的就是孩子节点中较小的那个节点的数据（叶子）下标
	//（已经取出的，不需要参与比较，将其值为 -1）
	index := make([]int, length, length)
	// 将待排序的 arr 转存入 data 数组中（树中的叶子节点位置）
	for i := 0; i <= (length-1)/2; i++ {
		if i < len(arr) {
			// 填入数据
			data[i+(length-1)/2] = arr[i]
			index[i+(length-1)/2] = i + (length-1)/2
		} else {
			// arr 的数据已经复制完毕但是有部分没填满
			data[i+(length-1)/2] = 0
			index[i+(length-1)/2] = -1
		}
	}
	// fmt.Println("转移后的数据：", data)
	// fmt.Println("是否为无穷：", index)
	// 开始调用构造函数，构造树（对数组前半部份的元素进行写入/修改）
	// 也即将这个数组修改为一棵树
	// 构造函数在构造的时候，数组的前半部分不放元素值，而是放叶子节点在数组的下标（构造的时候要依据是否为无穷的数组）
	for i := 0; i < len(arr); i++ {
		// 调用构造函数构造出树
		reshape(data[:], index[:])
		// fmt.Println("一次构造后的数组：", data)
		// fmt.Println("一次构造后是否为无穷：", index)
		// 取出树上的根节点并把数值所在的叶子节点的 index 对应的置为无穷
		index[index[0]] = -1
		// 放入 result 数组内
		result[i] = data[0]
	}
	return result
}

// reshape 构造出一棵胜者树的过程
func reshape(data []int, index []int) {
	// 获取数组长度
	length := len(data)
	// 循环写入数组的前半部份
	for i := length/2 - 1; i >= 0; i-- {
		// 如果叶子都是无穷大（即叶子对应的 index 值都为 -1）则 data 此位置无意义，index此位置置为 -1
		if index[2*i+1] == -1 && index[2*i+2] == -1 {
			index[i] = -1
		} else if index[2*i+1] == -1 && index[2*i+2] != -1 {
			// 左孩子提出来过了（无穷大），右孩子存在
			index[i] = index[2*i+2]
			data[i] = data[2*i+2]
		} else if index[2*i+1] != -1 && index[2*i+2] == -1 {
			// 右孩子提出来过了无穷大），左孩子存在
			index[i] = index[2*i+1]
			data[i] = data[2*i+1]
		} else {
			// 左右都存在。哪个小取哪个(一样的话先取左边)
			if data[2*i+1] <= data[2*i+2] {
				index[i] = index[2*i+1]
				data[i] = data[2*i+1]
			} else {
				index[i] = index[2*i+2]
				data[i] = data[2*i+2]
			}
		}
	}
}

// CockTailSort 鸡尾酒排序
func CockTailSort() {
	arr := [11]int{5, 11, 7, 2, 90, 45, 23, 8, 12, 28, 77}
	fmt.Println("鸡尾酒排序前数组：", arr)
	// 两个标志位，标识数组的首与尾
	i, j := 0, len(arr)-1
	// 开始循环
	for i < j {
		// 向大的冒泡，从 i 开始冒到 j 结束，冒泡完 j --
		for k := i; k < j; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
		j--
		if i == j {
			break
		}
		// 向小的冒泡，从 j 到 i 结束，冒泡完 i ++
		for k := j; k >= i; k-- {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
		i++
		if i == j {
			break
		}
	}
	fmt.Println("鸡尾酒排序后数组：", arr)
}
