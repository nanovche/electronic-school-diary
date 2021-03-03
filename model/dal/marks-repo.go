package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type IMarkRepository interface{
	InsertMark(mark model.Mark) error
	GetMarkByID(markID int) (mark *model.Mark, err error)
/*	UpdateMark(mark model.Mark) error
	DeleteMarkByID(markID int) error*/
}

type MarkRepositoryImpl struct {
	db *gorm.DB
}

func NewMarkRepository(db *gorm.DB) MarkRepositoryImpl{
	return MarkRepositoryImpl{db: db}
}

func (mr MarkRepositoryImpl) InsertMark(mark model.Mark) error {

	res := mr.db.Create(&mark)
	if res.Error != nil {
		return fmt.Errorf("failed to inser mark: %s", res.Error)
	}
	return nil
}

func (mr MarkRepositoryImpl) GetMarkByID(markID int) (mark *model.Mark, err error) {

	res := mr.db.First(&mark, markID)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to retrieve mark")
	}
	return mark, nil
}

/*func (mr MarkRepositoryImpl) UpdateMark(mark model.Mark) error  {

	res := mr.db.Create(mark)
	if res.Error != nil {
		return fmt.Errorf("failed to inser mark")
	}
	return nil
}

func (mr MarkRepositoryImpl) DeleteMarkByID(markID int) error {

	res := mr.db.First(&mark, markID)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to retrieve mark")
	}
	return mark, nil
}*/



