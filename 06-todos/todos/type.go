package todos

type Todo struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	IsComplete bool   `json:"isComplete"`
}