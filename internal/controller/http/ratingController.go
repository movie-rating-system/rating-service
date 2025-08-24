package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/kirillApanasiuk/movie-rating/domain/entity"
	"github.com/kirillApanasiuk/movie-rating/usecase/rating"
)

type HttpController struct {
	srv *rating.Service
}

func New(ctrl *rating.Service) *HttpController {
	return &HttpController{srv: ctrl}
}

func (h *HttpController) Handle(w http.ResponseWriter, req *http.Request) {
	recordID, recordType := entity.RecordID(req.FormValue("id")), entity.RecordType(req.FormValue("type"))
	if recordID == "" || recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		rsp, err := h.srv.GetAggregatedRating(req.Context(), &rating.GetAggregatedRatingReq{
			RecordType: recordType,
			RecordID:   recordID,
		})
		var httpRsp struct {
			AggregatedRating float64 `json:"aggregated_rating"`
		}

		httpRsp.AggregatedRating = rsp.Total()

		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(httpRsp); err != nil {
			log.Printf("failed to encode response: %v\n", err)
		}

	case http.MethodPut:
		userID := entity.UserId(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.srv.PutRating(req.Context(), recordID, recordType, &entity.Rating{
			userId: userID,
			value:  entity.RatingValue(v),
		}); err != nil {
			log.Printf("failed to encode response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
