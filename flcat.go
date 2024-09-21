package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// process a fixed length string record file
// func flcat(fn string, rlen, koff, klen int) {
// fn name of file to process
// rlen record length
// koff offset of key in rexord
// klen record key length
func flcat(fn string, rlen, koff, klen int) {
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
	} else {
		// ifp = os.Stdin
		ifp, err = os.Open("/dev/stdin")
		if err != nil {
			log.Fatal(err)
		}
	}

	r := make([]byte, rlen)
	for {
		if fsz != 0 && rsz == fsz {
			return
		}
		_, err := ifp.Read(r)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		if koff == 0 && klen == 0 {
			fmt.Println(string(r))
		} else {
			fmt.Println(string(r[koff : koff+klen]))
		}
		rsz += int64(rlen)
	}
}

func main() {
	var fn string
	var rlen, koff, klen int
	flag.StringVar(&fn, "fn", "", "name of fl file to emit")
	flag.IntVar(&rlen, "rlen", 0, "record length")
	flag.IntVar(&koff, "koff", 0, "offset of key in record")
	flag.IntVar(&rlen, "rlen", 0, "record key length")
	flag.Parse()

	if rlen == 0 {
		log.Fatal("rlen must be supplied")
	}
	if koff != 0 || klen != 0 {
		if koff+klen > rlen {
			log.Fatal("key must fall within record bounds")
		}
	}

	flcat(fn, rlen, koff, klen)

}
