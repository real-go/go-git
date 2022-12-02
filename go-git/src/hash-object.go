package gogit

import "fmt"

func HashObject(body []byte, tp ObjectType, isWrite bool) (string, error) {
	switch tp {
	case BlobType:
		var blob = Blob{Body: body}
		out, err := blob.HashObject(isWrite)
		if err != nil {
			return out, err
		}
		return out, nil
	default:
		return "", fmt.Errorf("unknown object type")
	}
}
