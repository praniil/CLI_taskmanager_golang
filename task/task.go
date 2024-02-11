package task

type TaskStruct struct {
	Description string
}

func NewTask (description string) *TaskStruct{
	return &TaskStruct{
		Description: description,
	}
}

func (t *TaskStruct) String() string{
	return t.Description
}