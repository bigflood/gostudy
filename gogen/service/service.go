package service

//go:generate asdfasdfasfd.exe -i xxxx.templ yyyy.go

type Service interface {
	Add(a, b int) int
	Mul(a, b int) int
}

func New() Service {
	return new(serviceImpl)
}

type serviceImpl struct {
}

func (svc *serviceImpl) Add(a, b int) int {
	return a + b
}

func (svc *serviceImpl) Mul(a, b int) int {
	return a * b
}
