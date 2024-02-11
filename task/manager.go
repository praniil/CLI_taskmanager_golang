package task

type ManagerStruct struct {
	tasks    []*TaskStruct
	taskChan chan string
	done     chan bool
}

func NewMananger(taskChan chan string, done chan bool) *ManagerStruct {
	return &ManagerStruct{
		tasks: []*TaskStruct{},
		taskChan: taskChan,
		done: done,
	}
}

func (m *ManagerStruct) AddTask (description string) {
	m.tasks = append(m.tasks, NewTask(description))
}

func (m *ManagerStruct) ListernForTasks() {
	for{
		select{
		case description := <-m.taskChan:
			m.AddTask(description)
		case <- m.done:
			return	
		}
	}
}