package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const DEFAULT_HOST string = "localhost"
const DEFAULT_PORT int64 = 8080

const CMD_DELIMITER string = "="

func main() {
	paramLength := len(os.Args)

	host := DEFAULT_HOST
	port := DEFAULT_PORT

	if paramLength > 1 {
		mappings := getMappings(os.Args[1:])

		if mappings["host"] == "" {
			host = DEFAULT_HOST
		} else {
			host = mappings["host"]
		}

		if mappings["port"] == "" {
			port = DEFAULT_PORT
		} else {
			port, _ = strconv.ParseInt(mappings["port"], 0, 0)
		}
	}

	fmt.Println("Preparing server on " + host + ":" + strconv.FormatInt(port, 10))

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.ListenAndServe(host + ":" + strconv.FormatInt(port, 10), nil)
}

/**
 *	Returns mappings key -> value from arguments
 */
func getMappings(args []string) map[string]string {
	mappings := map[string]string{}
	argsLength := len(args)

	for i := 0; i < argsLength; i++ {
		items := strings.Split(args[i], CMD_DELIMITER)

		// no pair
		if len(items) < 2 {
			continue
		}

		mappings[items[0]] = items[1]
	}

	return mappings
}