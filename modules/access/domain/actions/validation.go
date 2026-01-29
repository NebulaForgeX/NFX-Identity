package actions

func (a *Action) Validate() error {
	if a.Key() == "" {
		return ErrActionKeyRequired
	}
	if a.Name() == "" {
		return ErrActionNameRequired
	}
	if a.Service() == "" {
		return ErrActionServiceRequired
	}
	return nil
}
