package actions

import (
	"time"

	"github.com/google/uuid"
)

type NewActionParams struct {
	Key         string
	Service     string
	Status      string
	Name        string
	Description *string
	IsSystem    bool
}

func NewAction(p NewActionParams) (*Action, error) {
	if err := validateActionParams(p); err != nil {
		return nil, err
	}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	status := p.Status
	if status == "" {
		status = "active"
	}
	now := time.Now().UTC()
	return NewActionFromState(ActionState{
		ID:          id,
		Key:         p.Key,
		Service:     p.Service,
		Status:      status,
		Name:        p.Name,
		Description: p.Description,
		IsSystem:    p.IsSystem,
		CreatedAt:   now,
		UpdatedAt:   now,
	}), nil
}

func NewActionFromState(st ActionState) *Action {
	return &Action{state: st}
}

func validateActionParams(p NewActionParams) error {
	if p.Key == "" {
		return ErrActionKeyRequired
	}
	if p.Name == "" {
		return ErrActionNameRequired
	}
	if p.Service == "" {
		return ErrActionServiceRequired
	}
	return nil
}
