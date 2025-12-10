package profile_badge

import "time"

func (pb *ProfileBadge) EnsureEditable(e ProfileBadgeEditable) error {
	if err := e.Validate(); err != nil {
		return err
	}
	return nil
}

func (pb *ProfileBadge) Update(e ProfileBadgeEditable) error {
	if err := pb.EnsureEditable(e); err != nil {
		return err
	}

	pb.state.Editable = e
	pb.state.UpdatedAt = time.Now().UTC()
	return nil
}

