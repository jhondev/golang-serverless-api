package todo

func CreateTodo() map[string]string {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	return response
}
