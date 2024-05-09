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

func (t *BTree) Insert(val int) {
}

func (t *BTree) Search(val int) *Node {
	return nil
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
