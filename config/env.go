package config

import (
	"flag"
	"fmt"
)

var (
	PORT                                     int
	PGUSER, PGPASSWORD, PGHOST, PGPORT, PGDB string
)

func init() {
	fmt.Println("initialize config")
	flag.IntVar(&PORT, "port", 8080, "environment port")
	flag.StringVar(&PGUSER, "pguser", "", "postgres username")
	flag.StringVar(&PGPASSWORD, "pgpassword", "", "postgres password")
	flag.StringVar(&PGHOST, "pghost", "localhost", "postgres host")
	flag.StringVar(&PGPORT, "pgport", "5432", "postgres port")
	flag.StringVar(&PGDB, "pgdb", "", "postgres database")

	flag.Parse()
}
