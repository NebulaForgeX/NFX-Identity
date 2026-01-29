package actions

import "time"

func (a *Action) Update(key, service, status, name string, description *string) error {
	if a.DeletedAt() != nil {
		return ErrActionNotFound
	}
	if key == "" {
		return ErrActionKeyRequired
	}
	if name == "" {
		return ErrActionNameRequired
	}
	if service == "" {
		return ErrActionServiceRequired
	}
	a.state.Key = key
	a.state.Service = service
	if status != "" {
		a.state.Status = status
	}
	a.state.Name = name
	a.state.Description = description
	a.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (a *Action) Delete() error {
	if a.IsSystem() {
		return ErrSystemActionDelete
	}
	if a.DeletedAt() != nil {
		return nil
	}
	now := time.Now().UTC()
	a.state.DeletedAt = &now
	a.state.UpdatedAt = now
	return nil
}
