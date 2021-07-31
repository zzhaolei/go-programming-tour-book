package service

import "errors"

type AuthRequest struct {
	AppKey    string `json:"app_key" binding:"required"`
	AppSecret string `json:"app_secret" binding:"required"`
}

// CheckAuth 检测用户是否存在
func (s *Service) CheckAuth(param *AuthRequest) error {
	auth, err := s.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
