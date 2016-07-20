package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/tabwriter"

	"github.com/kardianos/osext"
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Println("Executable path:")

	osArgs := os.Args[0]
	fmt.Fprintln(w, "os.Args[0]:\t"+osArgs)

	dir, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "filepath.Abs(os.Args[0]):\t"+dir)

	_, fn0, _, _ := runtime.Caller(0)
	fmt.Fprintln(w, "runtime.Caller(0):\t"+fn0)

	filename, err := osext.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "osext.Executable:\t"+filename)
	_ = w.Flush()

	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.TabIndent)

	fmt.Println("")
	fmt.Println("Executable folder:")

	osArgs02 := os.Args[0]
	fmt.Fprintln(w, "filepath.Dir(osos.Args[0]):\t"+filepath.Dir(osArgs02))

	dir2, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "filepath.Abs(filepath.Dir(os.Args[0])):\t"+dir2)

	_, fn02, _, _ := runtime.Caller(0)
	fmt.Fprintln(w, "filepath.Dir(runtime.Caller(0)):\t"+filepath.Dir(fn02))

	filename2, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "osext.ExecutableFolder:\t"+filename2)
	_ = w.Flush()
}
