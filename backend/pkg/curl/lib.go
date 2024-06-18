package curl

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Curl struct {
	pathprefix string
	cookiepath string
}

func Init(cookiePath string) Curl {
	return Curl{
		cookiepath: cookiePath,
	}
}

func (c *Curl) PrefixPath(prefix string) *Curl {
	c.pathprefix = prefix
	return c
}

func (c Curl) String(urlPath string) (string, error) {
	fmt.Println("sending cURL request to: ", c.pathprefix+urlPath)
	b, err := exec.Command("curl", "-c", c.cookiepath, "-b", c.cookiepath, c.pathprefix+urlPath).Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// DO NOT PASS a AS VALUE. PASS AS A REFERENCE.
func (c Curl) JSON(urlPath string, a any) error {
	s, err := c.String(urlPath)
	if err != nil {
		return err
	}
	return json.NewDecoder(strings.NewReader(s)).Decode(a)
}
