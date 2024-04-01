// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:"用户名"`
	Password  string      `json:"password"  description:"密码"`
	RoleIds   string      `json:"roleIds"   description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeleteAt  *gtime.Time `json:"deleteAt"  description:""`
	UserSait  string      `json:"userSait"  description:"用户盐"`
	IsAdmin   int         `json:"isAdmin"   description:"是否超级管理员"`
}
