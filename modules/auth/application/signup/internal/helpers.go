package internal

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"nfxidentity/enums"
	"nfxidentity/modules/auth/domain/refreshtoken"
	"nfxidentity/pkgs/errx"
	"nfxidentity/pkgs/transaction"

	"github.com/google/uuid"
)

func None() transaction.UoW { return transaction.UoW{} }

func IsMissing(err error) bool {
	e := errx.AsError(err)
	return e != nil && e.Kind == errx.KindNotFound
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func LangOrDefault(lang string) string {
	switch lang {
	case "en", "zh", "fr":
		return lang
	default:
		return "zh"
	}
}

func RandomCode() string {
	b := make([]byte, 3)
	_, _ = rand.Read(b)
	n := (int(b[0])<<16 | int(b[1])<<8 | int(b[2])) % 1000000
	return fmt.Sprintf("%06d", n)
}

func VerifyKey(email string) string {
	return "nfxidentity:verify:" + strings.ToLower(strings.TrimSpace(email))
}

const EmptyProfileID = "00000000-0000-0000-0000-000000000000"

func NewRefreshToken(accountID uuid.UUID, identityID, profileID *uuid.UUID, scope *enums.AuthProfileScope, deviceID *string, hash string, now time.Time) *refreshtoken.RefreshToken {
	return refreshtoken.NewRefreshTokenFromState(refreshtoken.RefreshTokenState{
		ID: uuid.New(), AccountID: accountID, IdentityID: identityID, ProfileID: profileID,
		ProfileScope: scope, DeviceID: deviceID, TokenHash: hash, ExpiresAt: now.Add(30 * 24 * time.Hour), CreatedAt: now,
	})
}

func RoleStrings[T ~string](roles []T) []string {
	out := make([]string, 0, len(roles))
	for _, r := range roles {
		out = append(out, string(r))
	}
	return out
}
