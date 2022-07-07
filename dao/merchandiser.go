package dao

import (
	"database/sql"
	"staffplan-reports/domain"
	"staffplan-reports/model"
)

func NewMerchandiserRepo(db *sql.DB) domain.MerchandiserRepo {
	return &merchandiserDao{db}
}

type merchandiserDao struct {
	db *sql.DB
}

func (m *merchandiserDao) GetRouteStats() (*model.RouteStatsReponse, error) {
	var routeStat model.RouteStatsReponse
	err := m.db.QueryRow("select * from report_route_stats").Scan(
		&routeStat.Total,
		&routeStat.Ongoing,
		&routeStat.Completed,
	)

	return &routeStat, err
}
