package login_attempts

func (la *LoginAttempt) Validate() error {
	if la.Identifier() == "" {
		return ErrIdentifierRequired
	}
	return nil
}
