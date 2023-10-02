// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen_model

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNameDataTypeDemo = "data_type_demo"

// DataTypeDemo mapped from table <data_type_demo>
type DataTypeDemo struct {
	ID                string         `gorm:"column:id;primaryKey;comment:ID" json:"id"`                                         // ID
	DataTypeBool      bool           `gorm:"column:data_type_bool;comment:数据类型 bool" json:"dataTypeBool"`                       // 数据类型 bool
	DataTypeInt2      int16          `gorm:"column:data_type_int2;comment:数据类型 int2" json:"dataTypeInt2"`                       // 数据类型 int2
	DataTypeInt8      int64          `gorm:"column:data_type_int8;comment:数据类型 int8" json:"dataTypeInt8"`                       // 数据类型 int8
	DataTypeVarchar   string         `gorm:"column:data_type_varchar;default:test;comment:数据类型 varchar" json:"dataTypeVarchar"` // 数据类型 varchar
	DataTypeText      string         `gorm:"column:data_type_text;comment:数据类型 text" json:"dataTypeText"`                       // 数据类型 text
	DataTypeJSON      datatypes.JSON `gorm:"column:data_type_json;comment:数据类型 json" json:"dataTypeJson"`                       // 数据类型 json
	CreatedAt         time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"createdAt"`                          // 创建时间
	UpdatedAt         time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updatedAt"`                          // 更新时间
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deletedAt"`                                   // 删除时间
	DataTypeTimeNull  sql.NullTime   `gorm:"column:data_type_time_null;comment:数据类型 time null" json:"dataTypeTimeNull"`         // 数据类型 time null
	DataTypeTime      time.Time      `gorm:"column:data_type_time;not null;comment:数据类型 time not null" json:"dataTypeTime"`     // 数据类型 time not null
	DataTypeJsonb     datatypes.JSON `gorm:"column:data_type_jsonb;comment:数据类型 jsonb" json:"dataTypeJsonb"`                    // 数据类型 jsonb
	DataTypeDate      time.Time      `gorm:"column:data_type_date" json:"dataTypeDate"`
	DataTypeFloat4    float32        `gorm:"column:data_type_float4" json:"dataTypeFloat4"`
	DataTypeFloat8    float64        `gorm:"column:data_type_float8" json:"dataTypeFloat8"`
	UlID              string         `gorm:"column:_id;comment:验证下划线" json:"_id"`              // 验证下划线
	CacheKey          string         `gorm:"column:cacheKey;comment:特殊保留字段名称" json:"cacheKey"` // 特殊保留字段名称
	DataTypeTimestamp time.Time      `gorm:"column:data_type_timestamp" json:"dataTypeTimestamp"`
	DataTypeBytea     []uint8        `gorm:"column:data_type_bytea" json:"dataTypeBytea"`
	DataTypeNumeric   float64        `gorm:"column:data_type_numeric" json:"dataTypeNumeric"`
	DataTypeInterval  string         `gorm:"column:data_type_interval" json:"dataTypeInterval"`
	BatchAPI          string         `gorm:"column:batch_api" json:"batchApi"`
}

// TableName DataTypeDemo's table name
func (*DataTypeDemo) TableName() string {
	return TableNameDataTypeDemo
}
