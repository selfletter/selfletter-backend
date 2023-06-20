package initialization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfletter-backend/api"
	"selfletter-backend/config"
	"selfletter-backend/ratelimiting"
)

var server *http.Server

func InitializeRouter() error {
	cfg := config.GetConfig()

	if cfg.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	rateLimitHandler := ratelimiting.GetRateLimitHandler(cfg.RateLimitingTimeoutSeconds)
	adminRateLimitHandler := ratelimiting.GetRateLimitHandler(cfg.AdminRateLimitingTimeoutSeconds)
	router.GET(cfg.UrlPrefix+"/subscribe", rateLimitHandler, api.Subscribe)
	router.GET(cfg.UrlPrefix+"/unsubscribe", rateLimitHandler, api.Unsubscribe)
	router.GET(cfg.UrlPrefix+"/edit_subscribed_topics", rateLimitHandler, api.EditSubscribedTopics)
	router.GET(cfg.UrlPrefix+"/get_topics", api.GetTopics)
	router.GET(cfg.UrlPrefix+"/get_user_subscribed_topics", api.GetUserSubscribedTopics)
	router.POST(cfg.UrlPrefix+"/admin/add_topic", adminRateLimitHandler, api.AdminAddTopic)
	router.POST(cfg.UrlPrefix+"/admin/add_admin_key", adminRateLimitHandler, api.AdminAddAdminKey)
	router.POST(cfg.UrlPrefix+"/admin/get_users", adminRateLimitHandler, api.AdminGetUsers)
	router.POST(cfg.UrlPrefix+"/admin/remove_topic", adminRateLimitHandler, api.AdminRemoveTopic)
	router.POST(cfg.UrlPrefix+"/admin/send", adminRateLimitHandler, api.AdminSend)
	router.POST(cfg.UrlPrefix+"/admin/get_server_config", adminRateLimitHandler, api.AdminGetServerConfig)
	router.POST(cfg.UrlPrefix+"/admin/write_server_config", adminRateLimitHandler, api.AdminWriteServerConfig)

	server = &http.Server{
		Addr:    cfg.InternalAddress,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
