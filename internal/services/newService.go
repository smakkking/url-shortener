package services

func NewService(storage Storage) *Service {
	return &Service{
		urlStorage: storage,
	}
}
