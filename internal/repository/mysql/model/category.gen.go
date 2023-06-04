// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int64     `gorm:"column:parent_id;not null" json:"parent_id"`
	Sort      int64     `gorm:"column:sort;not null;comment:排序" json:"sort"`
	Title     string    `gorm:"column:title;not null;comment:名称" json:"title"`
	Icon      string    `gorm:"column:icon;not null;comment:图标" json:"icon"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	IsUsed    int64     `gorm:"column:is_used;default:-1;comment:是否启用 1:是 -1:否" json:"is_used"`
	Level     int64     `gorm:"column:level;comment:分类等级" json:"level"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
