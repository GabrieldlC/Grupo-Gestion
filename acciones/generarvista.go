package acciones

import (
	"Nueva/modelos"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GenerarVista(procesos []modelos.Proceso) modelos.Proceso {
	myApp := app.New()
	myWindow := myApp.NewWindow("Generador de documento")
	myWindow.Resize(fyne.NewSize(400, 400))

	var nombres []string
	nombres = append(nombres, "(Seleccione)")
	for _, proceso := range procesos {
		nombres = append(nombres, proceso.Nombre)
	}

	var selectedProcess modelos.Proceso

	baseSelect := widget.NewSelect(nombres, func(s string) {
		if s == "(Seleccione)" {
			return
		}
		for _, proceso := range procesos {
			if proceso.Nombre == s {
				fmt.Println("Proceso seleccionado:", s)
				selectedProcess = proceso
				break
			}
		}
	})

	tituloLista := widget.NewLabel("Procesos")

	baseSelect.SetSelected("(Seleccione)")

	selectContainer := container.NewVBox(
		tituloLista,
		baseSelect,
	)

	myWindow.SetContent(selectContainer)
	myWindow.ShowAndRun()

	return selectedProcess
}
