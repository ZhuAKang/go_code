package tree

import (
	"errors"
)

// RED 表示红节点
const RED bool = true

// BLACK 表示黑节点
const BLACK bool = false

// RBNode 红黑树的节点结构
type RBNode struct {
	key         int     // 键值
	value       int     // 相关联的值
	left, right *RBNode // 左右子树
	N           int     // 这棵子树中的节点总数
	color       bool    // 由其父节点指向它的链接的颜色
}

// RBTree 红黑树的结构体
type RBTree struct {
	root *RBNode
}

// isRed 测试一个节点和其父节点之间链接的颜色
// 是红色就返回 true，黑色就返回 false
func isRed(node *RBNode) bool {
	if node == nil {
		return false
	}
	return node.color == RED
}

// rotateLeft 左旋转，将右红链接旋转成为左红链接
func rotateLeft(node *RBNode) *RBNode {
	temp := node.right
	node.right = temp.left
	temp.left = node
	temp.color = node.color // 延续上一级链接的颜色
	node.color = RED
	temp.N = node.N
	node.N = 1 + size(node.left) + size(node.right)
	return temp
}

// rotateRight 右旋转，将左红链接旋转成为右红链接
func rotateRight(node *RBNode) *RBNode {
	temp := node.left
	node.left = temp.right
	temp.right = node
	temp.color = node.color // 延续上一级链接的颜色
	node.color = RED
	temp.N = node.N
	node.N = 1 + size(node.left) + size(node.right)
	return temp
}

// size 返回以此节点为子树的节点总数
func size(node *RBNode) int {
	if node == nil {
		return 0
	}
	return node.N
}

// IsEmpty 判断树是否为空
func (tree *RBTree) IsEmpty() bool {
	if tree != nil {
		return size(tree.root) == 0
	}
	return true
}

// Get 根据 key 值获取 value
func (tree *RBTree) Get(key int) int {
	if tree.root != nil {
		return get(tree.root, key)
	}
	return 0
}

// get 根据 key 值获取 value
func get(node *RBNode, key int) int {
	if node == nil {
		return 0
	}
	if key < node.key {
		return get(node.left, key)
	} else if key > node.key {
		return get(node.right, key)
	} else {
		return node.value
	}
}

// Min 求树上最小值
func (tree *RBTree) Min() int {
	return min(tree.root).key
}

// min 求此节点下的最小值
func min(node *RBNode) *RBNode {
	if node.left == nil {
		return node
	}
	return min(node.left)
}

// flipColors 转换一个节点的两个子节点的颜色。
// 在插入的时候：除了将子节点由红变黑之外，还需要将父节点由黑变红。
// 在删除最小的时候：用于从父结点借一个键值或者还给父节点一个键值
func flipColors(node *RBNode) {
	node.color = !node.color
	node.left.color, node.right.color = !node.left.color, !node.right.color
}

// Insert 向树内插入一个键值对
func (tree *RBTree) Insert(key int, value int) {
	// 首先找 key 值存不存在，存在更新其值，否则为它新建一个节点
	tree.root = put(tree.root, key, value)
	// 根节点颜色置为黑
	tree.root.color = BLACK
}

// put 节点的插入函数
func put(node *RBNode, key int, value int) *RBNode {
	if node == nil {
		return &RBNode{key, value, nil, nil, 1, RED}
	}
	// 比较 key 值和当前节点的 key 值，大了向右找位置，小了向左找位置
	// 相等更新当前节点的 value
	if key > node.key {
		node.right = put(node.right, key, value)
	} else if key < node.key {
		node.left = put(node.left, key, value)
	} else {
		node.value = value
	}

	// 接下来进行颜色的转换
	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if node.left != nil { // node.left.left 这一块可能空指针，判断一下
		if isRed(node.left) && isRed(node.left.left) {
			node = rotateRight(node)
		}
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

// moveRedLeft 主要实现以下两点：
// 当前节点的左右子节点都是2-节点，左右节点需要从父节点中借一个节点
// 如果该节点的右节点的左节点是红色节点，说明兄弟节点不是2-节点，可以从兄弟节点中借一个
func moveRedLeft(node *RBNode) *RBNode {
	// 从父结点中借一个
	flipColors(node)
	// 判断兄弟节点，如果是非红节点，也从兄弟节点中借一个
	if node.right != nil {
		if isRed(node.right.left) {
			node.right = rotateRight(node.right)
			node = rotateLeft(node)
			flipColors(node)
		}
	}
	return node
}

// DeleteMin 删除最小元素的函数入口
func (tree *RBTree) DeleteMin() error {
	if tree.root != nil {
		// 如果根节点的左右子节点是2-节点，我们可以将根设为红节点，这样才能进行后面的 moveRedLeft 操作，因为左子要从根节点借一个
		if !isRed(tree.root.left) && !isRed(tree.root.right) {
			tree.root.color = RED
		}
		tree.root = deleteMin(tree.root)
		// 借完以后，我们将根节点的颜色复原
		if !tree.IsEmpty() {
			tree.root.color = BLACK
		}
		return nil
	}
	return errors.New("树为空，不存在最小元素！！！")
}

// deleteMin 删除最小元素
func deleteMin(node *RBNode) *RBNode {
	// 删除
	if node.left == nil {
		return nil
	}
	// 判断 node 的左节点是不是2-节点（再向坐下递归的过程中，保证之前讨论的那三种情况成立）
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	// 递归
	node.left = deleteMin(node.left)
	// 解除临时组成的 4- 节点，更新以此节点的棵子树中的节点总数
	return balance(node)
}

// balance 解除临时组成的 4- 节点，更新以此节点的棵子树中的节点总数
func balance(node *RBNode) *RBNode {

	if isRed(node.right) {
		node = rotateLeft(node)
	}
	// TODO: 下面这个 if 可以去掉吧
	if isRed(node.right) && !isRed(node.left) {
		node = rotateLeft(node)
	}
	if node.left != nil { // node.left.left 这一块可能空指针，判断一下
		if isRed(node.left) && isRed(node.left.left) {
			node = rotateRight(node)
		}
	}
	if isRed(node.left) && isRed(node.right) {
		flipColors(node)
	}
	node.N = size(node.left) + size(node.right) + 1
	return node
}

// moveRedRight 主要实现以下两个功能：
// 当前节点的右子节点是 2- 节点，而兄弟节点不是 2- 节点，则从右子节点的兄弟节点那儿移动一个键值到右子节点中
// 前节点的右子节点和它的兄的节点都是 2- 节点，则将右子节点、父亲节点中的最大值和右子节点的兄弟节点合并成为一个 4- 节点
func moveRedRight(node *RBNode) *RBNode {
	// 从父节点中取一个
	flipColors(node)
	// 判断兄弟节点，如果不是 2- 节点，也从兄弟节点中借一个
	if node.left != nil {
		if isRed(node.left.left) {
			// 从兄弟那儿借来了一个
			node = rotateRight(node)
			// 从兄弟那儿取来了，那就再还给父节点一个
			flipColors(node)
		}
	}
	return node
}

// DeleteMax 删除最大元素的函数入口
func (tree *RBTree) DeleteMax() error {
	if tree.root != nil {
		if !isRed(tree.root.left) && !isRed(tree.root.right) {
			tree.root.color = RED
		}
		tree.root = deleteMax(tree.root)
		if !tree.IsEmpty() {
			tree.root.color = BLACK
		}
		return nil
	}
	return errors.New("树为空，不存在最大元素！！！")
}

// deleteMax 删除最大元素
func deleteMax(node *RBNode) *RBNode {
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	// 右侧节点是 2- 节点，则在递归的时候需要执行 moveRedRight 使其满足之前说的三种情况之一
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	node.right = deleteMax(node.right)
	return balance(node)
}

// Delete 删除函数，删除树上指定的 key
func (tree *RBTree) Delete(key int) error {
	if tree.root != nil {
		if !isRed(tree.root.left) && !isRed(tree.root.right) {
			tree.root.color = RED
		}
		tree.root = deleteKey(tree.root, key)
		if !tree.IsEmpty() {
			tree.root.color = BLACK
		}
		return nil
	}
	return errors.New("树为空，不存在最大元素！！！")
}

// deleteKey 删除树上指定的 key 值
func deleteKey(node *RBNode, key int) *RBNode {
	// 比较小，向左找
	if key < node.key {
		// 每当向左进行移动的时候，一定要将当前节点转换成一个非 2- 节点
		// 也就是每次向左走都要执行前面删除最小时候向左的转换 moveRedLeft
		// 删除
		if node.left == nil {
			return nil
		}
		// 判断 node 的左节点是不是2-节点（再向坐下递归的过程中，保证之前讨论的那三种情况成立）
		if !isRed(node.left) && !isRed(node.left.left) {
			node = moveRedLeft(node)
		}
		// 递归
		node.left = deleteKey(node.left, key)
	} else {
		// 同样，如果需要向右走或者删除当前节点（删除当前节点需要取后继）
		// 也需要执行前面删除最大的时候向右的变换，使其当前节点不是 2- 节点
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if node.right == nil && key != node.key{ // 没有找到目标键，可以退出了
			return nil
		}
		// 右侧节点是 2- 节点，则在递归的时候需要执行 moveRedRight 使其满足之前说的三种情况之一
		if !isRed(node.right) && !isRed(node.right.left) {
			node = moveRedRight(node)
		}

		// 上面对节点的变换完成了，下面就开始删除
		if key == node.key {
			// 取右侧最小（后继）填入当前节点
			node.value = get(node.right, min(node.right).key)
			node.key = min(node.right).key
			// 再删除右侧最小的（后继）
			node.right = deleteMin(node.right)
		} else {
			// 这是当前节点值还是小于 key 值，继续向右找
			node.right = deleteKey(node.right, key)
		}
	}
	return balance(node)
}
