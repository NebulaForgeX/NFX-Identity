package mapper

import (
	"nfxid/enums"
	"nfxid/modules/clients/domain/rate_limits"
	"nfxid/modules/clients/infrastructure/rdb/models"
)

// RateLimitDomainToModel 将 Domain RateLimit 转换为 Model RateLimit
func RateLimitDomainToModel(rl *rate_limits.RateLimit) *models.RateLimit {
	if rl == nil {
		return nil
	}

	return &models.RateLimit{
		ID:            rl.ID(),
		AppID:         rl.AppID(),
		LimitType:     rateLimitTypeDomainToEnum(rl.LimitType()),
		LimitValue:    rl.LimitValue(),
		WindowSeconds: rl.WindowSeconds(),
		Description:   rl.Description(),
		Status:        rl.Status(),
		CreatedAt:     rl.CreatedAt(),
		CreatedBy:     rl.CreatedBy(),
		UpdatedAt:     rl.UpdatedAt(),
		UpdatedBy:     rl.UpdatedBy(),
	}
}

// RateLimitModelToDomain 将 Model RateLimit 转换为 Domain RateLimit
func RateLimitModelToDomain(m *models.RateLimit) *rate_limits.RateLimit {
	if m == nil {
		return nil
	}

	state := rate_limits.RateLimitState{
		ID:            m.ID,
		AppID:         m.AppID,
		LimitType:     rateLimitTypeEnumToDomain(m.LimitType),
		LimitValue:    m.LimitValue,
		WindowSeconds: m.WindowSeconds,
		Description:   m.Description,
		Status:        m.Status,
		CreatedAt:     m.CreatedAt,
		CreatedBy:     m.CreatedBy,
		UpdatedAt:     m.UpdatedAt,
		UpdatedBy:     m.UpdatedBy,
	}

	return rate_limits.NewRateLimitFromState(state)
}

// RateLimitModelToUpdates 将 Model RateLimit 转换为更新字段映射
func RateLimitModelToUpdates(m *models.RateLimit) map[string]any {
	return map[string]any{
		models.RateLimitCols.LimitValue:    m.LimitValue,
		models.RateLimitCols.WindowSeconds: m.WindowSeconds,
		models.RateLimitCols.Description:   m.Description,
		models.RateLimitCols.Status:        m.Status,
		models.RateLimitCols.UpdatedAt:     m.UpdatedAt,
		models.RateLimitCols.UpdatedBy:     m.UpdatedBy,
	}
}

// 枚举转换辅助函数

// RateLimitTypeDomainToEnum 将 Domain RateLimitType 转换为 Enum RateLimitType（导出供其他包使用）
func RateLimitTypeDomainToEnum(rlt rate_limits.RateLimitType) enums.ClientsRateLimitType {
	return rateLimitTypeDomainToEnum(rlt)
}

func rateLimitTypeDomainToEnum(rlt rate_limits.RateLimitType) enums.ClientsRateLimitType {
	switch rlt {
	case rate_limits.RateLimitTypeRequestsPerSecond:
		return enums.ClientsRateLimitTypeRequestsPerSecond
	case rate_limits.RateLimitTypeRequestsPerMinute:
		return enums.ClientsRateLimitTypeRequestsPerMinute
	case rate_limits.RateLimitTypeRequestsPerHour:
		return enums.ClientsRateLimitTypeRequestsPerHour
	case rate_limits.RateLimitTypeRequestsPerDay:
		return enums.ClientsRateLimitTypeRequestsPerDay
	default:
		return enums.ClientsRateLimitTypeRequestsPerSecond
	}
}

func rateLimitTypeEnumToDomain(rlt enums.ClientsRateLimitType) rate_limits.RateLimitType {
	switch rlt {
	case enums.ClientsRateLimitTypeRequestsPerSecond:
		return rate_limits.RateLimitTypeRequestsPerSecond
	case enums.ClientsRateLimitTypeRequestsPerMinute:
		return rate_limits.RateLimitTypeRequestsPerMinute
	case enums.ClientsRateLimitTypeRequestsPerHour:
		return rate_limits.RateLimitTypeRequestsPerHour
	case enums.ClientsRateLimitTypeRequestsPerDay:
		return rate_limits.RateLimitTypeRequestsPerDay
	default:
		return rate_limits.RateLimitTypeRequestsPerSecond
	}
}
