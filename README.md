# SSE Handler

## Overview

This library is designed to be used by HTTP servers to setup server sent event (SSE) endpoints. The `SSEHandler` object handles setting up the necessary headers, and transmitting the data to the connected client.

### Multiple Clients

The endpoint created by the `SSEHandler` generates a unique id and channel and adds it to the client map. This allows for the data sent to the channels to be broadcast to all the connected clients.

### Custom Headers

The user can set custom headers that the SSE endpoint will set upon client connection. These headers will be specified upon `SSEHandler` initialization by passing in a `map[string]string` holding the header key-value pairs.

## References

1. [Mozilla Server Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events)
1. [W3 Schools](https://www.w3schools.com/html/html5_serversentevents.asp)
1. [Medium Article 1](https://medium.com/@rian.eka.cahya/server-sent-event-sse-with-go-10592d9c2aa1)
1. [The Developer Cafe](https://thedevelopercafe.com/articles/server-sent-events-in-go-595ae2740c7a)
