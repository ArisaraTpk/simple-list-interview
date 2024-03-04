package https

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"simple-list-interview/infra"
	"simple-list-interview/internal/adapters/handler"
	"simple-list-interview/internal/adapters/repository"
	"simple-list-interview/internal/core/services"
	"simple-list-interview/middleware/jwt"
)

var (
	validate = validator.New()
)

func InitRoutes() {
	r := gin.New()

	bindLogin(r)

	apisGroup := r.Group("/apis").Use(jwt.CheckJWT)

	bindGetInterviewList(&apisGroup)
	bindGetInterviewDetail(&apisGroup)
	bindUpdateInterviewDetail(&apisGroup)
	bindInterviewHistoryList(&apisGroup)
	bindEditComment(&apisGroup)
	bindCommentList(&apisGroup)

	appPort := viper.GetString("service.port")
	err := r.Run(":" + appPort)
	if err != nil {
		log.Error().Err(err).Msg("Start run error")
	}
}

func bindLogin(r *gin.Engine) {
	interviewDb := repository.NewUserRepo(infra.InterviewDB)
	loginSvc := services.NewLoginSvc(validate, interviewDb)
	loginHdl := handler.NewLoginHdl(loginSvc)

	r.POST("/login", loginHdl.Login)
}

func bindGetInterviewList(r *gin.IRoutes) {
	interviewDb := repository.NewInterviewsRepo(infra.InterviewDB)
	interviewSvc := services.NewInterviewListSvc(validate, interviewDb)
	interviewHdl := handler.NewInterviewListHdl(interviewSvc)

	(*r).GET("/interviews", interviewHdl.InterviewList)
}

func bindGetInterviewDetail(r *gin.IRoutes) {
	interviewDb := repository.NewInterviewsRepo(infra.InterviewDB)
	userDb := repository.NewUserRepo(infra.InterviewDB)
	interviewSvc := services.NewGetInterviewDetailSvc(validate, interviewDb, userDb)
	interviewHdl := handler.NewGetInterviewDetailHdl(interviewSvc)

	(*r).GET("/interviews/:interviewId/detail", interviewHdl.GetInterviewDetail)
}

func bindUpdateInterviewDetail(r *gin.IRoutes) {
	interviewDb := repository.NewInterviewsRepo(infra.InterviewDB)
	historyDb := repository.NewInterviewHistoryRepo(infra.InterviewDB)
	interviewSvc := services.NewUpdateInterviewDetailSvc(validate, interviewDb, historyDb)
	interviewHdl := handler.NewUpdateInterviewDetailHdl(interviewSvc)

	(*r).PATCH("/interviews/:interviewId/detail", interviewHdl.UpdateInterviewDetail)
}

func bindInterviewHistoryList(r *gin.IRoutes) {
	interviewDb := repository.NewInterviewHistoryRepo(infra.InterviewDB)
	interviewSvc := services.NewInterviewHistoryListSvc(validate, interviewDb)
	interviewHdl := handler.NewInterviewHistoryListHdl(interviewSvc)

	(*r).GET("/interviews/:interviewId/history", interviewHdl.InterviewHistoryList)
}

func bindEditComment(r *gin.IRoutes) {
	commentDb := repository.NewCommentRepo(infra.InterviewDB)
	editCommentSvc := services.NewEditCommentSvc(validate, commentDb)
	editCommentHdl := handler.NewEditCommentHdl(editCommentSvc)

	(*r).PUT("/interviews/:interviewId/comments/:commentId", editCommentHdl.EditComment)
}

func bindCommentList(r *gin.IRoutes) {
	commentDb := repository.NewCommentRepo(infra.InterviewDB)
	commentListSvc := services.NewCommentListSvc(validate, commentDb)
	commentListHdl := handler.NewCommentListHdl(commentListSvc)

	(*r).GET("/interviews/:interviewId/comments", commentListHdl.CommentList)
}
