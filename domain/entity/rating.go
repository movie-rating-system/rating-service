package entity

// TODO: Add more record types when needed (books, products, etc.)
const (
	RecordTypeMovie RecordType = "microservices_in_go"
)

// TODO: Consider using more specific types with validation
type RecordType string
type RecordID string
type UserId string
type RatingValue int

// TODO: Add validation constants for business rules
// const (
//     MinRatingValue = 1
//     MaxRatingValue = 5
// )

// Rating represents a user rating for a record (rich domain model)
type Rating struct {
	id         string      // TODO: Use UUID instead of string
	recordType string      // TODO: Should be RecordType, not string
	userId     UserId      // TODO: Add validation that userId is not empty
	value      RatingValue // TODO: Add validation that value is 1-5
	// TODO: Add missing fields: recordID, review, createdAt, updatedAt
}

// TODO: CRITICAL - Add validation to constructor!
// This violates Clean Architecture - no business rules validation
func NewRating(id string, recordType string, userId UserId, value RatingValue) *Rating {
	// TODO: Replace with proper validation:
	// if value < 1 || value > 5 {
	//     return nil, domain.ErrInvalidRating
	// }
	// if userId == "" {
	//     return nil, domain.ErrInvalidRequest
	// }
	return &Rating{
		id:         id,
		recordType: recordType,
		userId:     userId,
		value:      value,
	}
}

// TODO: Add more getters for all private fields
func (r *Rating) GetValue() RatingValue {
	return r.value
}

// TODO: Add business methods (rich domain model):
// func (r *Rating) IsHighRating() bool { return r.value >= 4 }
// func (r *Rating) UpdateRating(newValue RatingValue) error { /* validation */ }
// func (r *Rating) GetID() string { return r.id }
// func (r *Rating) GetUserID() UserId { return r.userId }
