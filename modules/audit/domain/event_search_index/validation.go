package event_search_index

import (
	auditErr "nfxid/errors/src/audit"

	"github.com/google/uuid"
)

func (esi *EventSearchIndex) Validate() error {
	if esi.EventID() == "" {
		return auditErr.ErrEventIDRequired
	}
	if esi.ActorType() == "" {
		return auditErr.ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[esi.ActorType()]; !ok {
		return auditErr.ErrInvalidActorType
	}
	if esi.ActorID() == uuid.Nil {
		return auditErr.ErrActorIDRequired
	}
	if esi.Action() == "" {
		return auditErr.ErrActionRequired
	}
	if esi.Result() == "" {
		return auditErr.ErrResultRequired
	}
	return nil
}
