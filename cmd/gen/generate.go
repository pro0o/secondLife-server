//go:generate go-jet -source=postgres -host=localhost -port=5432 -user=pgUser -password=pw -dbname=name -schema=public -path=./gen

package main

import "log"

func main() {
	log.Println("Running Go-Jet code generation...")
}
