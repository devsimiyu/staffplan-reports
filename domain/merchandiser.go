package domain

import "staffplan-reports/model"

type MerchandiserUseCase interface {
	RouteStats() (*model.RouteStatsReponse, error)
}

type MerchandiserRepo interface {
	GetRouteStats() (*model.RouteStatsReponse, error)
}
