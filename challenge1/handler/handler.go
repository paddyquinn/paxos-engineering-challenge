package handler

import (
  "crypto/sha256"
  "encoding/json"
  "encoding/hex"
  "io"
  "net/http"

  "github.com/gin-gonic/gin"
)

type Handler struct {
  // TODO: change this to a database
  db map[string]string
}

func NewHandler() *Handler {
  return &Handler{db: make(map[string]string)}
}

func (h *Handler) Post(ctx *gin.Context) {
  msg := &model{}
  decoder := json.NewDecoder(ctx.Request.Body)
  err := decoder.Decode(msg)
  if err != nil {
    statusCode := http.StatusInternalServerError
    errMsg := err.Error()
    if err == io.EOF {
      statusCode = http.StatusBadRequest
      errMsg = "Missing message data"
    }
    // TODO: should this be a model?
    ctx.JSON(statusCode, map[string]string{"err_msg": errMsg})
    return
  }
  digest := sha256.Sum256([]byte(msg.Message))
  hexDigest := hex.EncodeToString(digest[:])
  h.db[hexDigest] = msg.Message
  // TODO: should this be a model?
  ctx.JSON(http.StatusOK, map[string]string{"digest": hexDigest})
}

func (h *Handler) Get(ctx *gin.Context) {
  digest := ctx.Param("digest")
  msg, found := h.db[digest]
  if !found {
    // TODO: should this be a model?
    ctx.JSON(http.StatusNotFound, map[string]string{"err_msg": "Message not found"})
    return
  }

  ctx.JSON(http.StatusOK, map[string]string{"message": msg})
}

// TODO: move this to its own package when we use an actual database
type model struct {
  Message string `json:"message"`
}
