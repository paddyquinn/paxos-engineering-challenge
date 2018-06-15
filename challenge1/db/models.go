package db

type Error struct {
  Msg string `json:"err_msg"`
}

type Digest struct {
  Hex string `json:"digest"`
}

type Message struct {
  Text string `json:"message"`
}
