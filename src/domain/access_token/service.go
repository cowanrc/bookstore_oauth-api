package access_token

import (
	"bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(atId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(atId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil

}
