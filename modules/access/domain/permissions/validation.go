package permissions

func (p *Permission) Validate() error {
	if p.Key() == "" {
		return ErrPermissionKeyRequired
	}
	if p.Name() == "" {
		return ErrPermissionNameRequired
	}
	return nil
}
