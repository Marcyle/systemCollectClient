//go build -ldflags="-H windowsgui"
package main

import (
	"systemCollectClient/collect"
)

func main() {
	collect.Run()
}
