package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	var data = []byte(`
01:00.0 Ethernet controller: Intel Corporation Ethernet Controller 10G X550T (rev 01)
01:00.1 Ethernet controller: Intel Corporation Ethernet Controller 10G X550T (rev 01)
01:00.0 Ethernet controller: Intel Corporation Ethernet Controller 10G X550T (rev 01)
d8:00.0 Ethernet controller: Intel Corporation Ethernet Controller 10G X550T (rev 01)
ab:00.0 Ethernet controller: Intel Corporation Ethernet Controller 10G X550T (rev 01)`)
	r := regexp.MustCompile(`^(((\d+)|(\w+)):\d+.\d+)`)
	for _, row := range bytes.Split(data, []byte("\n")) {
		if bytes.Contains(row, []byte("Ethernet Controller")) {
			submatch := r.FindString(string(row))
			fmt.Println(submatch)
		}
	}
}
