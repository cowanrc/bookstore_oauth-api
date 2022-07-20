package access_token

import (
	"bookstore_oauth-api/src/utils/errors"
	"strings"
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
	atId = strings.TrimSpace((atId))
	if len(atId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(atId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil

}
