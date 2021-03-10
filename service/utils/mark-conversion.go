package utils

func GetIntToStringMarkMapper() map[int]string{
	return map[int]string{
		2 : "Слаб",
		3 : "Среден",
		4 : "Добър",
		5 : "Много Добър",
		6 : "Отличен",
	}
}

func StringToIntMarkMapper() map[string]int{
	return map[string]int{
		"Слаб" : 2,
		"Среден" : 3,
		"Добър" : 4,
		"Много Добър" : 5,
		"Отличен" : 6,
	}
}

func GetMarksAsWords() []string{
	return []string{"Слаб", "Среден", "Добър", "Много Добър", "Отличен" }
}

