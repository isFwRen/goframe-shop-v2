// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-shop-v2/internal/model"
)

type (
	IRole interface {
		Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error)
		// Delete 删除
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
