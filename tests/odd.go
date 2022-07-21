package tests

var (
	ODD_KEYWORD  = "ODD"
	EVEN_KEYWORD = "EVEN"
)

func IsOdd(value int64) string {
	if value%2 == 0 {
		return EVEN_KEYWORD
	}

	return ODD_KEYWORD
}
