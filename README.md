# What is this?

A demo of NGINX killing websocket connections if pings aren't sent from the application server.

Created for https://github.com/nhooyr/websocket/issues/265

# Getting started

Make sure you've installed Docker, Go, and some websocket client (I use [websocat](https://github.com/vi/websocat)).

Run the Go websocket server on port 8081:

```sh
go run main.go
```

Then, run NGINX on port 8080:
```
docker-compose up
```

The Go websocket server is an echo server. Try connecting directly to the Go server:

```sh
websocat ws://127.0.0.1:8081
foo
foo
```

Neat. This connection should last indefinitely. Now try connecting to the NGINX server:

```sh
websocat ws://127.0.0.1:8080
foo
foo
```

Neat. If you wait 60 seconds, and attempt to send another message, you won't get an echo. If you attempt to send another message, you'll get the following error (NGINX killed the connection):

```
websocat: WebSocketError: I/O failure
websocat: error running
```
