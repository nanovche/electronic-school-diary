package dal

import (
	"electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type IRepositoryAuth interface {

	StoreSession(session model.Session) error
	RetrieveEmail(token string) (email string , err error)
	DeleteSession(token string) error
	GetTeacherPass(teacherEmail string) (password string , err error)

}

type RepositoryAuthImpl struct {
	db *gorm.DB
}

func NewRepositoryAuth(db *gorm.DB) IRepositoryAuth {
	return RepositoryAuthImpl{db: db}
}

func (cl RepositoryAuthImpl) GetTeacherPass(teacherEmail string) (password string , err error) {

	res := cl.db.Where("email = ?", teacherEmail).Select("password").Find(&model.Teacher{})
	if res.Error != nil {
		return "", fmt.Errorf("failed querying database: %s", res.Error)
	}
	if err := res.Row().Scan(&password); err != nil {
		return "", fmt.Errorf("failed scanning data: %s", res.Error)
	}

	return password, nil
}

func(cl RepositoryAuthImpl) StoreSession(session model.Session) error {

	if res := cl.db.Create(&session); res.Error != nil {
		return fmt.Errorf("failed storing cookie value and email to database: %s", res.Error)
	}
	return nil
}

func(cl RepositoryAuthImpl) RetrieveEmail(token string) (email string , err error) {

	res := cl.db.Where("session_token = ?", token).Select("Email").Find(&model.Session{})
	if res.Error != nil {
		return "", fmt.Errorf("failed querying session token(%s): %s", token, res.Error)
	}
	res.Scan(&email)
	return
}

func(cl RepositoryAuthImpl) DeleteSession(token string) error {

	res := cl.db.Where("session_token = ?", token).Delete(&model.Session{})
	if res.Error != nil {
		return fmt.Errorf("failed deleting record with token (%s): %s", token, res.Error)
	}
	return nil
}