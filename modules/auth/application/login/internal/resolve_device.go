package internal

import (
	"strings"

	authErr "nfxidentity/errors/src/auth"
	"nfxidentity/pkgs/ptrx"
)

func ResolveRefreshDeviceID(storedDeviceID *string, clientDeviceID string) (*string, error) {
	clientDeviceID = strings.TrimSpace(clientDeviceID)
	if clientDeviceID == "" {
		return nil, authErr.ErrAccountDeviceIDRequired
	}
	if storedDeviceID == nil || strings.TrimSpace(*storedDeviceID) == "" {
		return ptrx.Ptr(clientDeviceID), nil
	}
	if strings.TrimSpace(*storedDeviceID) != clientDeviceID {
		return nil, authErr.ErrInvalidRefreshToken
	}
	return storedDeviceID, nil
}
