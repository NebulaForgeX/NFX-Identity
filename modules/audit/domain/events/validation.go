package events

import (
	auditErr "nfxid/errors/src/audit"

	"github.com/google/uuid"
)

func (e *Event) Validate() error {
	if e.EventID() == "" {
		return auditErr.ErrEventIDRequired
	}
	if e.ActorType() == "" {
		return auditErr.ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[e.ActorType()]; !ok {
		return auditErr.ErrInvalidActorType
	}
	if e.ActorID() == uuid.Nil {
		return auditErr.ErrActorIDRequired
	}
	if e.Action() == "" {
		return auditErr.ErrActionRequired
	}
	if e.Result() == "" {
		return auditErr.ErrResultRequired
	}
	return nil
}
