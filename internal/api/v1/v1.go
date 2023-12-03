package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zackarysantana/velocity/internal/api/middleware"
	"github.com/zackarysantana/velocity/internal/db"
)

type V1App struct {
	client db.Connection
}

func CreateV1App(client db.Connection) (*gin.Engine, error) {
	router := gin.Default()

	a := V1App{client: client}

	v1 := router.Group("/v1")

	v1.POST("/first_time_register", a.PostFirstTimeRegister()...)

	authorizedV1 := v1.Group("/", middleware.Auth(client))
	jobs := authorizedV1.Group("/jobs")
	jobs.GET("/", middleware.QueryAmount(100), a.GetJobs)
	jobs.POST("/enqueue", append(middleware.JobsFilter(nil), a.PostJobsEnqueue)...)
	jobs.POST("/dequeue", append(middleware.JobsFilter(nil), a.PostJobsDequeue)...)

	adminV1 := v1.Group("/admin", middleware.AdminAuth(client))
	adminV1.POST("/user", a.PostUser()...)

	return router, nil
}
