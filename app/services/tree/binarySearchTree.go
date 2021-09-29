package tree

type BinaryTree struct {
	Data int
	Left *BinaryTree
	Right *BinaryTree
}

type BinarySearchTree interface {
	Insert(num int)
	Find(num int) bool
	Delete(num int)
	Show() []int
	MaxDepth() int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(num int) {
	if b.Data == 0 {
		b.Data = num
		return
	}
	tmpTree := b
	for tmpTree != nil {
		if num > tmpTree.Data {
			if tmpTree.Right == nil {
				tmpTree.Right = &BinaryTree{Data: num}
				return
			}
			tmpTree = tmpTree.Right
		} else {
			if tmpTree.Left == nil {
				tmpTree.Left = &BinaryTree{Data: num}
				return
			}
			tmpTree = tmpTree.Left
		}
	}
}

func (b *BinaryTree) Find(num int) bool {
	tmpTree := b
	for tmpTree != nil {
		if num > tmpTree.Data {
			tmpTree = tmpTree.Right
		} else if num < tmpTree.Data {
			tmpTree = tmpTree.Left
		} else {
			return true
		}
	}
	return false
}

func (b *BinaryTree) Delete(num int) {
	cur := b
	var curParent *BinaryTree = nil
	for cur != nil && cur.Data != num {
		curParent = cur
		if num > cur.Data {
			cur = cur.Right
 		} else {
			 cur = cur.Left
		}
	}
	if cur == nil {
		return
	}

	//要删除的节点有两个子节点时
	if cur.Left != nil && cur.Right != nil {
		min := cur.Right
		minParent := cur
		for min.Left != nil {
			minParent = min
			min = min.Left
		}
		cur.Data = min.Data
		cur = min
		curParent = minParent
	}

	//删除节点是叶子节点或者仅有一个子节点
	var child *BinaryTree
	if cur.Left != nil {
		child = cur.Left
	} else if cur.Right != nil {
		child = cur.Right
	} else {
		child = nil
	}

	if curParent == nil {
		b = child
	} else if curParent.Left == cur {
		curParent.Left = child
	} else {
		curParent.Right = child
	}
}

func (b *BinaryTree) Show() (res []int) {
	var handleFunc func(tree *BinaryTree)
	handleFunc = func(tree *BinaryTree) {
		if tree == nil {
			return
		}

		handleFunc(tree.Left)
		res = append(res, tree.Data)
		handleFunc(tree.Right)
	}
	handleFunc(b)
	return
}


//递归方法：深度优先
//func (b *BinaryTree) MaxDepth() int {
//	if b == nil {
//		return 0
//	}
//	if b.Left == nil && b.Right == nil {
//		return 1
//	}
//	left := b.Left.MaxDepth()
//	right := b.Right.MaxDepth()
//	return max(left, right) + 1
//}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//层序遍历
func (b *BinaryTree) MaxDepth() int {
	if b == nil {
		return 0
	}
	node := b
	var queue []*BinaryTree
	queue = append(queue, node)
	var front int = 0 //队头
	var rear int = len(queue) //队尾
	var floor int = 1
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		front++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if front == rear {
			front = 0
			rear = len(queue)
			if rear != 0 {
				floor++
			}
		}
	}
	return floor
}





