package main

import (
	"crypto/rand"
	"flag"
	"path/filepath"

	"log"
	"os"
	"time"
)

var version string = "hatul"

func main() {

	bufferSize := flag.Int("buffer", 1024, "write buffer size ")
	fileName := flag.String("file", "large.file", "file name")
	size := flag.Int("size", 100, "size in MB")

	flag.Parse()

	log.Printf("this is version %#v", version)

	buf := make([]byte, *bufferSize)

	perent := filepath.Dir(*fileName)
	err := os.MkdirAll(perent, os.ModePerm)

	if err != nil {

		log.Printf("error creating dirs for file %#v %s", perent, err.Error())
	}
	file, err := os.Create(*fileName)
	if err != nil {

		log.Printf("fatal error opening file %s", err.Error())
		return
	}
	defer file.Close()

	start := time.Now()

	for i := 0; i < 1024**size; i++ {
		rand.Read(buf)
		file.Write(buf)

	}
	log.Printf("took %s to create %dMB ", time.Since(start), *size)

}
