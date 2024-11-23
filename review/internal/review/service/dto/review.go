package dto

type DataDTO struct {
	QuestionID int64
	AnswerID   int64
	Answer     string
	Rating     int64
}

type QuestionDTO struct {
	IsActive   bool
	QuestionID int64
	Question   string
	Answers    []AnswerDTO
}

type AnswerDTO struct {
	ID      int
	Content string
}
