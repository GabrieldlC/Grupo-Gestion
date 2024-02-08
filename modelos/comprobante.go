package modelos

import "time"

type Comprobante struct {
	Empresa string
	Factura string
	Total   int
	Iva     int
	Final   int
	Fecha   time.Time
}
