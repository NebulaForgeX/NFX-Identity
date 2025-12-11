package profile_education

import (
	"context"
	"nfxid/events"
	educationCommands "nfxid/modules/auth/application/profile_education/commands"
	educationDomain "nfxid/modules/auth/domain/profile_education"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/safeexec"
)

func (s *Service) CreateEducation(ctx context.Context, cmd educationCommands.CreateEducationCmd) (*educationDomain.Education, error) {
	// 使用 domain factory 创建实体
	e, err := educationDomain.NewEducation(educationDomain.NewEducationParams{
		ProfileID: cmd.ProfileID,
		Editable:  cmd.Editable,
	})
	if err != nil {
		return nil, err
	}

	if err := s.educationRepo.Create(ctx, e); err != nil {
		return nil, err
	}

	logx.S().Infof("✅ Created education: %s (ID: %s, ProfileID: %s)", e.Editable().School, e.ID(), cmd.ProfileID)

	// 发布 Kafka 事件（异步）
	educationID := e.ID()
	profileID := cmd.ProfileID
	safeexec.SafeGo(func() error {
		// 发布成功事件
		_ = eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_SuccessEvent{
			Operation: "education.created",
			EntityID:  educationID.String(),
			Details: map[string]interface{}{
				"profile_id": profileID.String(),
				"school":     e.Editable().School,
			},
		})

		// 发布缓存清除事件
		return eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_EducationInvalidateCacheEvent{
			EducationID: educationID.String(),
			ProfileID:   profileID.String(),
			Operation:   "created",
		})
	})

	return e, nil
}
