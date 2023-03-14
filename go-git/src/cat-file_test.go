package gogit

import (
	"reflect"
	"testing"
)

func TestCatFile(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"read blob", args{"95d09f2b10159347eece71399a7e2e907ea3df4f"}, []byte("hello world"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CatFile(tt.args.hash)
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

func Test_parse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"ParseData blob1", args{[]byte("blob 16\x00what is up, doc?")}, []byte("what is up, doc?")},
		{"ParseData blob2", args{[]byte("blob 17\x00hello world!!!!!!")}, []byte("hello world!!!!!!")},
		{"ParseData blob3", args{[]byte("blob 11\x00hello world")}, []byte("hello world")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
