package modelos

type Proceso struct {
	IDProceso     int
	Query         string
	Nombre        string
	FormatoSalida string
	ArchivoModelo string
	CantFechas    int
}
