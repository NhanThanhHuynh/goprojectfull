package model

type Todo struct {
	ID         string        `json:"id" :"id" :"id"`
	Text       string        `json:"text" :"text" :"text"`
	Done       bool          `json:"done" :"done" :"done"`
	UserId     string        `:"user_id"`
	TodoDetail []*TodoDetail `:"todo_detail" :"todo_detail"`
}
