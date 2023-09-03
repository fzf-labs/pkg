// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package user_model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

const TableNameUserDemo = "user_demo"

// UserDemo mapped from table <user_demo>
type UserDemo struct {
	ID        int64          `gorm:"column:id;primaryKey;comment:ID" json:"id"`                   // ID
	UID       string         `gorm:"column:uid;not null;comment:uid" json:"uid"`                  // uid
	Username  string         `gorm:"column:username;not null;comment:用户账号" json:"username"`       // 用户账号
	Password  string         `gorm:"column:password;not null;comment:密码" json:"password"`         // 密码
	Nickname  string         `gorm:"column:nickname;not null;comment:用户昵称" json:"nickname"`       // 用户昵称
	Remark    string         `gorm:"column:remark;comment:备注" json:"remark"`                      // 备注
	DeptID    int64          `gorm:"column:dept_id;comment:部门ID" json:"deptId"`                   // 部门ID
	PostIds   string         `gorm:"column:post_ids;comment:岗位编号数组" json:"postIds"`               // 岗位编号数组
	Email     string         `gorm:"column:email;comment:用户邮箱" json:"email"`                      // 用户邮箱
	Mobile    string         `gorm:"column:mobile;comment:手机号码" json:"mobile"`                    // 手机号码
	Sex       int16          `gorm:"column:sex;comment:用户性别" json:"sex"`                          // 用户性别
	Avatar    string         `gorm:"column:avatar;comment:头像地址" json:"avatar"`                    // 头像地址
	Status    int16          `gorm:"column:status;not null;comment:帐号状态（0正常 -1停用）" json:"status"` // 帐号状态（0正常 -1停用）
	LoginIP   string         `gorm:"column:login_ip;comment:最后登录IP" json:"loginIp"`               // 最后登录IP
	LoginDate sql.NullTime   `gorm:"column:login_date;comment:最后登录时间" json:"loginDate"`           // 最后登录时间
	TenantID  int64          `gorm:"column:tenant_id;not null;comment:租户编号" json:"tenantId"`      // 租户编号
	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`    // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`    // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`             // 删除时间
}

// TableName UserDemo's table name
func (*UserDemo) TableName() string {
	return TableNameUserDemo
}
