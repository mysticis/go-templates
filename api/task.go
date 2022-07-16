package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/golang-templates-app/db/sqlc"
)

type createTaskRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateTaskParams{
		Title:          req.Title,
		Content:        req.Content,
		CreatedDate:    time.Now(),
		LastModifiedAt: time.Now(),
	}

	task, err := server.store.CreateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

type getTaskRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTask(ctx *gin.Context) {

	var req getTaskRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	task, err := server.store.GetTask(ctx, req.ID)

	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (server *Server) listTasks(ctx *gin.Context) {

	tasks, err := server.store.ListTasks(ctx)

	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
