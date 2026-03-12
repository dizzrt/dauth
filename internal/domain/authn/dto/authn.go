package dto

type LoginStatus uint8

const (
	LoginStatusSuccess LoginStatus = iota
	LoginStatusFailed
	LoginStatusLocked
)

type LoginResponse struct {
	Token             string
	Status            LoginStatus
	RemainingAttempts *int32
	RemainingLockTime *int64
}
