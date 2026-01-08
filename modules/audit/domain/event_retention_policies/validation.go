package event_retention_policies

func (erp *EventRetentionPolicy) Validate() error {
	if erp.PolicyName() == "" {
		return ErrPolicyNameRequired
	}
	if erp.RetentionDays() <= 0 {
		return ErrRetentionDaysRequired
	}
	validActions := map[RetentionAction]struct{}{
		RetentionActionArchive: {},
		RetentionActionDelete:  {},
		RetentionActionExport:  {},
	}
	if _, ok := validActions[erp.RetentionAction()]; !ok {
		return ErrInvalidRetentionAction
	}
	return nil
}
