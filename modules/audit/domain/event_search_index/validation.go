package event_search_index

import "github.com/google/uuid"

func (esi *EventSearchIndex) Validate() error {
	if esi.EventID() == "" {
		return ErrEventIDRequired
	}
	if esi.ActorType() == "" {
		return ErrActorTypeRequired
	}
	validActorTypes := map[ActorType]struct{}{
		ActorTypeUser:    {},
		ActorTypeService: {},
		ActorTypeSystem:  {},
		ActorTypeAdmin:   {},
	}
	if _, ok := validActorTypes[esi.ActorType()]; !ok {
		return ErrInvalidActorType
	}
	if esi.ActorID() == uuid.Nil {
		return ErrActorIDRequired
	}
	if esi.Action() == "" {
		return ErrActionRequired
	}
	if esi.Result() == "" {
		return ErrResultRequired
	}
	return nil
}
