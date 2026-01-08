package image_types

func (it *ImageType) Validate() error {
	if it.Key() == "" {
		return ErrKeyRequired
	}
	return nil
}
