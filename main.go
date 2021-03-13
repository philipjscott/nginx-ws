package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"time"
)

func main() {
	s := http.Server{
		Addr:         ":8081",
		Handler:      echoServer{},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	log.Println("listening on " + s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type echoServer struct{}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "nginx ate my homework")

	for {
		err = echo(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			log.Printf("failed to echo with %v: %v\n", r.RemoteAddr, err)
			return
		}
	}
}

func echo(ctx context.Context, c *websocket.Conn) error {
	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}
	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}
	if _, err = io.Copy(w, r); err != nil {
		return err
	}
	return w.Close()
}
