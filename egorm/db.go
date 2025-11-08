// Package egorm gorm 封装
package egorm

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EDB struct {
	DB *gorm.DB
}

func NewEDB(db *gorm.DB) *EDB {
	return &EDB{
		DB: db,
	}
}

func (inst *EDB) WithContext(ctx context.Context) *EDB {
	var reqid string
	var ok bool
	reqid, ok = ctx.Value("requestID").(string)
	if !ok {
		reqid = uuid.New().String()
	}
	ctx = context.WithValue(ctx, "requestID", reqid)

	return inst.DB.WithContext(ctx)
}

func (e *EDB) StructToMap(model any, opts ...Option) (map[string]any, error) {
	options := &options{
		includeZero:    true,
		includePrimary: false,
		skipFields:     make([]string, 0),
		onlyFields:     make([]string, 0),
	}

	for _, opt := range opts {
		opt(options)
	}

	result := make(map[string]any)
	val := reflect.ValueOf(model)

	// 处理指针类型
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("model must be a struct")
	}

	typ := val.Type()
	numFields := val.NumField()

	for i := range numFields {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// 跳过非导出字段
		if !fieldValue.CanInterface() {
			continue
		}

		fieldName := getFieldName(field)
		if fieldName == "" {
			continue
		}

		// 检查是否跳过该字段
		if shouldSkipField(fieldName, options) {
			continue
		}

		// 检查主键字段
		if isPrimaryKey(field) && !options.includePrimary {
			continue
		}

		// 检查零值处理
		if !options.includeZero && fieldValue.IsZero() {
			continue
		}

		// 处理时间类型
		if t, ok := fieldValue.Interface().(time.Time); ok {
			result[fieldName] = t
			continue
		}

		// 处理其他类型
		result[fieldName] = fieldValue.Interface()
	}

	return result, nil
}

// UpdateWithMap 使用Map更新数据，避免零值问题
func (mc *EDB) UpdateWithMap(model any, opts ...Option) *gorm.DB {
	data, err := mc.StructToMap(model, opts...)
	if err != nil {
		return mc.Model(model).Where("1=0") // 返回一个不会执行的查询
	}

	return mc.Model(model).Updates(data)
}

// UpdateSelective 选择性更新（排除零值）
func (mc *EDB) UpdateSelective(model any, opts ...Option) *gorm.DB {
	opts = append(opts, WithExcludeZero())
	return mc.UpdateWithMap(model, opts...)
}

// 配置选项
type options struct {
	includeZero    bool
	includePrimary bool
	skipFields     []string
	onlyFields     []string
}

type Option func(*options)

// WithIncludeZero 包含零值字段
func WithIncludeZero() Option {
	return func(o *options) {
		o.includeZero = true
	}
}

// WithExcludeZero 排除零值字段
func WithExcludeZero() Option {
	return func(o *options) {
		o.includeZero = false
	}
}

// WithIncludePrimary 包含主键字段
func WithIncludePrimary() Option {
	return func(o *options) {
		o.includePrimary = true
	}
}

// WithSkipFields 跳过指定字段
func WithSkipFields(fields ...string) Option {
	return func(o *options) {
		o.skipFields = fields
	}
}

// WithOnlyFields 只包含指定字段
func WithOnlyFields(fields ...string) Option {
	return func(o *options) {
		o.onlyFields = fields
	}
}

// 工具函数
func getFieldName(field reflect.StructField) string {
	// 优先使用gorm标签
	gormTag := field.Tag.Get("gorm")
	if gormTag != "" {
		if gormTag == "-" {
			return ""
		}

		// 解析gorm标签中的column
		for _, part := range strings.Split(gormTag, ";") {
			if after, ok := strings.CutPrefix(part, "column:"); ok {
				return after
			}
		}
	}

	// 使用json标签
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" {
		if jsonTag == "-" {
			return ""
		}
		parts := strings.Split(jsonTag, ",")
		return parts[0]
	}

	// 默认使用字段名（转换为snake_case）
	return toSnakeCase(field.Name)
}

func isPrimaryKey(field reflect.StructField) bool {
	gormTag := field.Tag.Get("gorm")
	return strings.Contains(gormTag, "primaryKey")
}

func shouldSkipField(fieldName string, opts *options) bool {
	// 检查跳过字段
	if slices.Contains(opts.skipFields, fieldName) {
		return true
	}

	// 检查仅包含字段
	if len(opts.onlyFields) > 0 {
		found := slices.Contains(opts.onlyFields, fieldName)
		if !found {
			return true
		}
	}

	return false
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
