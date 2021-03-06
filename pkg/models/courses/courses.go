package courses

type Task struct {
	ID				int64 	`json:"id"`
	Name			string	`json:"name"`
	TaskHeaderID	string 	`json:"task_header_id"`
	ImgURL			string	`json:"imgURL"`
	contentURL		string	`json:"contentURL"`
}

type TaskHeader struct {
	ID				int64	`json:"id"`
	CourseID		string	`json:"course_id"`
	Name			string	`json:"name"`
}

type Course struct {
	ID 				int64 	`json:"id"`
	Name			string 	`json:"name"`
	ImgURL			string 	`json:"imgURL"`
}

type UsersInCourse struct {
	UserID			string 	`json:"user_id"`
	CourseID		string	`json:"course_id"`
}

func TasksToMap(tasks []*Task) map[string][]*Task {
	m := make(map[string][]*Task)

	for _, t := range tasks {
		if len(m[t.TaskHeaderID]) == 0 {
			m[t.TaskHeaderID] = []*Task{}
		}

		m[t.TaskHeaderID] = append(m[t.TaskHeaderID], t)
	}

	return m
}