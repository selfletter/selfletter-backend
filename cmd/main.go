package main

import (
	"github.com/gin-gonic/gin"
	"selfletter-backend/api"
	"selfletter-backend/config"
	"selfletter-backend/db"
	"selfletter-backend/helpers"
)

func main() {
	config.ParseConfig() // panics if fails
	cfg := config.GetConfig()

	err := db.Open()
	if err != nil {
		panic("db: failed to connect: " + err.Error())
	}

	if cfg.FirstRun {
		helpers.FirstRun()
	}

	if cfg.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
	
	router := gin.Default()
	rateLimitHandler := helpers.GetRateLimitHandler(cfg.RateLimitingTimeoutSeconds)
	adminRateLimitHandler := helpers.GetRateLimitHandler(cfg.AdminRateLimitingTimeoutSeconds)
	router.GET(cfg.UrlPrefix+"/subscribe", rateLimitHandler, api.Subscribe)
	router.GET(cfg.UrlPrefix+"/unsubscribe", rateLimitHandler, api.Unsubscribe)
	router.GET(cfg.UrlPrefix+"/get_topics", api.GetTopics)
	router.GET(cfg.UrlPrefix+"/get_user_subscribed_topics", api.GetUserSubscribedTopics)
	router.POST(cfg.UrlPrefix+"/admin/add_topic", adminRateLimitHandler, api.AdminAddTopic)
	router.POST(cfg.UrlPrefix+"/admin/add_admin_key", adminRateLimitHandler, api.AdminAddAdminKey)
	router.POST(cfg.UrlPrefix+"/admin/get_users", adminRateLimitHandler, api.AdminGetUsers)
	router.POST(cfg.UrlPrefix+"/admin/remove_topic", adminRateLimitHandler, api.AdminRemoveTopic)
	router.POST(cfg.UrlPrefix+"/admin/send", adminRateLimitHandler, api.AdminSend)

	err = router.Run(cfg.InternalAddress)
	if err != nil {
		panic("api: failed to start: " + err.Error())
	}
}
