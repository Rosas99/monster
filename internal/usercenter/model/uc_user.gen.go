// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserM = "uc_user"

// UserM mapped from table <uc_user>
type UserM struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username  string    `gorm:"column:username;not null;uniqueIndex:username,priority:1" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	Nickname  string    `gorm:"column:nickname;not null" json:"nickname"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Phone     string    `gorm:"column:phone;not null" json:"phone"`
	Status    string    `gorm:"column:status;comment:用户状态：registered,active,disabled,blacklisted,locked,deleted" json:"status"` // 用户状态：registered,active,disabled,blacklisted,locked,deleted
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// TableName UserM's table name
func (*UserM) TableName() string {
	return TableNameUserM
}