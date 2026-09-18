package profile

import (
	"time"
)

// * =============================== Authority avatar behavior =============================== !//
func (a *AuthorityProfileAvatar) SetIsActive(active bool) {
	a.state.IsActive = active
	a.state.UpdatedAt = time.Now().UTC()
}

// * =============================== Forger avatar behavior =============================== !//
func (a *ForgerProfileAvatar) SetIsActive(active bool) {
	a.state.IsActive = active
	a.state.UpdatedAt = time.Now().UTC()
}
