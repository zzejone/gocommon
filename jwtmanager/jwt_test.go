package jwtmanager

import (
	"reflect"
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func TestNewJWTManager(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *JWTManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJWTManager(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJWTManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSecretKey(t *testing.T) {
	type args struct {
		ipt string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSecretKey(tt.args.ipt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSecretKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetExpiresAt(t *testing.T) {
	type args struct {
		ipt time.Duration
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetExpiresAt(tt.args.ipt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExpiresAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetIssuer(t *testing.T) {
	type args struct {
		ipt string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetIssuer(tt.args.ipt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIssuer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSignMethod(t *testing.T) {
	type args struct {
		ipt jwt.SigningMethod
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetSignMethod(tt.args.ipt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSignMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_GenerateToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		data map[string]any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "GenerateToken",
			fields: fields{
				options: &JWTOptions{
					SecretKey:     "abc",
					ExpiresAt:     11 * time.Minute,
					Issuer:        "test",
					SigningMethod: nil,
				},
			},
			args: args{
				data: map[string]any{
					"userid": 1,
					"name":   "unknow",
				},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			_, err := jm.GenerateToken(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestJWTManager_ParseToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "parse",
			fields: fields{
				options: &JWTOptions{
					SecretKey:     "abb",
					ExpiresAt:     1 * time.Hour,
					Issuer:        "test",
					SigningMethod: jwt.SigningMethodHS256,
				},
			},
			args: args{
				tokenString: "",
			},
			want: map[string]any{
				"name": "name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			token, err := jm.GenerateToken(map[string]any{"name": "name"})
			if err != nil {
				t.Errorf("JWTManager.GenerateToken() error = %v", err)
				return
			}
			got, err := jm.ParseToken(token)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTManager.ParseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_ParseTokenWithClaims(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CustomClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			got, err := jm.ParseTokenWithClaims(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.ParseTokenWithClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTManager.ParseTokenWithClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_GetDataFromToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
		key         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			got, err := jm.GetDataFromToken(tt.args.tokenString, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.GetDataFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTManager.GetDataFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_GetStringFromToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
		key         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			got, err := jm.GetStringFromToken(tt.args.tokenString, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.GetStringFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWTManager.GetStringFromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_GetInt64FromToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
		key         string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			got, err := jm.GetInt64FromToken(tt.args.tokenString, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.GetInt64FromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWTManager.GetInt64FromToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJWTManager_RefreshToken(t *testing.T) {
	type fields struct {
		options *JWTOptions
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := &JWTManager{
				options: tt.fields.options,
			}
			got, err := jm.RefreshToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTManager.RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JWTManager.RefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
