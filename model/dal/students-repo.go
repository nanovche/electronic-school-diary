package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type IStudentRepository interface {
	GetAllStudents() (students []model.Student, err error)
	GetStudentIDByName(studentName string) (studentID uint, err error)
}

type StudentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepositoryImpl(db *gorm.DB) IStudentRepository{
	return StudentRepositoryImpl{db: db}
}

func(sr StudentRepositoryImpl) GetAllStudents() (students []model.Student, err error) {

	res := sr.db.Find(&model.Student{})
	if res.Error != nil {
		return nil, fmt.Errorf("getting all students query failed: %s", res.Error)
	}

	rows, err := res.Rows()
	if err != nil {
		return nil, fmt.Errorf("getting all students query failed")
	}
	for rows.Next() {
		var student_id uint
		var firstName, lastName string
		var number uint8

		if err := rows.Scan(&student_id, &firstName, &lastName, &number); err != nil {
			return nil, fmt.Errorf("error scanning student record: %s ", err)
		}

		students = append(students, model.Student{
			Student_ID: student_id,
			FirstName:  firstName,
			LastName:   lastName,
			Number:     number,
		})
	}
	return
}

func (sr StudentRepositoryImpl) GetStudentIDByName(studentName string) (studentID uint, err error) {

	firstAndLastName := strings.Split(studentName, " ")
	firstName := firstAndLastName[0]
	lastName := firstAndLastName[1]
	res := sr.db.Where("first_name = ? AND last_name = ? ", firstName, lastName).Select("student_id").Find(&model.Student{})
	if res.Error != nil {
		return 0, fmt.Errorf("failed to query for student_id: %s", res.Error)
	}

	if err := res.Row().Scan(&studentID); err != nil {
		return 0, fmt.Errorf("failed scanning student_id: %s", res.Error)
	}

	return
}

/*func(cl RepositoryTeacherImpl) GetAllStudentNames()(studentNames []string, err error){

	res := cl.db.Select("first_name, last_name").Find(&model.Student{})
	if res.Error != nil {
		return nil, fmt.Errorf("failed to query for student names: %s", res.Error)
	}

	rows, err := res.Rows()
	if err != nil {
		return nil, fmt.Errorf("")
	}
	for rows.Next() {
		var firstName string
		var lastName string
		if err := rows.Scan(&firstName, &lastName); err != nil {
			return nil, fmt.Errorf("error scanning student record: %s ", err)
		}

		fullName := firstName + " " + lastName
		studentNames = append(studentNames, fullName)
	}
	return
}
*/