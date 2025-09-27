// Package egorm gorm 封装
package egorm

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EDB struct {
	*gorm.DB
}

func (inst *EDB) WithContext(ctx context.Context) *gorm.DB {
	var reqid string
	var ok bool
	reqid, ok = ctx.Value("requestID").(string)
	if !ok {
		reqid = uuid.New().String()
	}
	ctx = context.WithValue(ctx, "requestID", reqid)

	return inst.DB.WithContext(ctx)
}
