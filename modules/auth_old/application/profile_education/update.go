package profile_education

import (
	"context"
	"nfxid/events"
	educationCommands "nfxid/modules/auth/application/profile_education/commands"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/safeexec"
)

func (s *Service) UpdateEducation(ctx context.Context, cmd educationCommands.UpdateEducationCmd) error {
	e, err := s.educationRepo.Get.ByID(ctx, cmd.EducationID)
	if err != nil {
		return err
	}

	profileID := e.ProfileID()

	if err := e.EnsureEditable(cmd.Editable); err != nil {
		return err
	}

	if err := e.Update(cmd.Editable); err != nil {
		return err
	}

	if err := s.educationRepo.Update.Generic(ctx, e); err != nil {
		return err
	}

	logx.S().Infof("✅ Updated education: ID=%s", cmd.EducationID)

	// 发布 Kafka 事件（异步）
	educationID := cmd.EducationID
	safeexec.SafeGo(func() error {
		// 发布成功事件
		_ = eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_SuccessEvent{
			Operation: "education.updated",
			EntityID:  educationID.String(),
		})

		// 发布缓存清除事件
		return eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_EducationInvalidateCacheEvent{
			EducationID: educationID.String(),
			ProfileID:   profileID.String(),
			Operation:   "updated",
		})
	})

	return nil
}
