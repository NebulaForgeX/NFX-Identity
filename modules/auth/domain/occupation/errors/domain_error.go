package errors

import "errors"

var (
	ErrOccupationNotFound        = errors.New("occupation not found")
	ErrOccupationCompanyRequired = errors.New("occupation company is required")
	ErrOccupationPositionRequired = errors.New("occupation position is required")
	ErrOccupationProfileIDRequired = errors.New("occupation profile_id is required")
)

