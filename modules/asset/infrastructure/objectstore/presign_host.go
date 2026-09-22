package objectstore

import (
	"context"
	"net"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type lanPresignKey struct{}

// WithLAN selects the LAN MinIO presigner (browser Host is an IP).
func WithLAN(ctx context.Context, lan bool) context.Context {
	return context.WithValue(ctx, lanPresignKey{}, lan)
}

func lanFromCtx(ctx context.Context) bool {
	v, _ := ctx.Value(lanPresignKey{}).(bool)
	return v
}

// HostIsLAN is true when the browser reached this API by IPv4/IPv6, not a DNS name.
func HostIsLAN(host string) bool {
	h := strings.TrimSpace(host)
	if h == "" {
		return false
	}
	if name, _, err := net.SplitHostPort(h); err == nil {
		h = name
	}
	h = strings.Trim(h, "[]")
	return net.ParseIP(h) != nil
}

// AttachLANPresign records whether this HTTP request should presign the LAN MinIO URL.
func AttachLANPresign() fiber.Handler {
	return func(c fiber.Ctx) error {
		c.SetContext(WithLAN(c.Context(), HostIsLAN(c.Host())))
		return c.Next()
	}
}
