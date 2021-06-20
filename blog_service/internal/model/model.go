package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(d *setting.DatabaseSetting) (*gorm.DB, error) {
	// 格式化数据库链接
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%scharset=%s&parseTime=%t&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)

	// 定义日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	// 链接数据库，并使用自定义配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   d.TablePrefix,
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		return nil, err
	}

	refDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	refDB.SetMaxIdleConns(d.MaxIdleConns)
	refDB.SetMaxOpenConns(d.MaxOpenConns)

	return db, nil
}
