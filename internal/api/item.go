package api

// Item models Search Response items (questions)
type Item struct {
	AnswerCount int      `json:"answer_count"`
	Answers     []Answer `json:"answers"`
	Body        string   `json:"body"`
	Link        string   `json:"link"`
	QuestionID  int      `json:"question_id"`
	Score       int      `json:"score"`
	Title       string   `json:"title"`
}
