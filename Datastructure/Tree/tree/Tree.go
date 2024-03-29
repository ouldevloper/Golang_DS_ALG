package tree

//	type TreeNode struct {
//		value     int
//		LeftNode  *Tree
//		RightNode *Tree
//	}
type Tree struct {
	value     int
	LeftNode  *Tree
	RightNode *Tree
}

func (tree *Tree) Init(value int) {
	tree = &Tree{

		value:     value,
		LeftNode:  nil,
		RightNode: nil,
	}
}

func (tree *Tree) Insert(value int) {
	node := &Tree{value, nil, nil}
	if tree == nil {
		tree = node
		return
	}
	if tree.LeftNode == nil {
		tree.LeftNode = node
		return
	}
	if tree.RightNode == nil {
		tree.RightNode = node
		return
	}
	if tree.LeftNode != nil {
		tree.LeftNode.Insert(value)
		return
	} else {
		tree.RightNode.Insert(value)
		return
	}
}

func (tree *Tree) Print(msg string) {
	if tree == nil {
		return
	}
	println(msg, " - Value : ", tree.value)
	tree.LeftNode.Print(msg + "-> Left")
	tree.RightNode.Print(msg + "-> Right")

}
