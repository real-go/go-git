package main

type Blob struct {
	Hash string
	Body string
}

func (b *Blob) HHash() string {
	return b.Hash
}

func (b *Blob) String() string {
	return b.Body
}

func (b *Blob) Type() ObjectType {
	return BlobType
}
