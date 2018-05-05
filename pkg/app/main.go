// Copyright © 2018 MTRNord <info@nordgedanken.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/loginUI"
	"github.com/Nordgedanken/Morpheusv2/pkg/mainUI"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/matrix-org/gomatrix"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
)

var args []string
var cli *gomatrix.Client
var windowHeight = 600
var windowWidth = 950
var window *widgets.QMainWindow

// Start prepares the Main QT Window and opens it
func Start(argsArg []string) error {
	args = argsArg
	log.Println("Starting Morpheus v2")

	initApp()

	user, err := util.DB.GetCurrentUser()
	if err != nil {
		return err
	}
	if user == nil {
		loginUIs := loginUI.NewLoginUI(windowWidth, windowHeight, window)
		go SetNewWindow(loginUIs, window)
	} else {
		mainUIs := mainUI.NewMainUI(windowWidth, windowHeight, window)
		go SetNewWindow(mainUIs, window)
	}

	widgets.QApplication_Exec()

	return nil
}

func initApp() {
	log.Println("Create QApp")
	app := widgets.NewQApplication(len(args), args)

	app.SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)
	app.SetApplicationName("Morpheus")
	app.SetApplicationVersion("0.1.0")
	appIcon := gui.NewQIcon5(":/qml/resources/logos/MorpheusBig.png")
	app.SetWindowIcon(appIcon)
	window = widgets.NewQMainWindow(nil, 0)
	app.SetActiveWindow(window)

	desktopApp := app.Desktop()
	primaryScreen := desktopApp.PrimaryScreen()
	screen := desktopApp.Screen(primaryScreen)
	windowX := (screen.Width() - windowHeight) / 2
	windowY := (screen.Height() - windowWidth) / 2

	window.Resize2(windowWidth, windowHeight)
	window.Move2(windowX, windowY)

	window.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		log.Println("Morpheus closed")
	})

}

// SetNewWindow loads the new UI into the QMainWindow
func SetNewWindow(ui ui, window *widgets.QMainWindow) error {
	ui.SetCli(cli)
	uiErr := ui.NewUI()
	if uiErr != nil {
		return uiErr
	}
	ui.GetWidget().Resize2(windowWidth, windowHeight)
	window.SetCentralWidget(ui.GetWidget())
	window.Show()
	return nil
}
