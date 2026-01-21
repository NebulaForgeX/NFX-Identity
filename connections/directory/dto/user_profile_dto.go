package dto

import (
	"time"

	userprofilepb "nfxid/protos/gen/directory/user_profile"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateUserProfileDTO 创建用户资料的 DTO
type CreateUserProfileDTO struct {
	UserID       string
	Role         *string
	FirstName    *string
	LastName     *string
	Nickname     *string
	DisplayName  *string
	AvatarID     *string
	BackgroundID *string
	BackgroundIDs []string
	Bio          *string
	Birthday     *time.Time
	Age          *int32
	Gender       *string
	Location     *string
	Website      *string
	Github       *string
	SocialLinks  map[string]interface{}
	Skills       map[string]interface{}
}

// ToCreateUserProfileRequest 转换为 protobuf 请求
func (d *CreateUserProfileDTO) ToCreateUserProfileRequest() (*userprofilepb.CreateUserProfileRequest, error) {
	req := &userprofilepb.CreateUserProfileRequest{
		UserId:        d.UserID,
		Role:          d.Role,
		FirstName:     d.FirstName,
		LastName:      d.LastName,
		Nickname:      d.Nickname,
		DisplayName:   d.DisplayName,
		AvatarId:      d.AvatarID,
		BackgroundId:  d.BackgroundID,
		BackgroundIds: d.BackgroundIDs,
		Bio:           d.Bio,
		Age:           d.Age,
		Gender:        d.Gender,
		Location:      d.Location,
		Website:       d.Website,
		Github:        d.Github,
	}

	if d.SocialLinks != nil {
		socialLinksStruct, err := structpb.NewStruct(d.SocialLinks)
		if err != nil {
			return nil, err
		}
		req.SocialLinks = socialLinksStruct
	}

	if d.Skills != nil {
		skillsStruct, err := structpb.NewStruct(d.Skills)
		if err != nil {
			return nil, err
		}
		req.Skills = skillsStruct
	}

	if d.Birthday != nil {
		req.Birthday = timestamppb.New(*d.Birthday)
	}

	return req, nil
}
