package driver

import (
  "database/sql"
  "log"
  _ "github.com/ibmdb/go_ibm_db"

)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
  con := "HOSTNAME=10.31.45.163;DATABASE=TEPS;PORT=50000;UID=db2inst1;PWD=N3tc00l" //DEV
  db, err := sql.Open("go_ibm_db", con)
  logFatal(err)
  err = db.Ping()
  logFatal(err)
  return db
}
