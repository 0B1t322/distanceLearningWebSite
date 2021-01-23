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

type UserInCourse struct {
	UserID			int64 	`json:"user_id"`
	CourseID		int64	`json:"course_id"`
}