package recover

import (
	"fmt"
	"nfxid/pkgs/logx"
	"nfxid/pkgs/netx/httpresp"
	"os"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RecoverMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Get logger from context or use global logger
				logger := logx.From(c.UserContext())
				if logger == nil {
					logger = logx.L()
				}

				// Log the panic with detailed information
				logger.Error("panic recovered",
					zap.Any("panic", r),
					zap.String("method", c.Method()),
					zap.String("path", c.Path()),
					zap.String("ip", c.IP()),
					zap.String("user_agent", c.Get("User-Agent")),
					zap.String("stack", string(debug.Stack())),
				)
				fmt.Fprintf(os.Stderr, "panic: %v\n%s\n", r, debug.Stack())

				// Return internal server error
				if err := httpresp.Error(c, fiber.StatusInternalServerError, "internal server error"); err != nil {
					logger.Error("failed to send error response", zap.Error(err))
				}
			}
		}()

		return c.Next()
	}
}
