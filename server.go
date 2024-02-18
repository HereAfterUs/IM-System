package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := Server{ip, port}
	return &server
}

// 启动服务器的接口
func (this *Server) Start() {

	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen.Error:", err)
		return
	}

	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept.Error:", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}

}

// 当前链接的业务
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")
}
