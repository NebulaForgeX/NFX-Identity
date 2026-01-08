package groups

import (
	"time"

	"github.com/google/uuid"
)

type GroupType string

const (
	GroupTypeDepartment GroupType = "department"
	GroupTypeTeam       GroupType = "team"
	GroupTypeGroup      GroupType = "group"
	GroupTypeOther      GroupType = "other"
)

type Group struct {
	state GroupState
}

type GroupState struct {
	ID            uuid.UUID
	GroupID       string
	TenantID      uuid.UUID
	Name          string
	Type          GroupType
	ParentGroupID *uuid.UUID
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     *uuid.UUID
	DeletedAt     *time.Time
	Metadata      map[string]interface{}
}

func (g *Group) ID() uuid.UUID                    { return g.state.ID }
func (g *Group) GroupID() string                  { return g.state.GroupID }
func (g *Group) TenantID() uuid.UUID              { return g.state.TenantID }
func (g *Group) Name() string                     { return g.state.Name }
func (g *Group) Type() GroupType                  { return g.state.Type }
func (g *Group) ParentGroupID() *uuid.UUID        { return g.state.ParentGroupID }
func (g *Group) Description() *string             { return g.state.Description }
func (g *Group) CreatedAt() time.Time             { return g.state.CreatedAt }
func (g *Group) UpdatedAt() time.Time             { return g.state.UpdatedAt }
func (g *Group) CreatedBy() *uuid.UUID            { return g.state.CreatedBy }
func (g *Group) DeletedAt() *time.Time            { return g.state.DeletedAt }
func (g *Group) Metadata() map[string]interface{} { return g.state.Metadata }
