
package servers_sql

import (
	"database/sql"

)

type Servers struct {}

func (s Servers) RemoveServer(db *sql.DB, server string, delstring string) (int, error) {
	//result, err := db.Exec("delete from kis.icmp where server='"+server+"'")
  result, err := db.Exec(delstring + server+"'")
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rowsDeleted), nil
}
