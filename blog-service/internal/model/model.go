package model

import (
	"fmt"

	setting "github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreateOn   string `json:"create_on"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  string `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	// 使用 gorm v2 版本
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	_db, err := db.DB()
	if err != nil {
		return nil, err
	}
	_db.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	_db.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
