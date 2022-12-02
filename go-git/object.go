package main

type ObjectType int

const (
	BlobType ObjectType = iota
	TreeType
	CommitType
)

func (o ObjectType)String() string {
	switch o {
	case BlobType:
		return "blob"
	case TreeType:
		return "tree"
	default:
		return "unknown"
	}
}

type Object interface {
	Type() ObjectType
	HHash() string
	String() string
}
