package model

const (
	RecordTypeMovie RecordType = "microservices_in_go"
)

type RecordType string
type RecordID string
type UserId string
type RatingValue int

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserId      `json:"userId"`
	Value      RatingValue `json:"value"`
}

type AggregatedRating struct {
	Ratings []Rating `json:"ratings"`
}

func (agr *AggregatedRating) Total() float64 {
	if agr.Ratings == nil {
		return 0
	}

	count := len(agr.Ratings)
	sum := 0.0
	for _, rating := range agr.Ratings {
		sum += float64(rating.Value)
	}
	return sum / float64(count)
}
