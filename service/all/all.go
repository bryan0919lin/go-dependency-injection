package all

type AllService interface {
	Welcome() string
}

type AllServiceImpl struct {
}

func New() *AllServiceImpl {
	return &AllServiceImpl{}
}

func (s *AllServiceImpl) Welcome() string {
	return "Hello"
}
