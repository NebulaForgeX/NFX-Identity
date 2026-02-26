package actor_snapshots

import (
	auditErr "nfxidentity/errors/src/audit"

	"github.com/google/uuid"
)

func (as *ActorSnapshot) Validate() error {
	if as.ActorType() == "" {
		return auditErr.ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[as.ActorType()]; !ok {
		return auditErr.ErrInvalidActorType
	}
	if as.ActorID() == uuid.Nil {
		return auditErr.ErrActorIDRequired
	}
	return nil
}
