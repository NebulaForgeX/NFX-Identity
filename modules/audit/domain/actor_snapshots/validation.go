package actor_snapshots

import "github.com/google/uuid"

func (as *ActorSnapshot) Validate() error {
	if as.ActorType() == "" {
		return ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[as.ActorType()]; !ok {
		return ErrInvalidActorType
	}
	if as.ActorID() == uuid.Nil {
		return ErrActorIDRequired
	}
	return nil
}
