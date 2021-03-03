package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type IRepositoryTeacher interface {
	GetTeacherIDByName(teacherName string) (teacherID uint, err error)
	GetTeacherSubjects(teacherID int) (subjects []model.Subject, err error)
}

type RepositoryTeacherImpl struct {
	db *gorm.DB
}

func NewRepositoryTeacher(db *gorm.DB) RepositoryTeacherImpl {
	return RepositoryTeacherImpl{db: db}
}

func (rt RepositoryTeacherImpl) GetTeacherIDByName(teacherName string) (teacherID uint, err error) {

	res := rt.db.Where("first_name = ?", teacherName).Select("teacher_id").Find(&model.Teacher{})
	if res.Error != nil {
		return 0, fmt.Errorf("teacher_id query failed: %s", res.Error)
	}

	if err := res.Row().Scan(&teacherID); err != nil {
		return 0, fmt.Errorf("failed scanning teacher_id: %s", res.Error)
	}

	return
}

func (rt RepositoryTeacherImpl) GetTeacherSubjects(teacherID int) (subjects []model.Subject, err error) {

	rows, err := rt.db.Where("teachers.teacher_id = ?", teacherID).Table("teacher_subject").
		Select("subjects.subject_id, subjects.title").
		Joins("JOIN teachers on teachers.teacher_id = teacher_subject.teacher_id").
		Joins("JOIN subjects on subjects.subject_id = teacher_subject.subject_id").Rows()
	if err != nil {
		return nil, fmt.Errorf("error retrieveing subjects the specified teacher teaches")
	}

	for rows.Next() {
		var subjectID uint
		var subjectTitle string
		err = rows.Scan(&subjectID, &subjectTitle); if err != nil {
			return nil, fmt.Errorf("error scanning subject for the specified teacher ")
		}
		subject := model.Subject {subjectID, subjectTitle}
		subjects = append(subjects, subject)
	}

	return
}





