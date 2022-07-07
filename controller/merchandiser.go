package controller

import (
	"encoding/json"
	"net/http"
	"staffplan-reports/domain"
)

type MerchandiserController struct {
	UseCase domain.MerchandiserUseCase
}

func (m MerchandiserController) RouteStats(res http.ResponseWriter, req *http.Request) {
	if routeStats, err := m.UseCase.RouteStats(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)

	} else {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(routeStats)
	}
}
