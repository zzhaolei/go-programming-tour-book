package model

import "github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         string `json:"state"`
}

func (a Article) TableName() string {
	return "bolg_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
