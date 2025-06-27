package commons

import "errors"

var (
	TypeBundle   = "bundle"
	TypeDiscount = "discount"
	TypeFreebie  = "freebie"

	// error message
	ErrNotFound = errors.New("data not found")
	ErrFailedGetData = errors.New("failed to get data")
	ErrUpdateData = errors.New("failed to update data")
	ErrFailedSaveData = errors.New("failed to save data")
)