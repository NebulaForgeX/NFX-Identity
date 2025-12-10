package grpcx

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"nebulaid/pkgs/logx"
	logmw "nebulaid/pkgs/logx/middleware"
	"nebulaid/pkgs/security/token"
	"nebulaid/pkgs/security/token/servertoken"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DefaultServerOptions(verifier token.Verifier) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()), // TODO: Add OpenTelemetry on client side to inherit parent traceID
		grpc.ChainUnaryInterceptor(
			logmw.UnaryLoggerInjectInterceptor(),
			logging.UnaryServerInterceptor(
				logmw.ZapLoggerAdapter(),
				logging.WithLogOnEvents(logging.FinishCall),
				logging.WithFieldsFromContext(logmw.FieldsFromCtx),
				logging.WithDurationField(func(d time.Duration) logging.Fields {
					return logging.Fields{"duration_ms", float64(d.Microseconds()) / 1000.0}
				}),
				logging.WithLevels(func(code codes.Code) logging.Level {
					switch code {
					case codes.OK:
						return logging.LevelInfo
					case codes.DeadlineExceeded, codes.Unavailable, codes.ResourceExhausted:
						return logging.LevelWarn
					default:
						return logging.LevelError
					}
				}),
			),
			servertoken.UnaryAuthInterceptor(verifier),
			recovery.UnaryServerInterceptor(
				recovery.WithRecoveryHandler(PanicRecoveryHandler),
			),
		),
	}
}

func PanicRecoveryHandler(p any) (err error) {
	logx.L().Error("grpc panic recovered",
		zap.Any("panic", p),
		zap.Stack("stack"),
	)
	fmt.Fprintf(os.Stderr, "panic: %v\n%s\n", p, debug.Stack())

	return status.Errorf(codes.Internal, "internal server error")
}
