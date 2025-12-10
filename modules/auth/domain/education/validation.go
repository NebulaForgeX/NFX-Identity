package education

import educationErrors "nebulaid/modules/auth/domain/education/errors"

func (e *EducationEditable) Validate() error {
	if e.School == "" {
		return educationErrors.ErrEducationSchoolRequired
	}
	if len(e.School) > 200 {
		return educationErrors.ErrEducationSchoolRequired
	}
	return nil
}
