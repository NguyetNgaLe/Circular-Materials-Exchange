package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const rpcTimeout = 10 * time.Second

func rpcContext(c *gin.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.Request.Context(), rpcTimeout)
}

func writeRPCError(c *gin.Context, err error, fallback string) {
	code := status.Code(err)
	httpStatus := http.StatusInternalServerError
	message := fallback
	if err != nil && err.Error() != "" {
		message = err.Error()
	}

	switch code {
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.DeadlineExceeded, codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
	default:
		lower := strings.ToLower(message)
		switch {
		case strings.Contains(lower, "invalid credential"):
			httpStatus = http.StatusUnauthorized
		case strings.Contains(lower, "already exists"), strings.Contains(lower, "duplicate"):
			httpStatus = http.StatusConflict
		case strings.Contains(lower, "not found"), strings.Contains(lower, "no rows"):
			httpStatus = http.StatusNotFound
		case strings.Contains(lower, "not verified"), strings.Contains(lower, "forbidden"):
			httpStatus = http.StatusForbidden
		}
	}

	c.JSON(httpStatus, gin.H{"success": false, "message": message})
}

func pagination(c *gin.Context) (int32, int32) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "50"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 200 {
		pageSize = 50
	}
	return int32(page), int32(pageSize)
}
