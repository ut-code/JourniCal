package curl

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Curl struct {
	pathprefix string
	cookie     bool
	cookiepath string
}

func WithCookie(cookiePath string) Curl {
	return Curl{
		cookie:     true,
		cookiepath: cookiePath,
	}
}

func New() Curl {
	return Curl{
		cookie: false,
	}
}

func (c *Curl) PrefixPath(prefix string) *Curl {
	c.pathprefix = prefix
	return c
}

func (c Curl) String(urlPath string) (string, error) {
	fmt.Println("sending cURL request to: ", c.pathprefix+urlPath)
	var args []string
	if c.cookie {
		args = append(args, "-c", c.cookiepath, "-b", c.cookiepath)
	}
	args = append(args, c.pathprefix+urlPath)
	b, err := exec.Command("curl", args...).Output()
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
