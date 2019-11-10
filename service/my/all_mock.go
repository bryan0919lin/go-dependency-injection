package my

type AllServiceMock struct {
}

func NewAllServiceMock() *AllServiceMock {
	return &AllServiceMock{}
}

func (s *AllServiceMock) Welcome() string {
	return "Welcome"
}
