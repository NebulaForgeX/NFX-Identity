package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"time"

	systemErr "nfxid/errors/src/system"
	bootstrapCommands "nfxid/modules/system/application/bootstrap/commands"
	systemStateDomain "nfxid/modules/system/domain/system_state"
	"nfxid/pkgs/logx"

	"github.com/google/uuid"
)

// BootstrapInit 系统初始化
// 流程：
// 1. 检查系统是否已经初始化
// 2. 检查所有服务的健康状态（包括基础设施：数据库、Redis等）
// 3. 清空所有服务的 schema（清空所有表数据，不删除表），确保 system_state 只有一条记录
// 4. 创建 system_state 记录（initialized = false），表示初始化开始
// 5. 通过 gRPC 调用其他服务初始化基础数据
// 6. 更新 metadata 并标记系统为已初始化
func (s *Service) BootstrapInit(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) error {
	logx.S().Info("🚀 Starting system bootstrap initialization...")
	// 步骤 1: 检查系统是否已经初始化
	if err := s.checkSystemInitialized(ctx); err != nil {
		return err
	}
	// 步骤 2: 检查所有服务的健康状态（包括基础设施：数据库、Redis等）
	if err := s.checkAllServicesHealth(ctx); err != nil {
		return err
	}
	// 步骤 3: 清空所有服务的 schema（清空所有表数据，不删除表） 注意：这会清空 system_state 表，确保只有一条记录
	schemaClearResults, err := s.clearAllSchemas(ctx)
	if err != nil {
		return err
	}
	// 步骤 3.1: 清空所有存储（如 Image 的 data 目录下所有图片）
	if err := s.clearAllStorages(ctx); err != nil {
		return err
	}
	// 步骤 4: 创建 system_state 记录（initialized = false），表示初始化开始
	systemState, err := s.createInitialSystemState(ctx, cmd)
	if err != nil {
		return err
	}
	// 步骤 5: 初始化各个服务的基础数据
	initializedServices := []string{}
	// 5.1 初始化 Directory 服务 - 创建第一个系统管理员用户
	logx.S().Info("📦 Initializing Directory service - creating admin user...")
	adminUserID, err := s.initDirectoryService(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to initialize directory service: %w", err)
	}
	initializedServices = append(initializedServices, "directory")
	logx.S().Infof("✅ Directory service initialized - admin user ID: %s", adminUserID)

	// 5.2 初始化 Access 服务 - 创建初始角色和权限
	logx.S().Info("📦 Initializing Access service - creating roles and permissions...")
	adminRoleID, err := s.initAccessService(ctx, adminUserID)
	if err != nil {
		return fmt.Errorf("failed to initialize access service: %w", err)
	}
	initializedServices = append(initializedServices, "access")
	logx.S().Infof("✅ Access service initialized - admin role ID: %s", adminRoleID)

	// 5.3 初始化 Auth 服务 - 创建用户凭证
	logx.S().Info("📦 Initializing Auth service - creating user credentials...")
	if err := s.initAuthService(ctx, adminUserID, cmd.AdminPassword); err != nil {
		return fmt.Errorf("failed to initialize auth service: %w", err)
	}
	initializedServices = append(initializedServices, "auth")
	logx.S().Info("✅ Auth service initialized")
	// 步骤 6: 更新 metadata 并标记系统为已初始化
	if err := s.finalizeSystemState(ctx, systemState, adminUserID, adminRoleID, initializedServices, schemaClearResults, cmd.Version); err != nil {
		return err
	}
	logx.S().Info("🎉 System bootstrap initialization completed successfully!")
	return nil
}

// checkSystemInitialized 检查系统是否已经初始化
func (s *Service) checkSystemInitialized(ctx context.Context) error {
	latestState, err := s.systemStateRepo.Get.Latest(ctx)
	if err != nil {
		if errors.Is(err, systemErr.ErrSystemStateNotFound) {
			logx.S().Info("ℹ️  No system state record found, proceeding with initialization...")
			return nil
		}
		return fmt.Errorf("failed to get latest system state: %w", err)
	}

	if latestState.Initialized() {
		return fmt.Errorf("system is already initialized")
	}

	return nil
}

// createInitialSystemState 创建 system_state 记录（initialized = false），表示初始化开始
func (s *Service) createInitialSystemState(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) (*systemStateDomain.SystemState, error) {
	now := time.Now().UTC()
	initialMetadata := map[string]interface{}{
		"bootstrap_started_at": now.Format(time.RFC3339),
		"admin_username":       cmd.AdminUsername,
		"services_initialized": []string{},
	}

	systemState, err := systemStateDomain.NewSystemState(systemStateDomain.NewSystemStateParams{
		Initialized:           false, // 默认未初始化
		InitializedAt:         nil,
		InitializationVersion: nil, // 在调用 Initialize 时设置
		LastResetAt:           nil,
		LastResetBy:           nil,
		ResetCount:            0, // 默认值
		Metadata:              initialMetadata,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create system state: %w", err)
	}

	// 保存初始状态记录
	if err := s.systemStateRepo.Create.New(ctx, systemState); err != nil {
		return nil, fmt.Errorf("failed to save initial system state: %w", err)
	}

	logx.S().Info("✅ Created initial system state record (initialized=false)")
	return systemState, nil
}

// checkAllServicesHealth 检查所有服务的健康状态（包括基础设施：数据库、Redis等）
func (s *Service) checkAllServicesHealth(ctx context.Context) error {
	logx.S().Info("⏳ Checking health of all 8 services (including infrastructure: database, Redis)...")
	maxRetries := 10
	retryInterval := 2 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		healthResults, err := s.grpcClients.CheckAllServicesHealth(ctx)
		if err != nil {
			logx.S().Warnf("Health check attempt %d/%d failed: %v", attempt, maxRetries, err)
			if attempt < maxRetries {
				time.Sleep(retryInterval)
				continue
			}
			return fmt.Errorf("failed to check service health after %d attempts: %w", maxRetries, err)
		}

		// 检查所有服务及其基础设施是否健康
		allHealthy := true
		unhealthyServices := []string{}
		unhealthyInfra := []string{}

		for serviceName, healthResp := range healthResults {
			if !healthResp.Healthy {
				allHealthy = false
				unhealthyServices = append(unhealthyServices, serviceName)
				logx.S().Warnf("❌ Service %s is not healthy", serviceName)
				continue
			}

			// 检查基础设施
			if healthResp.Infrastructure != nil {
				infraIssues := []string{}
				if healthResp.Infrastructure.Database != nil && !healthResp.Infrastructure.Database.Healthy {
					errorMsg := "unknown"
					if healthResp.Infrastructure.Database.ErrorMessage != nil {
						errorMsg = *healthResp.Infrastructure.Database.ErrorMessage
					}
					infraIssues = append(infraIssues, fmt.Sprintf("database: %s", errorMsg))
				}
				if healthResp.Infrastructure.Redis != nil && !healthResp.Infrastructure.Redis.Healthy {
					errorMsg := "unknown"
					if healthResp.Infrastructure.Redis.ErrorMessage != nil {
						errorMsg = *healthResp.Infrastructure.Redis.ErrorMessage
					}
					infraIssues = append(infraIssues, fmt.Sprintf("redis: %s", errorMsg))
				}

				if len(infraIssues) > 0 {
					allHealthy = false
					unhealthyInfra = append(unhealthyInfra, fmt.Sprintf("%s (%s)", serviceName, fmt.Sprint(infraIssues)))
					logx.S().Warnf("⚠️  Service %s is running but infrastructure unhealthy: %v", serviceName, infraIssues)
				} else {
					logx.S().Infof("✅ Service %s is healthy (service + database + redis)", serviceName)
				}
			} else {
				logx.S().Infof("✅ Service %s is healthy (no infrastructure info)", serviceName)
			}
		}

		if allHealthy {
			logx.S().Info("✅ All 8 services and their infrastructure (database, Redis) are healthy!")
			break
		}

		if attempt < maxRetries {
			allIssues := append(unhealthyServices, unhealthyInfra...)
			logx.S().Warnf("⚠️  Some services or infrastructure are not healthy (attempt %d/%d): %v. Retrying in %v...",
				attempt, maxRetries, allIssues, retryInterval)
			time.Sleep(retryInterval)
		} else {
			allIssues := append(unhealthyServices, unhealthyInfra...)
			return fmt.Errorf("some services or infrastructure are not healthy after %d attempts: %v", maxRetries, allIssues)
		}
	}

	return nil
}

// clearAllSchemas 清空所有服务的 schema（清空所有表数据，不删除表）
func (s *Service) clearAllSchemas(ctx context.Context) (map[string]int, error) {
	logx.S().Info("🧹 Clearing all schemas - removing all table data (keeping table structure)...")
	clearResults, err := s.grpcClients.ClearAllSchemas(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to clear schemas: %w", err)
	}
	// 检查清空结果并收集清除的表数量
	allCleared := true
	failedServices := []string{}
	schemaClearResults := make(map[string]int)
	for serviceName, result := range clearResults {
		if !result.Success {
			allCleared = false
			errMsg := "unknown error"
			if result.ErrorMessage != nil {
				errMsg = *result.ErrorMessage
			}
			failedServices = append(failedServices, fmt.Sprintf("%s: %s", serviceName, errMsg))
			logx.S().Warnf("⚠️  Failed to clear schema for %s: %s", serviceName, errMsg)
		} else {
			schemaClearResults[serviceName] = int(result.TablesCleared)
			logx.S().Infof("✅ Cleared schema for %s: %d tables cleared", serviceName, result.TablesCleared)
		}
	}

	if !allCleared {
		return nil, fmt.Errorf("failed to clear schemas for some services: %v", failedServices)
	}

	logx.S().Info("✅ All schemas cleared successfully!")
	return schemaClearResults, nil
}

// clearAllStorages 清空所有服务的存储（如 Image 的 data 目录下所有图片文件；NewGRPCClients 已保证 ImageClient 非 nil）
func (s *Service) clearAllStorages(ctx context.Context) error {
	logx.S().Info("🧹 Clearing all storages (e.g. image data files)...")
	success, errMsg, err := s.grpcClients.ImageClient.Image.ClearStorageData(ctx)
	if err != nil {
		return fmt.Errorf("failed to clear image storage: %w", err)
	}
	if !success {
		return fmt.Errorf("image clear storage failed: %s", errMsg)
	}
	logx.S().Info("✅ All storages cleared (image data)")
	return nil
}

// finalizeSystemState 更新 metadata 并标记系统为已初始化
func (s *Service) finalizeSystemState(
	ctx context.Context,
	systemState *systemStateDomain.SystemState,
	adminUserID, adminRoleID uuid.UUID,
	initializedServices []string,
	schemaClearResults map[string]int,
	version string,
) error {
	// 更新 metadata
	updatedMetadata := map[string]interface{}{
		"bootstrap_started_at":   systemState.Metadata()["bootstrap_started_at"],
		"admin_username":         systemState.Metadata()["admin_username"],
		"services_initialized":   initializedServices,
		"admin_user_id":          adminUserID.String(),
		"admin_role_id":          adminRoleID.String(),
		"bootstrap_completed_at": time.Now().UTC().Format(time.RFC3339),
		"schema_clear_results":   schemaClearResults,
	}

	if err := systemState.UpdateMetadata(updatedMetadata); err != nil {
		return fmt.Errorf("failed to update system state metadata: %w", err)
	}

	// 标记系统为已初始化
	logx.S().Info("✅ All services initialized, marking system as initialized...")
	if err := systemState.Initialize(version); err != nil {
		return fmt.Errorf("failed to initialize system state: %w", err)
	}

	// 一次性保存所有更新
	if err := s.systemStateRepo.Update.Generic(ctx, systemState); err != nil {
		return fmt.Errorf("failed to save finalized system state: %w", err)
	}

	return nil
}

// initDirectoryService 初始化 Directory 服务
// 创建第一个系统管理员用户及其关联数据
func (s *Service) initDirectoryService(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) (uuid.UUID, error) {
	// 1. 创建用户
	logx.S().Info("🔍 Creating admin user...")
	userIDStr, err := s.grpcClients.DirectoryClient.User.CreateUser(ctx, cmd.AdminUsername, "active", true)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create admin user: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID returned: %w", err)
	}

	// 2. 创建用户邮箱（如果提供了）
	if cmd.AdminEmail != nil && *cmd.AdminEmail != "" {
		logx.S().Info("🔍 Creating admin user email...")
		_, err := s.grpcClients.DirectoryClient.UserEmail.CreateUserEmailDefault(ctx, userID.String(), *cmd.AdminEmail)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to create admin user email: %w", err)
		}
	}

	// 3. 创建用户手机（如果提供了）
	if cmd.AdminPhone != nil && *cmd.AdminPhone != "" {
		countryCode := ""
		if cmd.AdminCountryCode != nil {
			countryCode = *cmd.AdminCountryCode
		}
		logx.S().Info("🔍 Creating admin user phone...")
		_, err := s.grpcClients.DirectoryClient.UserPhone.CreateUserPhoneDefault(ctx, userID.String(), *cmd.AdminPhone, countryCode)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to create admin user phone: %w", err)
		}
	}

	// 4. 创建用户资料（创建空的，后续可以更新）
	logx.S().Info("🔍 Creating admin user profile...")
	_, err = s.grpcClients.DirectoryClient.UserProfile.CreateUserProfileDefault(ctx, userID.String())
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create admin user profile: %w", err)
	}

	// 5. 创建用户偏好（创建空的，使用默认值）
	_, err = s.grpcClients.DirectoryClient.UserPreference.CreateUserPreferenceDefault(ctx, userID.String())
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create admin user preference: %w", err)
	}

	return userID, nil
}

// initAccessService 初始化 Access 服务：将首个管理员用户设为超级管理员（super_admin）
func (s *Service) initAccessService(ctx context.Context, adminUserID uuid.UUID) (uuid.UUID, error) {
	if err := s.grpcClients.AccessClient.Client.SuperAdmin.CreateSuperAdmin(ctx, adminUserID.String()); err != nil {
		return uuid.Nil, fmt.Errorf("failed to create super admin: %w", err)
	}
	// super_admin 无 role_id，仅 user_id；metadata 中 admin_role_id 保持为零值
	return uuid.Nil, nil
}

// initAuthService 初始化 Auth 服务
// 创建用户凭证（密码）
func (s *Service) initAuthService(ctx context.Context, userID uuid.UUID, password string) error {
	// 创建用户凭证，首次登录不需要强制修改密码
	if err := s.grpcClients.AuthClient.UserCredential.CreateUserCredential(ctx, userID.String(), password, false); err != nil {
		return fmt.Errorf("failed to create user credential: %w", err)
	}

	return nil
}
