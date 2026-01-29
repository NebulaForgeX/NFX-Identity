package images

import (
	"context"
	"os"
	"path/filepath"
)

// ClearStorageData removes all files under the storage base path (e.g. data/images).
// Used during system bootstrap to wipe image files when re-initializing.
func (s *Service) ClearStorageData(ctx context.Context) error {
	base := filepath.Clean(s.storageBasePath)
	info, err := os.Stat(base)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // nothing to clear
		}
		return err
	}
	if !info.IsDir() {
		return nil
	}
	entries, err := os.ReadDir(base)
	if err != nil {
		return err
	}
	for _, e := range entries {
		path := filepath.Join(base, e.Name())
		if e.IsDir() {
			if err := os.RemoveAll(path); err != nil {
				return err
			}
		} else {
			if err := os.Remove(path); err != nil {
				return err
			}
		}
	}
	return nil
}
