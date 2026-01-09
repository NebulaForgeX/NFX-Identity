package actor_snapshots

func (as *ActorSnapshot) Update(displayName, email, clientName *string, snapshotData map[string]interface{}) error {
	if displayName != nil {
		as.state.DisplayName = displayName
	}
	if email != nil {
		as.state.Email = email
	}
	if clientName != nil {
		as.state.ClientName = clientName
	}
	if snapshotData != nil {
		as.state.SnapshotData = snapshotData
	}
	return nil
}

// Actor snapshots are typically immutable once created
// This update method is for rare cases where snapshot data needs correction
func (as *ActorSnapshot) UpdateSnapshotData(snapshotData map[string]interface{}) error {
	as.state.SnapshotData = snapshotData
	return nil
}
