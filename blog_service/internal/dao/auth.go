package dao

import "github.com/zzhaolei/go-programming-tour-book/blog_service/internal/model"

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{
		AppKey:    appKey,
		AppSecret: appSecret,
		State:     1,
	}
	return auth.Get(d.engine)
}
