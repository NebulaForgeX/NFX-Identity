package profile

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	state ProfileState
}

type ProfileState struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Editable    ProfileEditable
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type ProfileEditable struct {
	FirstName       *string
	LastName        *string
	Nickname        *string // 昵称（唯一）
	DisplayName     *string // 显示名称（不唯一）
	AvatarID        *uuid.UUID
	BackgroundID    *uuid.UUID
	BackgroundIds   []uuid.UUID // 背景图片ID数组（UUID数组）
	Bio             *string
	Phone           *string
	Birthday        *time.Time
	Age             *int // 年龄
	Gender          *string
	Location        *string // 位置信息（字符串格式 "C P C T"：Country Province City Timezone，空格分隔）
	Website         *string
	Github          *string                // GitHub 用户名或 URL
	SocialLinks     map[string]interface{} // 社交链接（JSONB）：{twitter, linkedin, instagram, youtube} (不包含 github/website)
	Preferences     map[string]interface{}
	Skills          map[string]int         // 技能：{"golang": 10, "python": 8, ...}
	PrivacySettings map[string]interface{} // 隐私设置（JSONB）
}

func (p *Profile) ID() uuid.UUID              { return p.state.ID }
func (p *Profile) UserID() uuid.UUID           { return p.state.UserID }
func (p *Profile) Editable() ProfileEditable   { return p.state.Editable }
func (p *Profile) CreatedAt() time.Time        { return p.state.CreatedAt }
func (p *Profile) UpdatedAt() time.Time        { return p.state.UpdatedAt }
func (p *Profile) DeletedAt() *time.Time       { return p.state.DeletedAt }
