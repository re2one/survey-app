package response

type FullQuestionsResponse struct {
	Questions []*FullQuestion `json:"questions"`
	Finished  bool            `json:"finished"`
}

type FullQuestion struct {
	QuestionId uint   `json:"questionId"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Answered   bool   `json:"answered"`
	Example    string `json:"example"`
}
