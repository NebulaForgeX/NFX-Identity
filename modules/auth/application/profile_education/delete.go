package profile_education

import (
	"context"
	"nfxid/events"
	"nfxid/pkgs/eventbus"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/safeexec"

	"github.com/google/uuid"
)

type DeleteEducationCmd struct {
	EducationID uuid.UUID
}

func (s *Service) DeleteEducation(ctx context.Context, cmd DeleteEducationCmd) error {
	e, err := s.educationRepo.Get.ByID(ctx, cmd.EducationID)
	if err != nil {
		return err
	}

	profileID := e.ProfileID()
	school := e.Editable().School

	if err := e.Delete(); err != nil {
		return err
	}

	if err := s.educationRepo.Update.Generic(ctx, e); err != nil {
		return err
	}

	logx.S().Infof("🗑️ Deleted education: %s (ID: %s)", school, cmd.EducationID)

	// 发布 Kafka 事件（异步）
	educationID := cmd.EducationID
	safeexec.SafeGo(func() error {
		// 发布成功事件
		_ = eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_SuccessEvent{
			Operation: "education.deleted",
			EntityID:  educationID.String(),
		})

		// 发布缓存清除事件
		return eventbus.PublishEvent(context.Background(), s.busPublisher, events.AuthToAuth_EducationInvalidateCacheEvent{
			EducationID: educationID.String(),
			ProfileID:   profileID.String(),
			Operation:   "deleted",
		})
	})

	return nil
}
