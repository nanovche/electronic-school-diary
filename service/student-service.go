package service

import (
	"electronic-school-diary/model/dal"
	"fmt"
)

type IStudentService interface{
	GetAllStudentNames() (studentNames []string, err error)
	GetStudentIDByName(studentName string) (studentID uint, err error)
	GetIRepository() dal.IRepository
}

type StudentServiceImpl struct{
	repo dal.IRepository
}

func NewStudentService(repo dal.IRepository) IStudentService{
	return StudentServiceImpl{repo: repo}
}

func (ts StudentServiceImpl) GetIRepository() dal.IRepository {
	return ts.repo
}

func(ts StudentServiceImpl) GetAllStudentNames() (studentNames []string, err error) {

	students, err := ts.repo.GetStudentRepository().GetAllStudents()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain all students")
	}

	for _, student := range students {
		studentNames = append(studentNames, student.FullName)
	}

	return
}

func(ts StudentServiceImpl) GetStudentIDByName(studentName string) (studentID uint, err error) {

	studentID, err = ts.repo.GetStudentRepository().GetStudentIDByName(studentName)
	if err != nil {
		return 0, fmt.Errorf("failed to obtain id of student %s: %s", studentName, err)
	}
	return
}

