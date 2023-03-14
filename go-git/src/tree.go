package gogit

type TreeNode struct {
	Hash          string
	ObjName       string // path or file name
	FNode         string
	ChildrenNodes []string
}

func (t *TreeNode) add(objName, hash string) {
	t.ChildrenNodes = append(t.ChildrenNodes, objName)
}

func (t *TreeNode) Type() ObjectType {
	return TreeType
}
