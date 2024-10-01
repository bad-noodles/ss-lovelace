package modules

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/bad-noodles/ss-lovelace/pkg/message"
)

type ModuleDescriptor struct {
	Name      string
	Port      int
	Health    bool
	Connected bool
}

type ModuleHandler interface {
	Name() string
	SendChallenge() string
	ValidateChallenge(string) bool
	Message() message.Message
}

type Module struct {
	port      int
	handler   ModuleHandler
	connected bool
	healthy   bool
}

func NewModule(port int, handler ModuleHandler) *Module {
	module := Module{
		port,
		handler,
		false,
		false,
	}
	go module.Listen()

	return &module
}

func (r *Module) Descriptor() ModuleDescriptor {
	return ModuleDescriptor{
		r.handler.Name(),
		r.port,
		r.healthy,
		r.connected,
	}
}

func (r *Module) Listen() {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", r.port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		r.connected = true

		var netErr net.Error
		writer := bufio.NewWriter(conn)
		reader := bufio.NewReader(conn)
	writeLoop:
		for {
			writer.WriteString(r.handler.SendChallenge())
			writer.WriteRune('\n')
			writer.Flush()
			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				answer, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						r.healthy = false
						r.connected = false
						break writeLoop
					}
					if errors.As(err, &netErr) {
						if netErr.Timeout() {
							r.healthy = false
							continue
						}

						r.healthy = false
						r.connected = false
						break writeLoop
					}
					panic(err)
				}

				if r.handler.ValidateChallenge(answer[:len(answer)-1]) {
					r.healthy = true
				}

				break
			}
		}
		conn.Close()
	}
}
