package dependencyinjector

import (
	"time"

	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/config"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/infra/database"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/infra/web"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/infra/web/handlers"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/infra/web/middlewares"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/pkg/logger"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/pkg/ratelimiter"
	ratelimiter_strategies "github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/pkg/ratelimiter/strategies"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/pkg/responsehandler"
)

type DependencyInjectorInterface interface {
	Inject() (*Dependencies, error)
}

type DependencyInjector struct {
	Config *config.Conf
}

type Dependencies struct {
	Logger                logger.LoggerInterface
	ResponseHandler       responsehandler.WebResponseHandlerInterface
	HelloWebHandler       handlers.HelloWebHandlerInterface
	RateLimiterMiddleware middlewares.RateLimiterMiddlewareInterface
	WebServer             web.WebServerInterface
	RedisDatabase         database.RedisDatabaseInterface
	RateLimiter           ratelimiter.RateLimiterInterface
	RedisLimiterStrategy  ratelimiter_strategies.LimiterStrategyInterface
}

func NewDependencyInjector(c *config.Conf) *DependencyInjector {
	return &DependencyInjector{
		Config: c,
	}
}

func (di *DependencyInjector) Inject() (*Dependencies, error) {
	logger := logger.NewLogger(di.Config.LogLevel)
	responseHandler := responsehandler.NewWebResponseHandler()

	redisDB, err := database.NewRedisDatabase(*di.Config, logger.GetLogger())
	if err != nil {
		return nil, err
	}

	redisLimiterStrategy := ratelimiter_strategies.NewRedisLimiterStrategy(
		redisDB.Client,
		logger.GetLogger(),
		time.Now,
	)

	limiter := ratelimiter.NewRateLimiter(
		logger,
		redisLimiterStrategy,
		di.Config.RateLimiterIPMaxRequests,
		di.Config.RateLimiterTokenMaxRequests,
		di.Config.RateLimiterTimeWindowMilliseconds,
	)

	helloWebHandler := handlers.NewHelloWebHandler(responseHandler)
	rateLimiterMiddleware := middlewares.NewRateLimiterMiddleware(logger, responseHandler, limiter)

	webRouter := web.NewWebRouter(helloWebHandler, rateLimiterMiddleware)
	webServer := web.NewWebServer(
		di.Config.WebServerPort,
		logger.GetLogger(),
		webRouter.Build(),
		webRouter.BuildMiddlewares(),
	)

	return &Dependencies{
		Logger:                logger,
		ResponseHandler:       responseHandler,
		HelloWebHandler:       helloWebHandler,
		RateLimiterMiddleware: rateLimiterMiddleware,
		WebServer:             webServer,
		RedisDatabase:         redisDB,
		RateLimiter:           limiter,
		RedisLimiterStrategy:  redisLimiterStrategy,
	}, nil
}
