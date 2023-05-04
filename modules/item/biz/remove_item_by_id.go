package biz

import (
	"TodoApi/modules/item/model"
	"context"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{storage: storage}
}

func (biz *deleteItemBiz) DeleteItem(ctx context.Context, id int) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemIsDeleted
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
