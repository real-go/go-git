package gogit

import "testing"

func TestHashObject(t *testing.T) {
	type args struct {
		body    []byte
		tp      ObjectType
		isWrite bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"blob no write", args{[]byte("hello world"), BlobType, false}, "95d09f2b10159347eece71399a7e2e907ea3df4f", false},
		{"blob write", args{[]byte("hello world"), BlobType, true}, "95d09f2b10159347eece71399a7e2e907ea3df4f", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HashObject(tt.args.body, tt.args.tp, tt.args.isWrite)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HashObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}
