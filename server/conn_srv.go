package server

import (
	"net/http"
	"tapi-controller/controllers"
	"tapi-controller/models"

	"github.com/gin-gonic/gin"
)

type getConSrvReq struct {
	UUID string `uri:"uuid" binding:"required,min=1"`
}

func (server *Server) CreateConnSrvRoute(ctx *gin.Context) {
	// parse input payload
	var connSrv models.TapiCommonContext
	if err := ctx.ShouldBindJSON(&connSrv); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// res, err := controllers.CreateConSrv(ctx, connSrv, server.db)
	res, err := server.tapiCtrl.CreateConSrv(ctx, connSrv, server.db)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, res)
}

func (server *Server) GetAllConnSrvRoute(ctx *gin.Context) {
	res, err := server.tapiCtrl.GetAllConSrv(ctx, server.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) GetConnSrvRoute(ctx *gin.Context) {
	var req getConSrvReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	res, err := server.tapiCtrl.GetConSrv(ctx, server.db, req.UUID)
	if err != nil {
		if err.Error() == controllers.ConnSrvNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) UpdateConSrvRoute(ctx *gin.Context) {
	var req getConSrvReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var connSrv models.TapiConnectivityConnectivityServiceData
	if err := ctx.ShouldBindJSON(&connSrv); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	res, err := server.tapiCtrl.UpdateConSrv(ctx, server.db, req.UUID, connSrv)
	if err != nil {
		if err.Error() == controllers.ConnSrvNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, res)

}
