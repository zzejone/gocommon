package pwd

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "生成密码",
			args: args{
				password: "abc123",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		password string
		encoded  string
	}
	mima := "abc12333"
	encodedPwd, err := HashPassword(mima)
	if err != nil {
		t.Errorf("对比密码准备失败：%v", err)
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "对比密码",
			args: args{
				password: mima,
				encoded:  encodedPwd,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "对比密码时密码不对的情况",
			args: args{
				password: mima + "a",
				encoded:  encodedPwd,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyPassword(tt.args.password, tt.args.encoded)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHashPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashPassword("123456")
	}
}

func BenchmarkVerifyPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerifyPassword("pwd", "encodecpassword")
	}
}
