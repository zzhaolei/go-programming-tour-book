package model

type ArticleTag struct {
	*Model
	TagID        uint32 `json:"tag_id"`
	ArticleTagID string `json:"article_tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
