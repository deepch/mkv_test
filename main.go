package main

import (
	"fmt"
	"github.com/deepch/vdk/format/mkv"
	"github.com/deepch/vdk/format/mp4"
	"log"
	"os"
)

func main() {

	file, err := os.Open("i10-2021-10-12-05-dev-nvr2-79fb2f36-1e26-4899-8106-e2319f7dc89c-0.mkv")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fileOut, err := os.OpenFile("i10-2021-10-12-05-dev-nvr2-79fb2f36-1e26-4899-8106-e2319f7dc89c-0.mp4", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	defer fileOut.Close()

	demuxer := mkv.NewDemuxer(file)
	muxer := mp4.NewMuxer(fileOut)

	codec, err := demuxer.Streams()

	if err != nil {
		log.Println(err)
		return
	}

	err = muxer.WriteHeader(codec)

	if err != nil {
		log.Println(err)
		return
	}

	for {

		pkt, err := demuxer.ReadPacket()

		if err != nil {
			log.Println(err)
			break
		}

		err = muxer.WritePacket(pkt)

		if err != nil {
			log.Println(err)
			break
		}

	}
	log.Println("finish")
	muxer.WriteTrailer()

}
