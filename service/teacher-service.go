package service

import (
	"electronic-school-diary/model/dal"
	model "electronic-school-diary/model/entities"
	utils "electronic-school-diary/service/utils"
	"fmt"
	"strconv"
	"time"
)

type ITeacherService interface{
	AssessStudent(studentName, teacherName, subjectTitle, markValue string ) error
	GetAllMarksOfStudentByOneTeacher(studentID, teacherID uint) (marks map[string][][]interface{}, err error)
	GetTeacherIDByName(teacherName string) (teacherID uint, err error)
	UpdateMark(markID, date, markValue string ) error
	DeleteMark(markID string ) error

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

func (ts TeacherServiceImpl) AssessStudent(studentName, teacherName, subjectTitle, markValue string) error {

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
		Value:  utils.Mark(markValue).Int(),
	}

	err = ts.repo.GetMarkRepository().InsertMark(mark)
	if err != nil {
		return fmt.Errorf("Error adding mark: %s", err)
	}
	return nil

}

func (ts TeacherServiceImpl) GetAllMarksOfStudentByOneTeacher(studentID, teacherID uint) (marks map[string][][]interface{}, err error) {

	marks, err = ts.GetIRepository().GetMarkRepository().GetMarksByStudentIDTeacherID(studentID, teacherID)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain marks of student with id %d by teacher with od %d: %s", studentID, teacherID, err)
	}
	return
}

func (ts TeacherServiceImpl) UpdateMark(markID, date, markValue string ) error {

	dateInDateFormat, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("failed to parse date: %s", err)
	}

	markIDAsInt, err := strconv.Atoi(markID)
	if err != nil {
		return fmt.Errorf("failed to convert string to int: %s", err)
	}

	markValueAsInt := utils.Mark(markValue).Int()
	err = ts.repo.GetMarkRepository().UpdateMark(markIDAsInt, markValueAsInt, dateInDateFormat)
	if err != nil {
		return fmt.Errorf("Error adding mark: %s", err)
	}
	return nil
}

func (ts TeacherServiceImpl) DeleteMark(markID string ) error {

	markIDAsInt, err := strconv.Atoi(markID)
	if err != nil {
		return fmt.Errorf("failed to convert string to int: %s", err)
	}

	err = ts.repo.GetMarkRepository().DeleteMarkByID(markIDAsInt)
	if err != nil {
		return fmt.Errorf("error deleting mark with id %d: %s", markIDAsInt, err)
	}
	return nil
}

func (ts TeacherServiceImpl) GetTeacherIDByName(teacherName string) (teacherID uint, err error) {

	teacherID, err = ts.repo.GetTeacherRepository().GetTeacherIDByName(teacherName)
	if err != nil {
		return 0, fmt.Errorf("error obtaining id of teacher  %s: %s", teacherName, err)
	}
	return
}




