package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func flcat(fn string, rlen int) {
	var err error

	var fsz, rsz int64
	var ifp *os.File

	if fn != "" {
		ifp, err = os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		finf, err := ifp.Stat()
		if err != nil {
			log.Fatal(err)
		}
		fsz = finf.Size()
		defer ifp.Close()
	}

	r := make([]byte, rlen)
	for {
		if rsz == fsz {
			return
		}
		_, err := ifp.Read(r)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Println(string(r))
		rsz += int64(rlen)
	}
}

func main() {
	var fn string
	var rlen int
	flag.StringVar(&fn, "fn", "", "-fm name of fl file to emit")
	flag.IntVar(&rlen, "rlen", 0, "-f name of fl file to emit")
	flag.Parse()

	if rlen == 0 {
		log.Fatal("rlen must be supplied")
	}

	flcat(fn, rlen)

}
