package identifier

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() string
}

type uuidGenerator struct{}

func NewUUIDGenerator() UUIDGenerator {
	return uuidGenerator{}
}

func (u uuidGenerator) New() string {
	return uuid.NewString()
}
