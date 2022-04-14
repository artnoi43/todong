package internal

// TodoReqBody is used for creating/updating to-do
type TodoReqBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TodoDate    string `json:"todoDate"`
	Status      string `json:"status"`
}
