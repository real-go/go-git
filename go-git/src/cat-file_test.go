package gogit

import (
	"reflect"
	"testing"
)

func TestReadBlobObj(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    *Blob
		wantErr bool
	}{
		{"read blob", args{"95d09f2b10159347eece71399a7e2e907ea3df4f"}, &Blob{Hash: "95d09f2b10159347eece71399a7e2e907ea3df4f", Body: []byte("hello world")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadBlobObj(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBlobObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBlobObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBlob(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"parse 1", args{[]byte("blob 16\\x00what is up, doc?")}, []byte("what is up, doc?")},
		{"parse 2", args{[]byte("blob 17\\x00hello world!!!!!!")}, []byte("hello world!!!!!!")},
		{"parse 3", args{[]byte("blob 11\\x00hello world")}, []byte("hello world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBlob(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBlob() = %v, want %v", got, tt.want)
			}
		})
	}
}
