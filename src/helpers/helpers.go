package helpers

import (
	"bytes"
	"nonacoin/src/nonacoin"
)

func RoutePath(parts ...string) string {
	path := new(bytes.Buffer)
	path.Write([]byte(nonacoin.API_PATH))

	for i, part := range parts {
		path.Write([]byte(part))
		if i != len(parts)-1 {
			path.WriteByte('/')
		}

	}

	return path.String()
}
