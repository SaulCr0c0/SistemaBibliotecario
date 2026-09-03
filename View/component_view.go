package view

import "github.com/rivo/tview"

// ComponentView define la interfaz para todos los componentes de vista tview.
type ComponentView interface {
	GetPrimitive() tview.Primitive
}
