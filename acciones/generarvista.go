package acciones

import (
	"Nueva/calendar"
	"Nueva/modelos"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GenerarVista(procesos []modelos.Proceso) modelos.Proceso {
	myApp := app.New()
	myWindow := myApp.NewWindow("Generador de documento")
	myWindow.Resize(fyne.NewSize(400, 400))

	var nombres []string
	for _, proceso := range procesos {
		nombres = append(nombres, proceso.Nombre)
	}

	var selectedProcess modelos.Proceso

	var desde, hasta time.Time

	desdeW := calendar.CalendarWidget(myWindow, "Eija fecha desde", &desde)
	hastaW := calendar.CalendarWidget(myWindow, "Eija fecha hasta", &hasta)

	desdeW.Hide()
	hastaW.Hide()

	procesoSelect := widget.NewSelect(nombres, func(nombre string) {
		desdeW.Show()
		for _, proceso := range procesos {
			if proceso.Nombre == nombre {
				selectedProcess = proceso
				if proceso.CantFechas > 1 {
					hastaW.Show()
				} else {
					hastaW.Hide()
				}
				break
			}
		}
	})
	procesoSelect.PlaceHolder = "Seleccione"

	tituloLista := widget.NewLabel("Proceso:")
	selectContainer := container.NewVBox(
		tituloLista,
		procesoSelect,
	)

	btnSalir := widget.NewButtonWithIcon("Salir", theme.CancelIcon(), func() { myApp.Quit() })

	layoutContainer := container.NewBorder(
		container.NewGridWithRows(2, selectContainer,
			container.NewGridWithColumns(2, desdeW, hastaW)),
		container.NewHBox(layout.NewSpacer(), btnSalir),
		nil,
		nil,
	)

	myWindow.SetContent(layoutContainer)
	myWindow.ShowAndRun()

	return selectedProcess
}
