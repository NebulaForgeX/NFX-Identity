package commands

import (
	"github.com/google/uuid"
)

type CreateActionCmd struct {
	Key         string
	Service     string
	Status      string
	Name        string
	Description *string
	IsSystem    bool
}

type UpdateActionCmd struct {
	ActionID    uuid.UUID
	Key         string
	Service     string
	Status      string
	Name        string
	Description *string
}

type DeleteActionCmd struct {
	ActionID uuid.UUID
}
