package rating

import (
	"github.com/kirillApanasiuk/movie-rating/domain/entity"
)

type GetAggregatedRatingReq struct {
	RecordType entity.RecordType
	RecordID   entity.RecordID
}

func (req *GetAggregatedRatingReq) Validate() error {
	if req == nil || req.RecordType == "" || req.RecordID == "" {
		return InvalidArgumentReq
	}
	return nil
}

type AggregatedRating struct {
	Ratings []entity.Rating
}

func (agr *AggregatedRating) Total() float64 {
	if agr.Ratings == nil {
		return 0
	}

	count := len(agr.Ratings)
	sum := 0.0
	for _, rating := range agr.Ratings {
		sum += float64(rating.GetValue())
	}
	return sum / float64(count)
}
