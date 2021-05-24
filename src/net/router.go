package net

import (
	blockchainapi "nonacoin/src/apps/blockchain"

	"github.com/ChristianStefaniw/cgr-v2"
)

func Router() *cgr.Router {
	r := cgr.NewRouter()

	blockchainapi.GenRoutes(r, "blockchain")
	return r
}
