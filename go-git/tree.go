package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type TreeNode struct {
	Hash          string
	FatherNode    string
	ChildrenNodes map[string]interface{}
}

func (t *TreeNode) HHash() string {
	return t.Hash
}

func (t *TreeNode) String() string {
	var buf bytes.Buffer
	for k, v := range t.ChildrenNodes {
		switch t := v.(type) {
		case Blob:
			buf.WriteString(fmt.Sprintf("%s\t%s\t%s", t.HHash(), t.Type().String(), k))
		case TreeNode:
			buf.WriteString(fmt.Sprintf("%s\t%s\t%s", t.HHash(), t.Type().String(), k))
		}
	}
	return buf.String()
}

func (t *TreeNode) Type() ObjectType {
	return TreeType
}

func (t *TreeNode) Add(file, hash string, oType ObjectType) error {
	if t.ChildrenNodes == nil {
		t.ChildrenNodes = make(map[string]interface{})
	}
	switch oType {
	case BlobType:
		t.ChildrenNodes[file] = &Blob{
			Hash: hash,
		}
	case TreeType:
		t.ChildrenNodes[file] = &TreeNode{
			Hash:       hash,
			FatherNode: t.Hash,
		}
	}
	return nil
}

func ReadTree(path, hash string) (*TreeNode, error) {
	f, err := OpenFile(path, hash, os.O_RDWR|os.O_CREATE)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}(f)
	g := gob.NewDecoder(f)
	t := new(TreeNode)
	if err := g.Decode(&t); err != nil {
		return nil, err
	}
	return t, nil
}

func Write(treeHash, fileName, hashValue string, oType ObjectType) error {
	var t *TreeNode
	var file *os.File
	var err error
	if treeHash != "" {
		t, err = ReadTree(ObjectsPath, treeHash)
	} else {
		t, err = ReadTree(ConfigPath, TempTreeFile)
	}
	if err != nil {
		return err
	}
	if err := t.Add(fileName, hashValue, oType); err != nil {
		return err
	}
	if treeHash != "" {
		file, err = OpenFile(ObjectsPath, treeHash, os.O_TRUNC|os.O_RDWR|os.O_CREATE)
	} else {
		file, err = OpenFile(ConfigPath, TempTreeFile, os.O_TRUNC|os.O_RDWR|os.O_CREATE)
	}

	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			panic(err)
		}

	}(file)
	if err := gob.NewEncoder(file).Encode(t); err != nil {
		return err
	}
	return nil
}

func LsTree(treeHash string) (string, error) {
	var tree = &TreeNode{}
	var err error
	if treeHash != "" {
		tree, err = ReadTree(ObjectsPath, treeHash)
	} else {
		tree, err = ReadTree(ConfigPath, TempTreeFile)
	}
	if err != nil {
		return "", err
	}
	return tree.String(), nil
}

func init() {
	gob.Register(TreeNode{})
	gob.Register(Blob{})
	gob.Register(map[string]interface{}{})
}
