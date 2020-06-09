package layout

import (
	"github.com/gizak/termui/v3"
	"tui/src/storage/widget"
)

func Draw() {
	termui.Render(widget.MainMenu.Widget)
	termui.Render(widget.Command.Widget)
	termui.Render(widget.Device.Widget)
	termui.Render(widget.Volume.Widget)
	termui.Render(widget.Info.Widget)
}
