package model

type Student struct {
	name    string
	surname string
}

func (s *Student) GetName() string {
	return s.name
}
func New(name string, surname string) Student {
	return Student{name, surname}
}
