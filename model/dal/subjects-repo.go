package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type IRepositorySubject interface {
	GetSubjectIDByTitle(subjectTitle string) (subjectID uint, err error)
}

type SubjectRepositoryImpl struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) IRepositorySubject{
	return SubjectRepositoryImpl{db: db}
}

func (sr SubjectRepositoryImpl) GetSubjectIDByTitle(subjectTitle string) (subjectID uint, err error) {

	res := sr.db.Where("title = ?", subjectTitle).Select("subject_id").Find(&model.Subject{})
	if res.Error != nil {
		return 0, fmt.Errorf("failed to subject_id query: %s", res.Error)
	}

	if err := res.Row().Scan(&subjectID); err != nil {
		return 0, fmt.Errorf("failed scanning subject_id: %s", res.Error)
	}

	return
}
