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

package ui

import (
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/widgets"
)

var helloMatrixRespV helloMatrixResp
var CurrentUI UI

type UI interface {
	GetWidget() (widget *widgets.QWidget)
	NewUI() error
	Close()
	Extra() error
}

// SetNewWindow loads the new UI into the QMainWindow
func SetNewWindow(ui UI, window *widgets.QMainWindow, windowWidth, windowHeight int) error {
	log.Debugln("Start changing UI")
	uiErr := ui.NewUI()
	if uiErr != nil {
		return uiErr
	}
	ui.GetWidget().Resize2(windowWidth, windowHeight)
	window.SetCentralWidget(ui.GetWidget())
	log.Debugln("Finished changing UI")
	return nil
}
