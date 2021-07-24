package service

type ListArticleRequest struct {
	Name  string `form:"name" binding:"max=100"`
	Title string `form:"title" binding:"required,min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binding:"required,min=10"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binding:"required,min=10"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=255"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
