package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"time"

	systemErr "nfxidentity/errors/src/system"
	bootstrapCommands "nfxidentity/modules/system/application/bootstrap/commands"
	systemStateDomain "nfxidentity/modules/system/domain/system_state"
	"nfxidentity/pkgs/logx"
)

func (s *Service) BootstrapInit(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) error {
	logx.S().Info("Starting system bootstrap initialization...")
	if err := s.checkSystemInitialized(ctx); err != nil {
		return err
	}
	if err := s.checkAllServicesHealth(ctx); err != nil {
		return err
	}
	schemaClearResults, err := s.clearAllSchemas(ctx)
	if err != nil {
		return err
	}
	systemState, err := s.createInitialSystemState(ctx, cmd)
	if err != nil {
		return err
	}
	if err := s.finalizeSystemState(ctx, systemState, schemaClearResults, cmd.Version); err != nil {
		return err
	}
	logx.S().Info("System bootstrap initialization completed")
	return nil
}

func (s *Service) checkSystemInitialized(ctx context.Context) error {
	latestState, err := s.systemStateRepo.Get.Latest(ctx)
	if err != nil {
		if errors.Is(err, systemErr.ErrSystemStateNotFound) {
			return nil
		}
		return fmt.Errorf("failed to get latest system state: %w", err)
	}
	if latestState.Initialized() {
		return fmt.Errorf("system is already initialized")
	}
	return nil
}

func (s *Service) createInitialSystemState(ctx context.Context, cmd bootstrapCommands.BootstrapInitCmd) (*systemStateDomain.SystemState, error) {
	now := time.Now().UTC()
	initialMetadata := map[string]interface{}{
		"bootstrap_started_at": now.Format(time.RFC3339),
		"admin_username":       cmd.AdminUsername,
		"services_initialized": []string{"auth", "asset", "system"},
	}
	systemState, err := systemStateDomain.NewSystemState(systemStateDomain.NewSystemStateParams{
		Initialized:           false,
		InitializedAt:         nil,
		InitializationVersion: nil,
		LastResetAt:           nil,
		LastResetBy:           nil,
		ResetCount:            0,
		Metadata:              initialMetadata,
	})
	if err != nil {
		return nil, err
	}
	if err := s.systemStateRepo.Create.New(ctx, systemState); err != nil {
		return nil, fmt.Errorf("failed to save initial system state: %w", err)
	}
	return systemState, nil
}

func (s *Service) checkAllServicesHealth(ctx context.Context) error {
	logx.S().Info("Checking health of auth, asset, and system")
	maxRetries := 10
	retryInterval := 2 * time.Second
	for attempt := 1; attempt <= maxRetries; attempt++ {
		healthResults, err := s.grpcClients.CheckAllServicesHealth(ctx)
		if err != nil {
			if attempt < maxRetries {
				time.Sleep(retryInterval)
				continue
			}
			return fmt.Errorf("failed to check service health after %d attempts: %w", maxRetries, err)
		}
		allHealthy := true
		unhealthy := []string{}
		for serviceName, healthResp := range healthResults {
			if healthResp == nil || !healthResp.Healthy {
				allHealthy = false
				unhealthy = append(unhealthy, serviceName)
			}
		}
		if allHealthy {
			return nil
		}
		if attempt == maxRetries {
			return fmt.Errorf("unhealthy services: %v", unhealthy)
		}
		time.Sleep(retryInterval)
	}
	return nil
}

func (s *Service) clearAllSchemas(ctx context.Context) (map[string]int, error) {
	clearResults, err := s.grpcClients.ClearAllSchemas(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to clear schemas: %w", err)
	}
	schemaClearResults := make(map[string]int)
	failed := []string{}
	for serviceName, result := range clearResults {
		if result == nil || !result.Success {
			errMsg := "unknown error"
			if result != nil && result.ErrorMessage != nil {
				errMsg = *result.ErrorMessage
			}
			failed = append(failed, fmt.Sprintf("%s: %s", serviceName, errMsg))
			continue
		}
		schemaClearResults[serviceName] = int(result.TablesCleared)
	}
	if len(failed) > 0 {
		return nil, fmt.Errorf("failed to clear schemas: %v", failed)
	}
	return schemaClearResults, nil
}

func (s *Service) finalizeSystemState(
	ctx context.Context,
	systemState *systemStateDomain.SystemState,
	schemaClearResults map[string]int,
	version string,
) error {
	updatedMetadata := map[string]interface{}{
		"bootstrap_started_at":   systemState.Metadata()["bootstrap_started_at"],
		"admin_username":         systemState.Metadata()["admin_username"],
		"services_initialized":   []string{"auth", "asset", "system"},
		"bootstrap_completed_at": time.Now().UTC().Format(time.RFC3339),
		"schema_clear_results":   schemaClearResults,
	}
	if err := systemState.UpdateMetadata(updatedMetadata); err != nil {
		return err
	}
	if err := systemState.Initialize(version); err != nil {
		return err
	}
	return s.systemStateRepo.Update.Generic(ctx, systemState)
}
