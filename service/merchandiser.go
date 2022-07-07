package service

import (
	"staffplan-reports/domain"
	"staffplan-reports/model"
)

func NewMerchandiserUseCase(repo domain.MerchandiserRepo) domain.MerchandiserUseCase {
	return &merchandiserService{repo}
}

type merchandiserService struct {
	repo domain.MerchandiserRepo
}

func (m *merchandiserService) RouteStats() (*model.RouteStatsReponse, error) {
	return m.repo.GetRouteStats()
}
