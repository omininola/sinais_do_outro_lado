package lib

const (
	NumbersRangeStart  = 48
	NumbersRangeFinish = 57
)

const (
	LettersRangeStart  = 65
	LettersRangeFinish = 90
)

func GetSymbolsArr(theme string) []string {
	symbols := []string{"-", "$", "&"}
	if theme == "Sinais" {
		return symbols
	}

	return symbols
}
