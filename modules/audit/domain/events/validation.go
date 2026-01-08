package events

import "github.com/google/uuid"

func (e *Event) Validate() error {
	if e.EventID() == "" {
		return ErrEventIDRequired
	}
	if e.ActorType() == "" {
		return ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[e.ActorType()]; !ok {
		return ErrInvalidActorType
	}
	if e.ActorID() == uuid.Nil {
		return ErrActorIDRequired
	}
	if e.Action() == "" {
		return ErrActionRequired
	}
	if e.Result() == "" {
		return ErrResultRequired
	}
	return nil
}
