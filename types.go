package ssehandler

type HeaderMap map[string]string
type SSEChannel chan SSEMessage
type SSEChannelMap map[string]SSEChannel
type SSEHandlerOptFunc func(*SSEHandlerOption) error
