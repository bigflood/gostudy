package todo

type Task struct {
	Desc string
	Done bool
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

func (p *InMem) List() ([]Task, error) {
	return p.Tasks, nil
}

func (p *InMem) Done(index int) error {
	p.Tasks[index].Done = true
	return nil
}
