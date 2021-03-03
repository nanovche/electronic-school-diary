package utils

/*type Mark struct{
	Numeric int
	Title string
}

var Marks = []Mark{
	{2, "Слаб"},
	{3, "Среден"},
	{4, "Добър"},
	{5, "Много Добър"},
	{6, "Отличен"},
}
*/

var marks = make(map[string]int, 5)

func GetMarks() map[string]int {

	marks["Слаб"] = 2
	marks["Среден"] = 3
	marks["Добър"] = 4
	marks["Много Добър"] = 5
	marks["Отличен"] = 6

	return  marks
}

