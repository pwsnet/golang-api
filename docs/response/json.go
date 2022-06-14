package response

type Result struct {
	Result interface{} `json:"result"`
} // @name Result

type Error struct {
	Error string `json:"error"`
} // @name Error
