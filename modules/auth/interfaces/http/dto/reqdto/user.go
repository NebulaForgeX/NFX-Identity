package reqdto

import (
	userApp "nfxid/modules/auth/application/user"
	userAppCommands "nfxid/modules/auth/application/user/commands"
	userDomain "nfxid/modules/auth/domain/user"
	"nfxid/pkgs/query"
	"nfxid/pkgs/utils/ptr"

	"github.com/google/uuid"
)

type UserCreateRequestDTO struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Phone    *string `json:"phone,omitempty"`
	Password string  `json:"password" validate:"required,min=6"`
	Status   string  `json:"status,omitempty"`
}

type UserUpdateRequestDTO struct {
	ID       uuid.UUID `params:"id" validate:"required,uuid"`
	Username *string   `json:"username,omitempty"`
	Email    *string   `json:"email,omitempty"`
	Phone    *string   `json:"phone,omitempty"`
	Status   *string   `json:"status,omitempty"`
}

type UserLoginRequestDTO struct {
	Identifier string `json:"identifier" validate:"required"` // username, email 或 phone
	Password   string `json:"password" validate:"required"`
}

type UserRefreshTokenRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserByIDRequestDTO struct {
	ID uuid.UUID `params:"id" validate:"required,uuid"`
}

type UserQueryParamsDTO struct {
	Offset  *int     `query:"offset"`
	Limit   *int     `query:"limit"`
	Search  *string  `query:"search"`
	Status  *string  `query:"status"`
	RoleIDs []string `query:"role_ids"` // Changed to support multiple roles
	Sort    []string `query:"sort"`
}

func (r *UserCreateRequestDTO) ToCreateCmd() userApp.CreateUserCmd {
	editable := userDomain.UserEditable{
		Username: r.Username,
		Email:    r.Email,
		Phone:    r.Phone,
		Password: r.Password,
	}

	status := r.Status
	if status == "" {
		status = "active"
	}

	return userApp.CreateUserCmd{
		Editable: editable,
		Status:   status,
	}
}

func (r *UserUpdateRequestDTO) ToUpdateCmd() userApp.UpdateUserCmd {
	editable := userDomain.UserEditable{}
	if r.Username != nil {
		editable.Username = *r.Username
	}
	if r.Email != nil {
		editable.Email = *r.Email
	}
	if r.Phone != nil {
		editable.Phone = r.Phone
	}

	return userApp.UpdateUserCmd{
		UserID:   r.ID,
		Editable: editable,
	}
}

func (r *UserLoginRequestDTO) ToLoginCmd() userAppCommands.LoginCmd {
	return userAppCommands.LoginCmd{
		Identifier: r.Identifier,
		Password:   r.Password,
	}
}

func (r *UserRefreshTokenRequestDTO) ToRefreshCmd() userAppCommands.RefreshCmd {
	return userAppCommands.RefreshCmd{
		RefreshToken: r.RefreshToken,
	}
}

func (r *UserQueryParamsDTO) ToListQuery() userDomain.ListQuery {
	var status []string
	if r.Status != nil && *r.Status != "" {
		status = []string{*r.Status}
	}

	var roleIDs []uuid.UUID
	if len(r.RoleIDs) > 0 {
		roleIDs = make([]uuid.UUID, 0, len(r.RoleIDs))
		for _, idStr := range r.RoleIDs {
			if id, err := uuid.Parse(idStr); err == nil {
				roleIDs = append(roleIDs, id)
			}
		}
	}

	return userDomain.ListQuery{
		DomainPagination: query.DomainPagination{
			Offset: ptr.Deref(r.Offset),
			Limit:  ptr.Deref(r.Limit),
		},
		DomainSorts: query.ParseSortParams(r.Sort, map[string]userDomain.SortField{
			"created_time": userDomain.SortByCreatedTime,
			"username":     userDomain.SortByUsername,
			"email":        userDomain.SortByEmail,
		}),
		Search:  r.Search,
		Status:  status,
		RoleIDs: roleIDs,
	}
}
