package model

import "github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
