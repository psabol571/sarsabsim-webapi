package api

import (
    _ "embed"
    "net/http"

    "github.com/gin-gonic/gin"
)

//go:embed hospital-mgmt.openapi.yaml
var openapiSpec []byte

func HandleOpenApi(ctx *gin.Context) {
    ctx.Data(http.StatusOK, "application/yaml", openapiSpec)
}