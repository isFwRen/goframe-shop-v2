// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure of table rotation_info for DAO operations like Where/Data.
type RotationInfo struct {
	g.Meta    `orm:"table:rotation_info, do:true"`
	Id        interface{} // 主键id
	PicUrl    interface{} //
	Link      interface{} //
	Sort      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
