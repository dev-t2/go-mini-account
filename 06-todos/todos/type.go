package todos

type Todo struct {
	Id         int    `json:"id"`
	IsComplete bool   `json:"isComplete"`
	Content    string `json:"content"`
	Order      int    `json:"order"`
}