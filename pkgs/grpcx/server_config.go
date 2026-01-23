package grpcx

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"nfxid/pkgs/logx"
	logmw "nfxid/pkgs/logx/middleware"
	"nfxid/pkgs/security/token"
	"nfxid/pkgs/security/token/servertoken"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

func DefaultServerOptions(verifier token.Verifier) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()), // TODO: Add OpenTelemetry on client side to inherit parent traceID
		// 配置 keepalive enforcement policy，允许客户端每 5 分钟发送一次 ping
		// MinTime 设置为 1 分钟，确保不会因为 ping 太频繁而拒绝连接
		// 客户端配置是 5 分钟发送一次，所以 1 分钟足够宽松，可以容忍网络延迟和抖动
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             1 * time.Minute, // 允许客户端每 1 分钟发送一次 ping（客户端配置是 5 分钟，所以这个值足够宽松）
			PermitWithoutStream: true,              // 允许无流时发送 ping
		}),
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
