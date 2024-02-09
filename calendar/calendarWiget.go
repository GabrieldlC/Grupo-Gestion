package calendar

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CalendarWidget(w fyne.Window, txt string, dateSelected *time.Time) *widget.Button {
	var dateButton *widget.Button

	if dateSelected.IsZero() {
		*dateSelected = time.Now()
	}

	calendar := newCalendar(*dateSelected, func(t time.Time) {
		dateButton.Text = fmt.Sprintf("%d %s %d", t.Day(), nombreMes(t.Month().String()), t.Year())
		*dateSelected = t
		dateButton.Refresh()
	})

	dateButton = widget.NewButtonWithIcon(txt, theme.MenuDropDownIcon(), func() {
		position := fyne.CurrentApp().Driver().AbsolutePositionForObject(dateButton)
		position.Y += dateButton.Size().Height
		calendar.showAtPos(w.Canvas(), position)
	})
	dateButton.Alignment = widget.ButtonAlignLeading
	dateButton.IconPlacement = widget.ButtonIconTrailingText
	dateButton.Importance = widget.LowImportance
	return dateButton
}
