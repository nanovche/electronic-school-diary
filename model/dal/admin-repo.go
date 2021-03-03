package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type IRepositoryAdmin interface {

	CreateTeacher(teacher *model.Teacher) error
 	GetTeacherPass(teacherEmail string) (password string , err error)
 	//GetTeacherID(teacherEmail string) (teacherID uint , err error)
	/*UpdateTeacher(student *model.Teacher) error
	DeleteTeacher(student *model.Teacher) error*/

	AddStudent(student *model.Student) error
	GetAllStudents() (student []*model.Student, err error)
	UpdateStudent(studentID int) (student *model.Student, err error)
	RemoveStudent(studentID int) (student *model.Student, err error)
}

type RepositoryAdminImpl struct {
	db *gorm.DB
}

func NewRepositoryAdmin(db *gorm.DB) RepositoryAdminImpl {
	return RepositoryAdminImpl{db: db}
}


func (cl RepositoryAdminImpl) CreateTeacher(teacher *model.Teacher) error {

	res := cl.db.Create(teacher)
	if res.Error != nil {
		return fmt.Errorf("failed to create teacher: %s", res.Error)
	}
	return nil
}


func (cl RepositoryAdminImpl) AddStudent(student *model.Student) error {

	res := cl.db.Create(student)
	if res.Error != nil {
		return fmt.Errorf("failed to create student: %s", res.Error)
	}
	return nil
}

func (cl RepositoryAdminImpl) CreateStudent(student *model.Student) error {

	res := cl.db.Create(student)
	if res.Error != nil {
		return fmt.Errorf("failed to inser mark")
	}
	return nil
}

func (cl RepositoryAdminImpl) GetStudentByID(studentID int) (student *model.Student, err error) {

	res := cl.db.First(&student, studentID)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to retrieve mark")
	}
	return

}
