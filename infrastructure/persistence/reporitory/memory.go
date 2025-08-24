package reporitory

import (
	"context"
	"errors"

	"github.com/kirillApanasiuk/movie-rating/domain"
	"github.com/kirillApanasiuk/movie-rating/domain/entity"
)

var ErrNotFound = errors.New("repository not found")

type recordRatingsMap map[entity.RecordID][]entity.Rating
type typesRecordMap map[entity.RecordType]recordRatingsMap
type Repository struct {
	data typesRecordMap
}

func NewRepository() *Repository {
	return &Repository{
		data: make(typesRecordMap),
	}
}

func (r *Repository) Get(_ context.Context, recordType entity.RecordType, id entity.RecordID) ([]entity.Rating, error) {
	recordsForType, ok := r.data[recordType]
	if !ok {
		return nil, domain.ErrRecordNotFound
	}

	ratings, ok := recordsForType[id]
	if !ok || len(ratings) == 0 {
		return nil, domain.ErrRecordNotFound
	}

	return ratings, nil
}

func (r *Repository) Put(_ context.Context, recordType entity.RecordType, id entity.RecordID, rating *entity.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = make(recordRatingsMap)
	}

	r.data[recordType][id] = append(r.data[recordType][id], *rating)

	return nil
}
