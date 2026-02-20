package mapper

import (
	sessionAppResult "nfxid/modules/auth/application/sessions/results"
	sessionDomain "nfxid/modules/auth/domain/sessions"
	sessionpb "nfxid/protos/gen/auth/session"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// SessionROToProto 将 SessionRO 转换为 proto Session 消息
func SessionROToProto(v *sessionAppResult.SessionRO) *sessionpb.Session {
	if v == nil {
		return nil
	}

	session := &sessionpb.Session{
		Id:        v.ID.String(),
		SessionId: v.SessionID,
		UserId:    v.UserID.String(),
		CreatedAt: timestamppb.New(v.CreatedAt),
		LastSeenAt: timestamppb.New(v.LastSeenAt),
		ExpiresAt: timestamppb.New(v.ExpiresAt),
		UpdatedAt: timestamppb.New(v.UpdatedAt),
	}

	if v.AppID != nil {
		appIDStr := v.AppID.String()
		session.ApplicationId = &appIDStr
	}
	if v.ClientID != nil {
		session.ClientId = v.ClientID
	}
	if v.IP != nil {
		session.Ip = v.IP
	}
	if v.UAHash != nil {
		session.UaHash = v.UAHash
	}
	if v.DeviceID != nil {
		session.DeviceId = v.DeviceID
	}
	if v.DeviceFingerprint != nil {
		session.DeviceFingerprint = v.DeviceFingerprint
	}
	if v.DeviceName != nil {
		session.DeviceName = v.DeviceName
	}
	if v.RevokedAt != nil {
		session.RevokedAt = timestamppb.New(*v.RevokedAt)
	}
	if v.RevokeReason != nil {
		reason := revokeReasonToProto(*v.RevokeReason)
		session.RevokeReason = &reason
	}
	if v.RevokedBy != nil {
		session.RevokedBy = v.RevokedBy
	}

	return session
}

// SessionListROToProto 批量转换 SessionRO 到 proto Session
func SessionListROToProto(results []sessionAppResult.SessionRO) []*sessionpb.Session {
	sessions := make([]*sessionpb.Session, len(results))
	for i, v := range results {
		sessions[i] = SessionROToProto(&v)
	}
	return sessions
}

func revokeReasonToProto(reason sessionDomain.SessionRevokeReason) sessionpb.AuthSessionRevokeReason {
	switch reason {
	case sessionDomain.SessionRevokeReasonUserLogout:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_USER_LOGOUT
	case sessionDomain.SessionRevokeReasonAdminRevoke:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_ADMIN_REVOKE
	case sessionDomain.SessionRevokeReasonPasswordChanged:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_PASSWORD_CHANGED
	case sessionDomain.SessionRevokeReasonDeviceChanged:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_DEVICE_CHANGED
	case sessionDomain.SessionRevokeReasonAccountLocked:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_ACCOUNT_LOCKED
	case sessionDomain.SessionRevokeReasonSuspiciousActivity:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_SUSPICIOUS_ACTIVITY
	case sessionDomain.SessionRevokeReasonSessionExpired:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_SESSION_EXPIRED
	case sessionDomain.SessionRevokeReasonOther:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_OTHER
	default:
		return sessionpb.AuthSessionRevokeReason_AUTH_SESSION_REVOKE_REASON_UNSPECIFIED
	}
}
