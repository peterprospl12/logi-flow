package order

import "errors"

var (
	ErrEmptyShipperID   = errors.New("shipper ID cannot be empty")
	ErrEmptyOrigin      = errors.New("origin cannot be empty")
	ErrEmptyDestination = errors.New("destination cannot be empty")
	ErrInvalidPrice     = errors.New("initial price must be greater than zero")
)

type Status string

const (
	Created    Status = "created"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
	Cancelled  Status = "cancelled"
)

type TransportType string

const (
	Road       TransportType = "road"
	Air        TransportType = "air"
	Sea        TransportType = "sea"
	Rail       TransportType = "rail"
	Intermodal TransportType = "intermodal"
)
