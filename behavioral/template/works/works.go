package works

import "fmt"

type Maider interface {
	GetName() string
	DoWork() string
}

func NewMaidTemplate(m Maider) *MaidTemplate {
	return &MaidTemplate{
		m,
	}
}

type MaidTemplate struct {
	Maider
}

func (m *MaidTemplate) Do() string {
	return m.DoWork()
}

type MaiderImpl struct {
	Name string
	Responsibility string
}

func (m *MaiderImpl) GetName() string {
	return m.Name
}

func (m *MaiderImpl) DoWork() string {
	return fmt.Sprintf("%s is doing %s", m.GetName(), m.Responsibility)
}