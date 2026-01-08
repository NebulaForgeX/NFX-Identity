package event_retention_policies

import (
	"time"
)

func (erp *EventRetentionPolicy) Update(actionPattern *string, dataClassification *DataClassification, riskLevel *RiskLevel, retentionDays int, retentionAction RetentionAction, archiveLocation *string) error {
	if actionPattern != nil {
		erp.state.ActionPattern = actionPattern
	}
	if dataClassification != nil {
		erp.state.DataClassification = dataClassification
	}
	if riskLevel != nil {
		erp.state.RiskLevel = riskLevel
	}
	if retentionDays > 0 {
		erp.state.RetentionDays = retentionDays
	}
	if retentionAction != "" {
		validActions := map[RetentionAction]struct{}{
			RetentionActionArchive: {},
			RetentionActionDelete:  {},
			RetentionActionExport:  {},
		}
		if _, ok := validActions[retentionAction]; !ok {
			return ErrInvalidRetentionAction
		}
		erp.state.RetentionAction = retentionAction
	}
	if archiveLocation != nil {
		erp.state.ArchiveLocation = archiveLocation
	}
	erp.state.UpdatedAt = time.Now().UTC()
	return nil
}

func (erp *EventRetentionPolicy) UpdateStatus(status string) error {
	validStatuses := map[string]struct{}{
		"active":   {},
		"disabled": {},
	}
	if _, ok := validStatuses[status]; !ok {
		return ErrInvalidStatus
	}
	erp.state.Status = status
	erp.state.UpdatedAt = time.Now().UTC()
	return nil
}
