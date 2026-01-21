package bootstrap

import (
	"context"
	"fmt"
	"time"

	bootstrapCommands "nfxid/modules/system/application/bootstrap/commands"
	systemStateDomain "nfxid/modules/system/domain/system_state"
	"nfxid/pkgs/logx"

	"github.com/google/uuid"
)

// BootstrapInit 系统初始化
// 流程：
// 1. 创建 system_state 记录（initialized = false），表示初始化开始
// 2. 通过 gRPC 调用其他服务初始化基础数据
// 3. 等待所有服务初始化完成
// 4. 更新 system_state 为 initialized = true
func (s *Service) BootstrapInit(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) error {
	logx.S().Info("🚀 Starting system bootstrap initialization...")

	// 步骤 1: 检查系统是否已经初始化
	latestState, err := s.systemStateAppSvc.GetLatestSystemState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get latest system state: %w", err)
	}

	if latestState.Initialized {
		return fmt.Errorf("system is already initialized")
	}

	// 步骤 2: 创建 system_state 记录（initialized = false），表示初始化开始
	// 使用最小化的参数创建，默认值由 factory 处理
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
		return fmt.Errorf("failed to create system state: %w", err)
	}

	// 保存初始状态记录
	if err := s.systemStateRepo.Create.New(ctx, systemState); err != nil {
		return fmt.Errorf("failed to save initial system state: %w", err)
	}

	logx.S().Info("✅ Created initial system state record (initialized=false)")

	// 步骤 3: 初始化各个服务的基础数据
	initializedServices := []string{}

	// 3.1 初始化 Directory 服务 - 创建第一个系统管理员用户
	logx.S().Info("📦 Initializing Directory service - creating admin user...")
	adminUserID, err := s.initDirectoryService(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to initialize directory service: %w", err)
	}
	initializedServices = append(initializedServices, "directory")
	logx.S().Infof("✅ Directory service initialized - admin user ID: %s", adminUserID)

	// 3.2 初始化 Access 服务 - 创建初始角色和权限
	logx.S().Info("📦 Initializing Access service - creating roles and permissions...")
	adminRoleID, err := s.initAccessService(ctx, adminUserID)
	if err != nil {
		return fmt.Errorf("failed to initialize access service: %w", err)
	}
	initializedServices = append(initializedServices, "access")
	logx.S().Infof("✅ Access service initialized - admin role ID: %s", adminRoleID)

	// 3.3 初始化 Auth 服务 - 创建用户凭证
	logx.S().Info("📦 Initializing Auth service - creating user credentials...")
	if err := s.initAuthService(ctx, adminUserID, cmd.AdminPassword); err != nil {
		return fmt.Errorf("failed to initialize auth service: %w", err)
	}
	initializedServices = append(initializedServices, "auth")
	logx.S().Info("✅ Auth service initialized")

	// 步骤 4: 使用 UpdateMetadata 方法更新 metadata 记录已初始化的服务
	updatedMetadata := map[string]interface{}{
		"bootstrap_started_at":   systemState.Metadata()["bootstrap_started_at"],
		"admin_username":         systemState.Metadata()["admin_username"],
		"services_initialized":   initializedServices,
		"admin_user_id":          adminUserID.String(),
		"admin_role_id":          adminRoleID.String(),
		"bootstrap_completed_at": time.Now().UTC().Format(time.RFC3339),
	}

	// 使用 domain 方法更新 metadata
	if err := systemState.UpdateMetadata(updatedMetadata); err != nil {
		return fmt.Errorf("failed to update system state metadata: %w", err)
	}

	// 保存更新后的 metadata
	if err := s.systemStateRepo.Update.Generic(ctx, systemState); err != nil {
		return fmt.Errorf("failed to save updated system state metadata: %w", err)
	}

	// 步骤 5: 等待所有服务初始化完成（这里可以添加健康检查）
	logx.S().Info("⏳ Waiting for all services to be ready...")
	// TODO: 实现服务健康检查逻辑
	time.Sleep(1 * time.Second) // 临时等待，实际应该检查服务健康状态

	// 步骤 6: 使用 domain entity 的 Initialize 方法更新 system_state 为 initialized = true
	logx.S().Info("✅ All services initialized, marking system as initialized...")
	if err := systemState.Initialize(cmd.Version); err != nil {
		return fmt.Errorf("failed to initialize system state: %w", err)
	}

	// 保存更新后的 system_state
	if err := s.systemStateRepo.Update.Generic(ctx, systemState); err != nil {
		return fmt.Errorf("failed to save initialized system state: %w", err)
	}

	logx.S().Info("🎉 System bootstrap initialization completed successfully!")
	return nil
}

// initDirectoryService 初始化 Directory 服务
// 创建第一个系统管理员用户
func (s *Service) initDirectoryService(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) (uuid.UUID, error) {
	userIDStr, err := s.grpcClients.DirectoryClient.User.CreateUser(ctx, cmd.AdminUsername, "active", true)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create admin user: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID returned: %w", err)
	}

	return userID, nil
}

// initAccessService 初始化 Access 服务
// 创建初始角色和权限
func (s *Service) initAccessService(ctx context.Context, adminUserID uuid.UUID) (uuid.UUID, error) {
	// 1. 创建系统管理员角色
	adminRoleDesc := "系统管理员角色，拥有所有权限"
	adminRoleID, err := s.grpcClients.AccessClient.Role.CreateRole(ctx, "system.admin", "系统管理员", &adminRoleDesc, "global", true)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create admin role: %w", err)
	}

	// 2. 创建基础权限
	permissions := []struct {
		key         string
		name        string
		description string
	}{
		{"system.*", "系统管理", "系统所有权限"},
		{"user.*", "用户管理", "用户所有权限"},
		{"role.*", "角色管理", "角色所有权限"},
		{"permission.*", "权限管理", "权限所有权限"},
	}

	var permissionIDs []string
	for _, perm := range permissions {
		permDesc := perm.description
		permID, err := s.grpcClients.AccessClient.Permission.CreatePermission(ctx, perm.key, perm.name, &permDesc, true)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to create permission %s: %w", perm.key, err)
		}
		permissionIDs = append(permissionIDs, permID)
	}

	// 3. 将权限分配给角色
	for _, permID := range permissionIDs {
		_, err := s.grpcClients.AccessClient.RolePermission.CreateRolePermission(ctx, adminRoleID, permID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to assign permission %s to role: %w", permID, err)
		}
	}

	// 4. 将角色分配给用户（通过 Grant）
	_, err = s.grpcClients.AccessClient.Grant.CreateGrant(ctx, "user", adminUserID.String(), "role", adminRoleID, nil)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to grant role to user: %w", err)
	}

	roleID, err := uuid.Parse(adminRoleID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid role ID returned: %w", err)
	}

	return roleID, nil
}

// initAuthService 初始化 Auth 服务
// 创建用户凭证（密码）
func (s *Service) initAuthService(ctx context.Context, userID uuid.UUID, password string) error {
	// 创建用户凭证，首次登录不需要强制修改密码
	if err := s.grpcClients.AuthClient.UserCredential.CreateUserCredential(ctx, userID.String(), password, nil, false); err != nil {
		return fmt.Errorf("failed to create user credential: %w", err)
	}

	return nil
}
