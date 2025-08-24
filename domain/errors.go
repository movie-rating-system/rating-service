package domain

import "errors"

// Domain errors - принадлежат бизнес-логике
var (
	ErrRatingNotFound = errors.New("rating not found")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInvalidRating  = errors.New("invalid rating value")
	ErrUserNotFound   = errors.New("user not found")
	ErrRecordNotFound = errors.New("record not found")
)
