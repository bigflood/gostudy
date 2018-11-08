package store

type Task struct {
	Desc string
	Done bool
}

type Filter struct {
	Done *bool
}

func NewInMem() *InMem {
	return &InMem{}
}

type InMem struct {
	Tasks []Task
}

func (p *InMem) Add(desc string) error {
	p.Tasks = append(p.Tasks, Task{Desc: desc})
	return nil
}

func (p *InMem) List(filter Filter) ([]Task, error) {
	if filter.Done == nil {
		return p.Tasks, nil
	}

	tasks := make([]Task, 0, len(p.Tasks))
	for _, task := range p.Tasks {
		if *filter.Done != task.Done {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (p *InMem) Done(index int) error {
	p.Tasks[index].Done = true
	return nil
}
