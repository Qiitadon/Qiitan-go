package main

import (
	"flag"
	"fmt"
	"runtime/debug"

	"github.com/KEINOS/go-utiles/util"
)

var version string
var versionFlag = flag.Bool("version", false, "show version")

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(getVersion())

		return
	}

	fmt.Println("Hello, Qiitan!")
}

func getVersion() string {
	nameBin := util.GetNameBin()
	verBin := "(unknown)"

	if version != "" {
		verBin = version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		verBin = buildInfo.Main.Version
	}

	return fmt.Sprintf("%s %s", nameBin, verBin)
}
