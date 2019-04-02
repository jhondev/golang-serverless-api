package todo

func DeleteTodo(todoID string) map[string]string {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully" + todoID
	return response
}
