package service

import (
	"context"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	s := Service{ctx: ctx}
	s.dao = dao.New(global.DBEngine)
	return s
}
