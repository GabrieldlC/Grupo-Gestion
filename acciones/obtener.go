package acciones

import (
	"Nueva/modelos"
	"database/sql"
	"strings"
)

func ObtenerProceso(oracleDB *sql.DB) (modelos.Proceso, error) {
	query := "SELECT * FROM PRUEBA.PROCESOS WHERE NOMBRE = 'TODOS'"

	row := oracleDB.QueryRow(query)

	var proceso modelos.Proceso

	err := row.Scan(&proceso.IDProceso, &proceso.Query, &proceso.Nombre, &proceso.FormatoSalida, &proceso.ArchivoModelo, &proceso.CantFechas)
	if err != nil {
		return modelos.Proceso{}, err
	}

	return proceso, nil
}

func ObtenerComprobantes(oracleDB *sql.DB, query string, desde string, hasta string) ([]modelos.Comprobante, error) {
	query = strings.Replace(query, "$1", desde, 1)
	query = strings.Replace(query, "$2", hasta, 1)
	rows, err := oracleDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comprobantes []modelos.Comprobante

	for rows.Next() {
		var comprobante modelos.Comprobante

		err := rows.Scan(&comprobante.Empresa, &comprobante.Factura, &comprobante.Total, &comprobante.Iva, &comprobante.Final, &comprobante.Fecha)
		if err != nil {
			return nil, err
		}

		comprobantes = append(comprobantes, comprobante)
	}

	return comprobantes, nil
}
