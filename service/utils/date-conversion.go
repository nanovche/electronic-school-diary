package utils

import "time"

const (
	layoutISOFormat = "2006-01-02"
)

func StringToDate(dateAsString string) (date time.Time, err error){
	date, err = time.Parse(layoutISOFormat, dateAsString); if err != nil {

	}
	return
}

func DateToString(dateAsTime time.Time) string {
	return dateAsTime.Format(layoutISOFormat)
}