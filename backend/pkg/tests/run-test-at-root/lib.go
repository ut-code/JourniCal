package run_test_at_root

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")
	fmt.Println("running test at", dir)
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}
}
