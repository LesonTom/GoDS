package Tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BTree struct {
	Root *Node
}

func NewBTree() *BTree {
	return &BTree{}
}

func insert(node *Node, val int) *Node {
	if node == nil {
		return &Node{Val: val}
	}

	if val < node.Val {
		node.Left = insert(node.Left, val)
	}
	if val > node.Val {
		node.Right = insert(node.Right, val)
	}
	return node
}

func (t *BTree) Insert(val int) {
	t.Root = insert(t.Root, val)
}

func search(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if node.Val == val {
		return node
	}

	if node.Val > val {
		return search(node.Left, val)
	}
	return search(node.Right, val)
}

func (t *BTree) Search(val int) *Node {
	node := search(t.Root, val)
	return node
}

func delete(node *Node, val int) {
	if node == nil {
		return
	}

}

func (t *BTree) Delete(val int) {

}

func (t *BTree) Height() int {
	return 0
}

func (t *BTree) InorderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) PreorderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) PostorderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) LevelOrderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) PostLevelOrderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) InLevelOrderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) LevelInOrderTraversal(root *Node) []int {
	return nil
}

func (t *BTree) PostInLevelOrderTraversal(root *Node) []int {
	return nil
}
