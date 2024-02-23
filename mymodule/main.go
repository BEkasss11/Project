package main

import (
	"database/sql"
	"fmt"//

	"github.com/lib/pq"
)

func main() {
    const (
        host     = "localhost"
        port     =  5432
        user     = "postgres"
        password = "your-password"
        dbname   = "calhounio_demo"
      )
      db , err := sql.Open("host = localhost " , pq)
}
