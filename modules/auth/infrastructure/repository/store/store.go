package store

import (
	"context"

	"gorm.io/gorm"
)

type Handle struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Handle { return &Handle{db: db} }

func (h *Handle) S(ctx context.Context) *gorm.DB {
	return h.db.WithContext(ctx)
}

func (h *Handle) WithTx(ctx context.Context, fn func(tx *Handle) error) error {
	return h.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(&Handle{db: tx})
	})
}

func (h *Handle) First(ctx context.Context, dest any, query string, args ...any) error {
	return h.db.WithContext(ctx).Where(query, args...).First(dest).Error
}

func (h *Handle) Find(ctx context.Context, dest any, query string, args ...any) error {
	q := h.db.WithContext(ctx)
	if query != "" {
		q = q.Where(query, args...)
	}
	return q.Find(dest).Error
}

func (h *Handle) FindOrder(ctx context.Context, dest any, order, query string, args ...any) error {
	return h.db.WithContext(ctx).Where(query, args...).Order(order).Find(dest).Error
}

func (h *Handle) Create(ctx context.Context, value any) error {
	return h.db.WithContext(ctx).Create(value).Error
}

func (h *Handle) Updates(ctx context.Context, model any, updates map[string]any, query string, args ...any) *gorm.DB {
	return h.db.WithContext(ctx).Model(model).Where(query, args...).Updates(updates)
}

func (h *Handle) Update(ctx context.Context, model any, column string, value any, query string, args ...any) error {
	return h.db.WithContext(ctx).Model(model).Where(query, args...).Update(column, value).Error
}

func (h *Handle) Count(ctx context.Context, model any, query string, args ...any) (int64, error) {
	var n int64
	err := h.db.WithContext(ctx).Model(model).Where(query, args...).Count(&n).Error
	return n, err
}

func (h *Handle) LimitOffsetFind(ctx context.Context, dest any, model any, order string, limit, offset int, query string, args ...any) (int64, error) {
	q := h.db.WithContext(ctx).Model(model).Where(query, args...)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return 0, err
	}
	err := q.Order(order).Limit(limit).Offset(offset).Find(dest).Error
	return total, err
}
