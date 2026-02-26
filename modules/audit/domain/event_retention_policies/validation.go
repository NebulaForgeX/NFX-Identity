package event_retention_policies

import (
	auditErr "nfxidentity/errors/src/audit"
)

func (erp *EventRetentionPolicy) Validate() error {
	if erp.PolicyName() == "" {
		return auditErr.ErrPolicyNameRequired
	}
	if erp.RetentionDays() <= 0 {
		return auditErr.ErrRetentionDaysRequired
	}
	validActions := map[RetentionAction]struct{}{
		RetentionActionArchive: {},
		RetentionActionDelete:  {},
		RetentionActionExport:  {},
	}
	if _, ok := validActions[erp.RetentionAction()]; !ok {
		return auditErr.ErrInvalidRetentionAction
	}
	return nil
}
