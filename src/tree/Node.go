package tree

type Node struct {
	Label       rune
	IsWord      bool
	Children    []*Node
	Prefix      []rune
	Explanation string
}

func CreateNodeTree(words map[string]string) *Node {
	root := Node{
		IsWord: false,
	}
	for k, v := range words {
		addWord([]rune(k), v, &root)
	}
	return &root
	//traversal(&root)
}

func addWord(left []rune, explain string, root *Node) bool {
	if len(left) > 0 {
		children := root.Children
		for _, child := range children {
			if child.Label == left[0] {
				return addWord(left[1:], explain, child)
			}
		}
		node := Node{
			Label: left[0],
		}
		node.Prefix = append(root.Prefix, root.Label)
		root.Children = append(root.Children, &node)
		return addWord(left[1:], explain, &node)
	}
	root.IsWord = true
	root.Explanation = explain
	return true
}

func (n *Node) Search(str string) (string, error) {
	return search([]rune(str), n)
}

func search(left []rune, root *Node) (string, error) {
	if len(left) > 0 {
		children := root.Children
		for _, child := range children {
			if child.Label == left[0] {
				return search(left[1:], child)
			}
		}
		return "", &MyError{Msg: "未找到"}
	}
	if root.IsWord {
		return root.Explanation, nil
	}
	return "", &MyError{Msg: "未找到"}
}

func (n *Node) Traversal() {
	traversal(n)
}

func traversal(node *Node) {
	println(node)
	println("Label:", string(node.Label))
	println("Prefix:", string(node.Prefix))
	println("Explanation:", string(node.Explanation))
	println("IsWord:", node.IsWord)
	println("--------------分割线-----------------")
	for _, child := range node.Children {
		traversal(child)
	}
}
