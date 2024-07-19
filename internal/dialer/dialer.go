package dialer

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if _, err := con.Write(d.file); err != nil {
			con.Close()
			log.Fatalln(err)
		}
	}(wg)

	wg.Add(1)

	go func(c net.Conn, w *sync.WaitGroup) {
		defer wg.Done()

		data, dataErr := io.ReadAll(c)

		fmt.Println(string(data))

		if dataErr != nil {
			log.Fatalln(dataErr)
		}

	}(con, wg)

	wg.Wait()
	con.Close()

	return nil
}

func (d *Dialer) getFile() error {
	dir, dirErr := os.Getwd()
	if dirErr != nil {
		return dirErr
	}

	var fileName string

	fmt.Fscan(os.Stdin, &fileName)

	file, fileErr := os.ReadFile(dir + "/" + fileName)

	if fileErr != nil {
		return fileErr
	}

	d.file = file

	return nil
}
