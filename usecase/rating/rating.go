package rating

import (
	"context"
	"errors"

	"github.com/kirillApanasiuk/movie-rating/domain"
	"github.com/kirillApanasiuk/movie-rating/domain/entity"
)

var ErrNotFound = errors.New("rating not found for a record")
var InvalidArgumentReq = errors.New("invalid argument")

type ratingRepository interface {
	Get(ctx context.Context, recordType entity.RecordType, recordID entity.RecordID) ([]entity.Rating, error)
	Put(ctx context.Context, recordType entity.RecordType, recordID entity.RecordID, rating *entity.Rating) error
}

type Service struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (r *Service) GetAggregatedRating(ctx context.Context, req *GetAggregatedRatingReq) (*AggregatedRating, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	recordType := req.RecordType
	recordID := req.RecordID

	ratings, err := r.repo.Get(ctx, recordType, recordID)
	if err != nil && errors.Is(err, domain.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	return &AggregatedRating{Ratings: ratings}, nil
}

func (r *Service) PutRating(ctx context.Context, recordID entity.RecordID, recordType entity.RecordType, rating *entity.Rating) error {
	if recordID == "" || recordType == "" {
		return InvalidArgumentReq
	}

	return r.repo.Put(ctx, recordType, recordID, rating)
}
