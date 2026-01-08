package badges

func (b *Badge) Validate() error {
	if b.Name() == "" {
		return ErrNameRequired
	}
	return nil
}
