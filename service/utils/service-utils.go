package utils

//type Mark int

type Mark string

const (
	Слаб Mark = "Слаб"
	Среден = "Среден"
	Добър = "Добър"
	МногоДобър = "Много Добър"
	Отличен = "Отличен"
	)

var marks []string

func GetMarksAsSliceOfStrings() []string {
	marks = append(marks, string(Слаб))
	marks = append(marks, Среден)
	marks = append(marks, Добър)
	marks = append(marks, МногоДобър)
	marks = append(marks, Отличен)
	return marks
}

func(m Mark) Int() int {

	switch m {
	case Слаб:
		return 2
	case Среден:
		return 3
	case Добър:
		return 4
	case МногоДобър:
		return 5
	default:
		return 6
	}
}

/*const(
	Слаб Mark = iota + 2
	Среден
	Добър
	МногоДобър
	Отличен
)*/

/*func Int(val string) int {

	switch va {
	case Слаб:
		return 2
	case Среден:
		return 3
	case Добър:
		return 4
	case МногоДобър:
		return 5
	default:
		return 6
	}
}

func (m Mark) String() string {

	switch m {
	case Слаб:
		return "Слаб"
	case Среден:
		return "Среден"
	case Добър:
		return "Добър"
	case МногоДобър:
		return "Много Добър"
	default:
		return "Отличен"
	}
}*/

var marksMappedWordToDigit = make(map[string]int, 5)
var marksMappedDigitToWord = make(map[int]string, 5)

func GetMarksMappedWordToDigit() map[string]int {

	marksMappedWordToDigit["Слаб"] = 2
	marksMappedWordToDigit["Среден"] = 3
	marksMappedWordToDigit["Добър"] = 4
	marksMappedWordToDigit["Много Добър"] = 5
	marksMappedWordToDigit["Отличен"] = 6

	return  marksMappedWordToDigit
}
func GetMarksMappedDigitToWord() map[int]string {

	marksMappedDigitToWord[2] = "Слаб"
	marksMappedDigitToWord[3] = "Среден"
	marksMappedDigitToWord[4] = "Добър"
	marksMappedDigitToWord[5] = "Много Добър"
	marksMappedDigitToWord[6] = "Отличен"

	return  marksMappedDigitToWord
}

/*func GetMarksAsSliceOfStrings() []string {

	marksMappedDigitToWord[2] = "Слаб"
	marksMappedDigitToWord[3] = "Среден"
	marksMappedDigitToWord[4] = "Добър"
	marksMappedDigitToWord[5] = "Много Добър"
	marksMappedDigitToWord[6] = "Отличен"

	return  marksMappedDigitToWord
}*/
