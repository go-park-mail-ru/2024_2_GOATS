package models

type Data struct {
	QuestionID int64
	AnswerID   int64
	Answer     string
}

type Question struct {
	IsActive   bool
	QuestionID int64
	Question   string
	Answers    []string
}
