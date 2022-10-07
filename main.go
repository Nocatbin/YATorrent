package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	// "$HOME/ws/bittorent-go/bencode/decode.go"

	"github.com/golang/glog"
)

func main() {
	torrentPath := flag.String("torrent", "test.torrent", "input torrent file path")
	flag.Parse()
	defer glog.Flush()
	flag.Args()
	glog.Infoln("in main")
	glog.Infoln("input torrent file: ", *torrentPath)

	// str, err := os.ReadFile(*torrentPath)
	// if err != nil {
	// 	return
	// }
	// glog.Infoln(str)

	// "test.torrent"
	file, err := os.Open(*torrentPath)
	if err != nil {
		glog.Fatalln("Open Torrent Err", err)
		return
	}
	reader := bufio.NewReader(file)
	// // ReadNumber Tester
	// val := ReadNumber(reader)
	// glog.Infoln("val: ", val)

	// //DecodeBString Tester
	// str, _ := DecodeBString(reader)
	// glog.Infoln(str)

	// //DecodeBInteger Tester
	// number, _ := DecodeBInteger(reader)
	// glog.Infoln(number)

	list, _ := RecursiveParse(reader)
	fmt.Println(list)
}
