package todo

func GetAllTodos() []Todo {
	return []Todo{
		{
			Slug:  "slug 1",
			Title: "Hello world 1",
			Body:  "Heloo world from planet earth 1",
		},
		{
			Slug:  "slug 2",
			Title: "Hello world 2",
			Body:  "Heloo world from planet earth 2",
		},
	}
}

func GetAToDo(todoID string) Todo {
	return Todo{
		Slug:  todoID,
		Title: "Hello world",
		Body:  "Heloo world from planet earth",
	}
}
