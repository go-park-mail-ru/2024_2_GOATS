package dto

type DataDTO struct {
	QuestionID int64
	AnswerID   int64
	Answer     string
}

type QuestionDTO struct {
	IsActive   bool
	QuestionID int64
	Question   string
	Answers    []string
}
