package constants

type SortingType string

const (
	SortingType1 SortingType = "sortByCreateDate"              // 가입순
	SortingType2 SortingType = "sortByCreateDateReverse"       // 가입역순
	SortingType3 SortingType = "sortByGradeClassNumber"        // 학년반번호순
	SortingType4 SortingType = "sortByGradeClassNumberReverse" // 학년반번호역순
)
