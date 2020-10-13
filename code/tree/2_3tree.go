package tree

// keyNumber 为 B 树的阶数限制
const keyNumber int = 3

// TTNode 2-3 查找树的节点
type TTNode struct {
	// 因为有时候需要分裂结点，所以 keys 切片最多存放三个 key 值
	// ptrs 切片最多存放 4 个节点指针
	keys      []int     // 元素值
	ptrs      []*TTTree // 子节点指针
	fa        *TTTree   // 父节点
	eleNumber int       // 此节点的元素个数
}

// TTTree 2-3 查找树，也就是 3 阶 B 树
type TTTree struct {
	root *TTNode
}

// InitTree 初始化函数，返回一个 TTTree 的指针
func InitTree() *TTTree {
	tree := new(TTTree)
	return tree
}
