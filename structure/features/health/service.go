package health

type Service struct { }

func NewService() (*Service, error) {
	return &Service{
	}, nil
}