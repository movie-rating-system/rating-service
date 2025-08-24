package service // the same as controller

import (
	"context"
	"errors"

	"github.com/kirillApanasiuk/movie-rating/internal/reporitory"
	"github.com/kirillApanasiuk/movie-rating/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrNotFound = errors.New("rating not found for a record")
var InvalidArgumentReq = errors.New("invalid argument")

type ratingRepository interface {
	Get(ctx context.Context, recordType model.RecordType, recordID model.RecordID) ([]model.Rating, error)
	Put(ctx context.Context, recordType model.RecordType, recordID model.RecordID, rating *model.Rating) error
}

type Service struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (r *Service) GetAggregatedRating(ctx context.Context, req *GetAggregatedRatingReq) (*model.AggregatedRating, error) {
	recordType := req.RecordType
	recordID := req.RecordID

	if req == nil || recordType == "" || recordID == "" {
		return nil, status.Error(codes.InvalidArgument, InvalidArgumentReq.Error())
	}

	ratings, err := r.repo.Get(ctx, recordType, recordID)
	if err != nil && errors.Is(err, reporitory.ErrNotFound) {
		return nil, ErrNotFound
	}

	return &model.AggregatedRating{Ratings: ratings}, nil
}

func (r *Service) PutRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType, rating *model.Rating) error {
	if recordID == "" || recordType == "" {
		return status.Error(codes.InvalidArgument, InvalidArgumentReq.Error())
	}
	return r.repo.Put(ctx, recordType, recordID, rating)
}
