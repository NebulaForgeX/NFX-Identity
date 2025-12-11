package occupation

import occupationErrors "nfxid/modules/auth/domain/occupation/errors"

func (e *OccupationEditable) Validate() error {
	if e.Company == "" {
		return occupationErrors.ErrOccupationCompanyRequired
	}
	if len(e.Company) > 200 {
		return occupationErrors.ErrOccupationCompanyRequired
	}
	if e.Position == "" {
		return occupationErrors.ErrOccupationPositionRequired
	}
	if len(e.Position) > 100 {
		return occupationErrors.ErrOccupationPositionRequired
	}
	return nil
}
