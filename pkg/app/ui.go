package app

import (
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/widgets"
)

type ui interface {
	SetCli(cli *gomatrix.Client)
	GetWidget() (widget *widgets.QWidget)
	NewUI() error
}
