package service

type DivarService interface {
}

type divarService struct{}

func NewDivarService() DivarService {
	return &divarService{}
}
