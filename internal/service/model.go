package service

import "github.com/kirillApanasiuk/movie-rating/model"

type GetAggregatedRatingReq struct {
	RecordType model.RecordType
	RecordID   model.RecordID
}
