package dal

import (
	model "electronic-school-diary/model/entities"
	"electronic-school-diary/service/utils"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IMarkRepository interface{
	InsertMark(mark model.Mark) error
	GetMarkByID(markID int) (mark *model.Mark, err error)
	GetMarksByStudentIDTeacherID(studentID , teacherID uint) (markData map[string][][]interface{}, err error)
 	UpdateMark(markID, markValue int , date time.Time) error
	DeleteMarkByID(markID int) error
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

func (mr MarkRepositoryImpl) GetMarksByStudentIDTeacherID(studentID, teacherID uint) (map[string][][]interface{}, error) {

	rows, err := mr.db.Where("marks.student_id = ? AND teacher_id = ?", studentID, teacherID).Table("marks").
		Select("marks.inserted_at, subjects.title, marks.value, mark_id").
		Joins("JOIN subjects on subjects.subject_id = marks.subject_id").Rows()
	if err != nil {
		return nil, fmt.Errorf("failed to query mark records")
	}

	markData := make(map[string][][]interface{})
	for rows.Next(){
		var markID int
		var subjectTitle string
		var markValue int
		var date time.Time
		err = rows.Scan(&date, &subjectTitle, &markValue, &markID); if err != nil {
			return nil, fmt.Errorf("failed to scan mark record")
		}

		marks := utils.GetMarksMappedDigitToWord()
		mark := marks[markValue]
		mark = mark + "." + strconv.Itoa(markID)
		markData[subjectTitle] = append(markData[subjectTitle], []interface{}{date.Format("2006-01-02"), mark})
	}
	return markData, nil
}

func (mr MarkRepositoryImpl) UpdateMark(markID, markValue int, date time.Time) error {

	res := mr.db.Table("marks").
		Where("mark_id = ?", markID).
		Updates(map[string]interface{}{"value": markValue, "inserted_at": date})
	if res.Error != nil {
		return fmt.Errorf("failed to update mark with id %d: %s", markID, res.Error)
	}
	return nil

}

func (mr MarkRepositoryImpl) DeleteMarkByID(markID int) error {

	res := mr.db.Where("mark_id = ?", markID).Delete(&model.Mark{})
	if res.Error != nil {
		return fmt.Errorf("failed to delete mark with id %d :, %s", markID, res.Error)
	}
	return nil

}



