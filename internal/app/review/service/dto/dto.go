package dto

type CheckReviewData struct {
	CSAT bool
	NPS  bool
	CSI  bool
}

type CreateReviewData struct {
	ID         int
	AnswerID   int
	AnswerText string
}

type ReviewData struct {
	ID      int
	Title   string
	Answers []Answer
	Type    string
}

type Answer struct {
	ID      int
	Content string
}

type Statistic struct {
	Rating   float64
	Type     string
	Comments []string
}
