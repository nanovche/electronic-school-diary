package service

import (
	"electronic-school-diary/model/dal"
	model "electronic-school-diary/model/entities"
	"fmt"
)

type ITeacherService interface{
	AssessStudent(studentName, teacherName, subjectTitle string, markValue int) error
	GetIRepository() dal.IRepository

}

type TeacherServiceImpl struct{
	repo dal.IRepository
}

func NewTeacherService(repo dal.IRepository) ITeacherService{
	return TeacherServiceImpl{repo: repo}
}

func (ts TeacherServiceImpl) GetIRepository() dal.IRepository {
	return ts.repo
}

func (ts TeacherServiceImpl) AssessStudent(studentName, teacherName, subjectTitle string, markValue int) error {

	studentID, err := ts.repo.GetStudentRepository().GetStudentIDByName(studentName)
	if err != nil {
		return fmt.Errorf("failed getting student id: %s", err)
	}
	subjectID, err := ts.repo.GetSubjectRepository().GetSubjectIDByTitle(subjectTitle)
	if err != nil {
		return fmt.Errorf("failed getting subject id: %s", err)
	}
	teacherID, err := ts.repo.GetTeacherRepository().GetTeacherIDByName(teacherName)
	if err != nil {
		return fmt.Errorf("failed getting teacher id: %s", err)
	}

	mark := model.Mark{
		Student_ID: studentID,
		Teacher_ID: teacherID,
		Subject_ID: subjectID,
		Value:  markValue,
	}

	err = ts.repo.GetMarkRepository().InsertMark(mark)
	if err != nil {
		return fmt.Errorf("Error adding mark: %s", err)
	}
	return nil

}
