package service

import (
	"electronic-school-diary/model/dal"
)

type IMarkService interface{
	GetIRepository() dal.IRepository
}

type MarkServiceImpl struct{
	repo dal.IRepository
}

func NewMarkService(repo dal.IRepository) IMarkService{
	return MarkServiceImpl{repo: repo}
}

func (ts MarkServiceImpl) GetIRepository() dal.IRepository {
	return ts.repo
}


