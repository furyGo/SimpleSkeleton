package app

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type ListenType string

const (
	CommandLine ListenType = "COMMAND_LINE"
	TCP         ListenType = "TCP"
)

type Process struct {
	config     map[string]interface{}
	containers []Container
	listenType ListenType
}

// CreateConfiguration set default config
func CreateConfiguration(filePath *string) App {
	reader := YamlReader{}
	configMap := reader.Read(filePath)

	p := &Process{}
	p.overrideConfigs(configMap)

	return p
}

func (p *Process) overrideConfigs(configs map[string]interface{}) {
	p.config = configs
}

func (p *Process) CreateContainer(container func() Container) Container {
	con := container()
	p.containers = append(p.containers, con)
	return con
}

func (p *Process) ListenTCP() App {
	p.listenType = TCP
	return p
}

func (p *Process) ListenCommands() App {
	p.listenType = CommandLine
	return p
}

func (p *Process) Run() {
	switch p.listenType {
	default:
		log.Fatalf("Please choose one listen type.")
		return
	case TCP:
		log.Printf(
			"Now listen on TCP:%d, request data must not be exceeded over 2048 bytes",
			p.config["server-port"])
		tcpHandle(*p)
	case CommandLine:
		commandLineHandle()
	}
}

// TODO
func commandLineHandle() {
	log.Println("Feature Does Not Be Implemented Yet.")
}

func tcpHandle(process Process) {
	ln, err := net.Listen(
		"tcp",
		fmt.Sprintf(":%d", process.config["server-port"]))
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConnection(conn, process)
	}
}

// request exceeded 2048 bytes
func handleConnection(conn net.Conn, process Process) {
	buf := make([]byte, 2048)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}

	x := strings.Split(string(buf), "\n")
	// decode & get route
	route := strings.TrimPrefix(strings.Split(x[0], " ")[1], "/")
	// fan out message for containers
	msg := produce(x[len(x)-1])
	if msg != nil {
		for _, container := range process.containers {
			container.Receive(route, msg)
		}
		log.Println(
			fmt.Sprintf(
				"processing flow:[%s] executed done",
				route))
	}

	er := conn.Close()
	if er != nil {
		return
	}
}

func produce(msg string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- msg
	}()
	return out
}
