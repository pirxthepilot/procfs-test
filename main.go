package main

import (
	"flag"
	"fmt"

	"github.com/prometheus/procfs"
)

func main() {
	pidPtr := flag.Int("pid", -1, "PID of process")
	//fdPtr := flag.String("fd", "", "A FD number of the process")
	flag.Parse()

	proc, _ := procfs.NewProc(*pidPtr)
	fdinfos, _ := proc.FileDescriptorsInfo()
	fmt.Printf("%+v\n", fdinfos)

	length, _ := proc.InotifyWatchLen()
	fmt.Printf("Inotify watches: %v\n", length)
	// fdinfo, _ := proc.NewFDInfo(*fdPtr)
	// fmt.Printf("%+v\n", fdinfo)
}
