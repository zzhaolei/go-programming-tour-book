package model

import (
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// Count 统计相同标签数量
func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Create 创建标签
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新标签
func (t Tag) Update(db *gorm.DB) error {
	return db.Model(&t).Updates(&t).Error
}

// Delete 删除标签
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
