package net

import "os"

func ServeHTTP() {
	r := Router()
	r.Run(os.Getenv("PORT"))
}
