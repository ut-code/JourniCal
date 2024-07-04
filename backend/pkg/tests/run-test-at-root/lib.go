package run_test_at_root

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
)

/*
go tests run at the directory where the test is located at rather than working directory.
this package allow to run tests at /backend instead of /backend/wherever/the/test/file/is/at,
s.t. it can actually read files such as credentials.json without copying them all around.
ref:
https://intellij-support.jetbrains.com/hc/en-us/community/posts/360009685279-Go-test-working-directory-keeps-changing-to-dir-of-the-test-file-instead-of-value-in-template
*/

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")
	fmt.Println("running test at", dir)
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}
}
