package controllers

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
  "../models"
  "../servers_sql"
  "../utils"
	"strconv"
	"fmt"
)

type Controller struct{}

func (c Controller) RemoveServer(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)

		serversSQL := servers_sql.Servers{}

		rowsDeletedICMP, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.icmp where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}


		rowsDeletedHTTP, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.http where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		rowsDeletedHTTPS, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.https where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		rowsDeletedLDAP, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.ldap where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		rowsDeletedTCPPORT, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.tcpport where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		rowsDeletedFTP, err := serversSQL.RemoveServer(db, params["server"], "delete from kis.ftp where server='")
		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}



    w.Header().Set("Content-Type", "text/plain")

/*
		if rowsDeletedICMP == 0 {
			error.Message = "Not Found in KIS.ICMP"
			utils.SendError(w, http.StatusNotFound, error) //404
			return
		}
    utils.SendSuccess(w, rowsDeletedICMP)
*/

		/*
				{
					“rowsDeletedICMP”: 2,
					“rowsDeletedHTTP”: 0,
					“rowsDeletedHTTPS”: 1,
					“rowsDeletedLDAP”: 8,
					“rowsDeletedTCPPORT”: 0,
					“rowsDeletedFTP”: 0
				}

		*/

/*
		rowsDeletedResponse := `{`+
															`rowsDeletedICMP: `+ strconv.Itoa(rowsDeletedICMP) +`,` +
															`rowsDeletedHTTP: 0,` +
															`rowsDeletedHTTPS: 1,` +
															`rowsDeletedLDAP: 8,` +
															`rowsDeletedTCPPORT: 0,` +
															`rowsDeletedFTP: 0` +
													  `}`
*/

		fmt.Fprint(w,"{"+
									"\"rowsDeletedICMP\":"+ strconv.Itoa(rowsDeletedICMP) +","+
									"\"rowsDeletedHTTP\":"+ strconv.Itoa(rowsDeletedHTTP) +","+
									"\"rowsDeletedHTTPS\":"+ strconv.Itoa(rowsDeletedHTTPS) +","+
									"\"rowsDeletedLDAP\":"+ strconv.Itoa(rowsDeletedLDAP) +","+
									"\"rowsDeletedTCPPORT\":"+ strconv.Itoa(rowsDeletedTCPPORT) +","+
									"\"rowsDeletedFTP\":"+ strconv.Itoa(rowsDeletedFTP) +
									"}")

		//utils.SendSuccess(w, rowsDeletedResponse)








	}

}
