// Package egorm gorm 封装
package egorm

import (
	"context"
	"reflect"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newdb() {
	dsn := "root:root@tcp(127.0.0.1:3306)/card?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func TestEDB_WithContext(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &EDB{
				DB: tt.fields.DB,
			}
			if got := inst.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EDB.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEDB_StructToMap(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		model any
		opts  []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EDB{
				DB: tt.fields.DB,
			}
			got, err := e.StructToMap(tt.args.model, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("EDB.StructToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EDB.StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEDB_UpdateWithMap(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		model any
		opts  []Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &EDB{
				DB: tt.fields.DB,
			}
			if got := mc.UpdateWithMap(tt.args.model, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EDB.UpdateWithMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEDB_UpdateSelective(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		model any
		opts  []Option
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &EDB{
				DB: tt.fields.DB,
			}
			if got := mc.UpdateSelective(tt.args.model, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EDB.UpdateSelective() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithIncludeZero(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithIncludeZero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithIncludeZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithExcludeZero(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithExcludeZero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithExcludeZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithIncludePrimary(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithIncludePrimary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithIncludePrimary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithSkipFields(t *testing.T) {
	type args struct {
		fields []string
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
			if got := WithSkipFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithSkipFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOnlyFields(t *testing.T) {
	type args struct {
		fields []string
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
			if got := WithOnlyFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOnlyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
