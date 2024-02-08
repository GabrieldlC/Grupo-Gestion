package conexiones

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func ConexionOracle() (*sql.DB, error) {
	connString := fmt.Sprintf("%s/%s@%s:%s/%s", "prueba", "gestion", "66.97.36.229", "1521", "XE")

	origenDB, err := sql.Open("godror", connString)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return nil, err
	}

	err = origenDB.Ping()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return origenDB, nil
}
