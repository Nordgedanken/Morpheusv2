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
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

const redBorder = "border: 1px solid red"
const selectMessage = "Select a Server"

// RegisterUI defines the data for the register ui
type RegisterUI struct {
	widget       *widgets.QWidget
	window       *widgets.QMainWindow
	windowWidth  int
	windowHeight int

	localpart            string
	password             string
	confirmpassword      string
	server               string
	passwordInput        *widgets.QLineEdit
	passwordConfirmInput *widgets.QLineEdit
	localpartInput       *widgets.QLineEdit

	serverDropdown *widgets.QComboBox
}

// NewRegisterUI gives you a MainUI struct with profiled data
func NewRegisterUI(windowWidth, windowHeight int, window *widgets.QMainWindow) (registerUI *RegisterUI) {
	registerUI = &RegisterUI{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		window:       window,
	}
	return
}

// GetWidget returns the QWidget of the RegisterUI
func (r *RegisterUI) GetWidget() (widget *widgets.QWidget) {
	return r.widget
}

func (r *RegisterUI) Close() {}

// NewUI prepares the new UI
func (r *RegisterUI) NewUI() error {
	r.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/register.ui")

	file.Open(core.QIODevice__ReadOnly)
	registerWidget := loader.Load(file, r.widget)
	file.Close()

	r.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	registerWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	r.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		registerWidget.Resize(event.Size())
		event.Accept()
	})

	// Run Setup for all fields
	go r.setupDropdown()
	go r.setupUsername()
	go r.setupPassword()
	go r.setupConfirmPassword()
	go r.setupRegisterButton()
	go r.setupLoginButton()

	r.window.SetWindowTitle("Morpheus - Register")

	return nil
}

func (r *RegisterUI) setupUsername() (err error) {
	// UsernameInput
	r.localpartInput = widgets.NewQLineEditFromPointer(r.widget.FindChild("UsernameInput", core.Qt__FindChildrenRecursively).Pointer())

	r.localpartInput.ConnectTextChanged(func(value string) {
		if r.localpartInput.StyleSheet() == redBorder {
			r.localpartInput.SetStyleSheet("")
		}
		r.localpart = value
	})

	r.localpartInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if r.serverDropdown.CurrentText() == selectMessage {
				r.serverDropdown.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == "" {
				r.passwordInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.localpart == "" {
				r.localpartInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.confirmpassword == "" {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == r.confirmpassword {
				r.server = r.serverDropdown.CurrentText()
				RegisterErr := r.register()
				if RegisterErr != nil {
					err = RegisterErr
					return
				}

				r.passwordInput.Clear()
				ev.Accept()
			} else {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
		} else {
			r.localpartInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})
	return
}

func (r *RegisterUI) setupPassword() (err error) {
	// PasswordInput
	r.passwordInput = widgets.NewQLineEditFromPointer(r.widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())

	r.passwordInput.ConnectTextChanged(func(value string) {
		if r.passwordInput.StyleSheet() == redBorder {
			r.passwordInput.SetStyleSheet("")
		}
		r.password = value
	})

	r.passwordInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if r.serverDropdown.CurrentText() == selectMessage {
				r.serverDropdown.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == "" {
				r.passwordInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.localpart == "" {
				r.localpartInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.confirmpassword == "" {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == r.confirmpassword {
				r.server = r.serverDropdown.CurrentText()
				RegisterErr := r.register()
				if RegisterErr != nil {
					err = RegisterErr
					return
				}

				r.passwordInput.Clear()
				ev.Accept()
			} else {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
		} else {
			r.passwordInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})

	return
}

func (r *RegisterUI) setupConfirmPassword() (err error) {
	// PasswordConfirmInput
	r.passwordConfirmInput = widgets.NewQLineEditFromPointer(r.widget.FindChild("PasswordConfirmInput", core.Qt__FindChildrenRecursively).Pointer())

	r.passwordConfirmInput.ConnectTextChanged(func(value string) {
		if r.passwordConfirmInput.StyleSheet() == redBorder {
			r.passwordConfirmInput.SetStyleSheet("")
		}
		r.confirmpassword = value
	})

	r.passwordConfirmInput.ConnectKeyPressEvent(func(ev *gui.QKeyEvent) {
		if int(ev.Key()) == int(core.Qt__Key_Enter) || int(ev.Key()) == int(core.Qt__Key_Return) {
			if r.serverDropdown.CurrentText() == selectMessage {
				r.serverDropdown.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == "" {
				r.passwordInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.localpart == "" {
				r.localpartInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.confirmpassword == "" {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
			if r.password == r.confirmpassword {
				r.server = r.serverDropdown.CurrentText()
				RegisterErr := r.register()
				if RegisterErr != nil {
					err = RegisterErr
					return
				}

				r.passwordInput.Clear()
				ev.Accept()
			} else {
				r.passwordConfirmInput.SetStyleSheet(redBorder)
				ev.Ignore()
			}
		} else {
			r.passwordConfirmInput.KeyPressEventDefault(ev)
			ev.Ignore()
		}
	})

	return
}

func (r *RegisterUI) setupRegisterButton() (err error) {
	// registerButton
	registerButton := widgets.NewQPushButtonFromPointer(r.widget.FindChild("RegisterButton", core.Qt__FindChildrenRecursively).Pointer())

	registerButton.ConnectClicked(func(_ bool) {
		if r.serverDropdown.CurrentText() == selectMessage {
			r.serverDropdown.SetStyleSheet(redBorder)
		}
		if r.password == "" {
			r.passwordInput.SetStyleSheet(redBorder)
		}
		if r.localpart == "" {
			r.localpartInput.SetStyleSheet(redBorder)
		}
		if r.confirmpassword == "" {
			r.passwordConfirmInput.SetStyleSheet(redBorder)
		}
		if r.password == r.confirmpassword {
			r.server = r.serverDropdown.CurrentText()
			RegisterErr := r.register()
			if RegisterErr != nil {
				err = RegisterErr
				return
			}

			r.passwordInput.Clear()
		} else {
			r.passwordConfirmInput.SetStyleSheet(redBorder)
		}
	})
	return
}

func (r *RegisterUI) setupLoginButton() (err error) {
	// loginButton
	loginButton := widgets.NewQPushButtonFromPointer(r.widget.FindChild("loginButton", core.Qt__FindChildrenRecursively).Pointer())
	loginButton.ConnectClicked(func(_ bool) {
		r.Close()
		loginUIs := NewLoginUI(r.windowWidth, r.windowHeight, r.window)
		err := SetNewWindow(loginUIs, r.window, r.windowWidth, r.windowHeight)
		if err != nil {
			log.Panicln(err)
		}
	})
	return
}

func (r *RegisterUI) register() (err error) {

	return
}

func (r *RegisterUI) setupDropdown() (err error) {
	// ServerDropdown
	r.serverDropdown = widgets.NewQComboBoxFromPointer(r.widget.FindChild("ServerChooserDropdown", core.Qt__FindChildrenRecursively).Pointer())

	var helloMatrixRespErr error
	if helloMatrixRespV == nil {
		helloMatrixRespV, helloMatrixRespErr = getHelloMatrixList()
		if helloMatrixRespErr != nil {
			log.Println(helloMatrixRespErr)
			err = helloMatrixRespErr
			return
		}
	}

	hostnames := convertHelloMatrixRespToNameSlice(helloMatrixRespV)
	r.serverDropdown.AddItems(hostnames)
	return
}
