package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/kirillApanasiuk/movie-rating/internal/service"
	"github.com/kirillApanasiuk/movie-rating/model"
)

type HttpController struct {
	srv *service.Service
}

func New(ctrl *service.Service) *HttpController {
	return &HttpController{srv: ctrl}
}

func (h *HttpController) Handle(w http.ResponseWriter, req *http.Request) {
	recordID, recordType := model.RecordID(req.FormValue("id")), model.RecordType(req.FormValue("type"))
	if recordID == "" || recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		rsp, err := h.srv.GetAggregatedRating(req.Context(), &service.GetAggregatedRatingReq{
			RecordType: recordType,
			RecordID:   recordID,
		})
		var httpRsp struct {
			AggregatedRating float64 `json:"aggregated_rating"`
		}

		httpRsp.AggregatedRating = rsp.Total()

		if err != nil && errors.Is(err, service.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(httpRsp); err != nil {
			log.Printf("failed to encode response: %v\n", err)
		}

	case http.MethodPut:
		userID := model.UserId(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.srv.PutRating(req.Context(), recordID, recordType, &model.Rating{
			UserID: userID,
			Value:  model.RatingValue(v),
		}); err != nil {
			log.Printf("failed to encode response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
