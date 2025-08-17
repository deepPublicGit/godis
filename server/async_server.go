package server

import (
	"log"
	"net"
	"syscall"

	"golang.org/x/sys/windows"
)

func HandleAsync() {

	log.Printf("Listening async on %s:%d", Host, Port)
	events := make([]windows.Handle, MAX_CLIENTS)
	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)

	if err != nil {
		panic(err)
	}

	defer syscall.Close(fd)

	if err = syscall.SetNonblock(fd, true); err != nil {
		panic(err)
	}

	ipv4 := net.ParseIP(Host)

	if err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: Port,
		Addr: [4]byte{ipv4[0], ipv4[1], ipv4[2], ipv4[3]},
	}); err != nil {
		panic(err)
	}

	if err = syscall.Listen(fd, MAX_CLIENTS); err != nil {
		panic(err)
	}

	iocpFd, err := windows.CreateIoCompletionPort(windows.Handle(fd), 0, 0, 0)
	if err != nil {
		panic(err)
	}
	defer windows.Close(iocpFd)

	windows.
		log.Print("[STARTED] PROCESSING CLIENT ", conn.RemoteAddr())

}
