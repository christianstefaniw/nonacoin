package blockchain

import (
	"nonacoin/src/helpers"

	"github.com/ChristianStefaniw/cgr-v2"
)

func GenRoutes(r *cgr.Router, subdir string) {
	r.Route(helpers.RoutePath(subdir, "test")).Method("GET").Handler(helloWorld).Insert()
}
