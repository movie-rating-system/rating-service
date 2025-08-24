package grpc

import (
	"context"

	"github.com/kirillApanasiuk/movie-rating/gen"
	"github.com/kirillApanasiuk/movie-rating/internal/service"
	"github.com/kirillApanasiuk/movie-rating/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	unspecified = 0
)

type GrpcController struct {
	gen.UnimplementedRatingServiceServer
	svc *service.Service
}

func New(svc *service.Service) *GrpcController {
	return &GrpcController{
		svc: svc,
	}
}

func (h *GrpcController) GetAggregatedRating(ctx context.Context, req *gen.GetAggregatedRatingRequest) (*gen.GetAggregatedRatingResponse, error) {
	if req == nil || req.RecordId == "" || req.RecordType == unspecified {
		return nil, status.Errorf(codes.InvalidArgument, "nil")
	}

	rsp, err := h.svc.GetAggregatedRating(ctx, &service.GetAggregatedRatingReq{
		RecordType: model.RecordType(req.RecordType),
		RecordID:   model.RecordID(req.RecordId),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	grpcRsp := &gen.GetAggregatedRatingResponse{
		RatingValue: rsp.Total(),
	}

	return grpcRsp, nil
}

func (h *GrpcController) PutRating(ctx context.Context, req *gen.PutRatingRequest) (*gen.PutRatingResponse, error) {
	//TODO implement me
	panic("implement me")
}
