package structs

import "syscall"

type FdConn struct {
	Fd int
}

func (f FdConn) Read(b []byte) (n int, err error) {
	return syscall.Read(f.Fd, b)
}

func (f FdConn) Write(b []byte) (n int, err error) {
	return syscall.Write(f.Fd, b)
}
