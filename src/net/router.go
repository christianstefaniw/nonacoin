package net

import (
	"nonacoin/src/apps/blockchain"

	"github.com/ChristianStefaniw/cgr-v2"
)

func Router() *cgr.Router {
	r := cgr.NewRouter()

	blockchain.GenRoutes(r, "blockchain")
	return r
}
