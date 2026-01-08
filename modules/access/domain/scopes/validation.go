package scopes

func (s *Scope) Validate() error {
	if s.ScopeKey() == "" {
		return ErrScopeRequired
	}
	return nil
}
