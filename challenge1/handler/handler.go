package handler

import (
  "crypto/sha256"
  "database/sql"
  "encoding/json"
  "encoding/hex"
  "io"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/paxos-engineering-challenge/challenge1/db"
)

type Handler struct {
  db *db.SQLite
}

func NewHandler() (*Handler, error) {
  db, err := db.NewSQLite()
  if err != nil {
    return nil, err
  }
  return &Handler{db: db}, nil
}

func (h *Handler) Post(ctx *gin.Context) {
  // Decode the POST data into a
  msg := &db.Message{}
  decoder := json.NewDecoder(ctx.Request.Body)
  err := decoder.Decode(msg)
  if err != nil {
    // Default to an internal server error if the decoding fails.
    statusCode := http.StatusInternalServerError
    errMsg := err.Error()

    // If the decoding fails with an EOF error, no message was passed so return a bad request error code.
    if err == io.EOF {
      statusCode = http.StatusBadRequest
      errMsg = "Missing message data"
    }

    ctx.JSON(statusCode, db.Error{Msg: errMsg})
    return
  }

  // Calculate the digest and encode it as a hex string.
  digest := sha256.Sum256([]byte(msg.Text))
  hexDigest := &db.Digest{Hex: hex.EncodeToString(digest[:])}

  // Insert into the database.
  err = h.db.Insert(hexDigest, msg.Text)
  if err != nil {
    ctx.JSON(http.StatusInternalServerError, db.Error{Msg: err.Error()})
    return
  }

  ctx.JSON(http.StatusOK, hexDigest)
}

func (h *Handler) Get(ctx *gin.Context) {
  // Get the digest from the URL and query the database for it.
  digest := ctx.Param("digest")
  msg, err := h.db.Get(digest)
  if err != nil {
    // Default to an internal server error if the database query fails.
    statusCode := http.StatusInternalServerError
    errMsg := err.Error()

    // If no rows were found for the message return a 404 error.
    if err == sql.ErrNoRows {
      statusCode = http.StatusNotFound
      errMsg = "Message not found"
    }

    ctx.JSON(statusCode, db.Error{Msg: errMsg})
    return
  }

  ctx.JSON(http.StatusOK, msg)
}
