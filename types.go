package ssehandler

import "github.com/google/uuid"

type HeaderMap map[string]string
type SSEChannel chan SSEMessage
type SSEChannelMap map[uuid.UUID]SSEChannel
type SSEHandlerOptFunc func(*SSEHandlerOption) error
