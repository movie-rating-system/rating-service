package reporitory

import (
	"context"
	"errors"

	"github.com/kirillApanasiuk/movie-rating/model"
)

var ErrNotFound = errors.New("reporitory not found")

type recordRatingsMap map[model.RecordID][]model.Rating
type typesRecordMap map[model.RecordType]recordRatingsMap
type Repository struct {
	data typesRecordMap
}

func NewRepository() *Repository {
	return &Repository{
		data: make(typesRecordMap),
	}
}

func (r *Repository) Get(_ context.Context, recordType model.RecordType, id model.RecordID) ([]model.Rating, error) {
	recordsForType, ok := r.data[recordType]
	if !ok {
		return nil, ErrNotFound
	}

	ratings, ok := recordsForType[id]
	if !ok || len(ratings) == 0 {
		return nil, ErrNotFound
	}

	return ratings, nil
}

func (r *Repository) Put(_ context.Context, recordType model.RecordType, id model.RecordID, rating *model.Rating) error {
	if _, ok := r.data[recordType]; !ok {
		r.data[recordType] = make(recordRatingsMap)
	}

	r.data[recordType][id] = append(r.data[recordType][id], *rating)

	return nil
}
