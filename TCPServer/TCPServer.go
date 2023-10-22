package tcpserver

import "net"

/*Пакет с TCP сервером*/
func Start() {

	net.Listen("tcp", "127.0.0.1")

}
