package dal

import (
	model "electronic-school-diary/model/entities"
	"fmt"
	"gorm.io/gorm"
)

type ITermRepository interface{
	GetStartAndEndDate()([]string , error)
}

type TermRepositoryImpl struct {
	db *gorm.DB
}

func NewTermRepository(db *gorm.DB) ITermRepository{
	return TermRepositoryImpl{db: db}
}

func (tr TermRepositoryImpl) GetStartAndEndDate()([]string , error) {

	res := tr.db.Where("CURDATE() BETWEEN start_date AND end_date").
		Select("start_date, end_date").
		Find(&model.Term{})
	if res.Error != nil {
		return nil, fmt.Errorf("failed to query term dates: %s", res.Error)
	}

	var startDate, endDate string
	if err := res.Row().Scan(&startDate, &endDate); err != nil {
		return nil, fmt.Errorf("failed to scan term dates: %s", res.Error)
	}

	return []string{startDate, endDate}, nil
}
