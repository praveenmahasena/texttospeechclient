package dialer

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Dialer struct {
	Port string
	file []byte
}

func New(p string) *Dialer {
	return &Dialer{
		Port: p,
	}
}

func (d *Dialer) Start() error {
	if err := d.getFile(); err != nil {
		return err
	}

	con, conErr := net.Dial("tcp", d.Port)

	if conErr != nil {
		return conErr
	}

	if _, err := con.Write(d.file); err != nil {
		return err
	}

	go func(c net.Conn) {
		r := bufio.NewReader(c)
		r.WriteTo(os.Stdout)
		os.Exit(0)
	}(con)

	select {}

	// return nil
}

func (d *Dialer) getFile() error {
	dir, dirErr := os.Getwd()
	if dirErr != nil {
		return dirErr
	}

	var fileName string = "Ai.m4a"

	fmt.Println("write file name")

	file, fileErr := os.ReadFile(dir + "/" + fileName)

	if fileErr != nil {
		return nil
	}

	d.file = file

	return nil
}
