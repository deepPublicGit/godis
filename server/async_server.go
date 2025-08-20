package server

import (
	"godis/config"
	"godis/core"
	"godis/core/structs"
	"log"
	"net"
	"syscall"
	"time"
)

func HandleAsync() {

	log.Printf("Listening async on %s:%d with max clients:%d", config.Host, config.Port, config.MAX_CLIENTS)
	ticker := time.NewTicker(time.Duration(config.ExpiryCron))

	events := make([]syscall.EpollEvent, config.MAX_CLIENTS)
	sfd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)

	if err != nil {
		log.Fatal(err)
	}

	defer syscall.Close(sfd)

	if err = syscall.SetNonblock(sfd, true); err != nil {
		log.Fatal(err)
	}

	ipv4 := net.ParseIP(config.Host)

	if err = syscall.Bind(sfd, &syscall.SockaddrInet4{
		Port: config.Port,
		Addr: [4]byte{ipv4[0], ipv4[1], ipv4[2], ipv4[3]},
	}); err != nil {
		log.Fatal(err)
	}

	if err = syscall.Listen(sfd, config.MAX_CLIENTS); err != nil {
		log.Fatal(err)
	}

	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Close(epfd)

	registerEvent(sfd, epfd)

	for {
		select {
		case <-ticker.C:
			structs.DelExpiredKeys(config.ExpirySample)
		}
		numEvents, err := syscall.EpollWait(epfd, events, -1)
		if err != nil {
			log.Fatal(err)
		}
		for i := range numEvents {
			if int(events[i].Fd) == sfd {
				nsfd, _, err := syscall.Accept(sfd)
				if err != nil {
					log.Print(err)
				}

				syscall.SetNonblock(nsfd, true)
				registerEvent(nsfd, epfd)

			} else {
				fdconn := structs.FdConn{Fd: int(events[i].Fd)}
				cmds, err := readClient(fdconn)
				if err != nil {
					log.Print(err)
					syscall.Close(int(events[i].Fd))
					continue
				}
				output, err := core.Eval(cmds)
				if err != nil {
					writeClient(fdconn, err)
					continue
				}
				writeClient(fdconn, output)
			}
		}
	}
}

func registerEvent(sfd int, epfd int) {
	socketEvent := syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(sfd),
	}

	if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, sfd, &socketEvent); err != nil {
		log.Fatal(err)
	}
}
