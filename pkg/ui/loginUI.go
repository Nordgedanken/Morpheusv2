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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/users"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/matrix-org/gomatrix"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
	"strings"
)

// LoginUI defines the data for the login ui
type LoginUI struct {
	widget       *widgets.QWidget
	window       *widgets.QMainWindow
	windowWidth  int
	windowHeight int

	localpart      string
	password       string
	server         string
	passwordInput  *widgets.QLineEdit
	localpartInput *widgets.QLineEdit

	serverDropdown *widgets.QComboBox
}

// NewLoginUI gives you a MainUI struct with profiled data
func NewLoginUI(windowWidth, windowHeight int, window *widgets.QMainWindow) (loginUI *LoginUI) {
	loginUI = &LoginUI{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		window:       window,
	}
	return
}

// GetWidget returns the QWidget of the LoginUI
func (l *LoginUI) GetWidget() (widget *widgets.QWidget) {
	return l.widget
}

// NewUI prepares the new UI
func (l *LoginUI) NewUI() error {
	l.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/login.ui")

	file.Open(core.QIODevice__ReadOnly)
	loginWidget := loader.Load(file, l.widget)
	file.Close()

	l.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	loginWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	l.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		loginWidget.Resize(event.Size())
		event.Accept()
	})

	// Run Setup for all fields
	go l.setupDropdown()
	go l.setupLocalpartInput()
	go l.setupPasswordInput()
	go l.setupLoginButton()
	go l.setupRegisterButton()

	l.window.SetWindowTitle("Morpheus - Login")

	return nil
}

func (l *LoginUI) setupLocalpartInput() {
	// LocalpartInput
	l.localpartInput = widgets.NewQLineEditFromPointer(l.widget.FindChild("LocalpartInput", core.Qt__FindChildrenRecursively).Pointer())

	l.localpartInput.ConnectTextChanged(func(value string) {
		if l.localpartInput.StyleSheet() == redBorder {
			l.localpartInput.SetStyleSheet("")
		}
		l.localpart = value
	})
}

func (l *LoginUI) setupPasswordInput() {
	// PasswordInput
	l.passwordInput = widgets.NewQLineEditFromPointer(l.widget.FindChild("PasswordInput", core.Qt__FindChildrenRecursively).Pointer())

	l.passwordInput.ConnectTextChanged(func(value string) {
		if l.passwordInput.StyleSheet() == redBorder {
			l.passwordInput.SetStyleSheet("")
		}
		l.password = value
	})

}

func (l *LoginUI) setupLoginButton() (err error) {
	// loginButton
	loginButton := widgets.NewQPushButtonFromPointer(l.widget.FindChild("LoginButton", core.Qt__FindChildrenRecursively).Pointer())

	loginButton.ConnectClicked(func(_ bool) {
		if l.localpart != "" && l.password != "" && l.server != selectMessage {
			err = l.login()
			if err != nil {
				return
			}
		} else {
			if l.localpart == "" {
				l.localpartInput.SetStyleSheet(redBorder)
			}
			if l.password == "" {
				l.passwordInput.SetStyleSheet(redBorder)
			}
			if l.server == selectMessage {
				l.serverDropdown.SetStyleSheet(redBorder)
			}
		}
	})
	return
}

func (l *LoginUI) setupRegisterButton() (err error) {
	// registerButton
	registerButton := widgets.NewQPushButtonFromPointer(l.widget.FindChild("RegisterButton", core.Qt__FindChildrenRecursively).Pointer())

	registerButton.ConnectClicked(func(_ bool) {
		l.Close()
		registerUIs := NewRegisterUI(l.windowWidth, l.windowHeight, l.window)
		err := SetNewWindow(registerUIs, l.window, l.windowWidth, l.windowHeight)
		if err != nil {
			log.Panicln(err)
		}
	})

	return
}

func (l *LoginUI) login() (err error) {
	mainUIs := NewMainUI(l.windowWidth, l.windowHeight, l.window)
	go l.loginUser(l.localpart, l.password, l.server, mainUIs)

	err = SetNewWindow(mainUIs, l.window, l.windowWidth, l.windowHeight)
	if err != nil {
		return err
	}

	return
}

//getClient returns a Client
func getClient(homeserverURL string) (client *gomatrix.Client, err error) {
	client, ClientErr := gomatrix.NewClient(homeserverURL, "", "")
	if ClientErr != nil {
		err = ClientErr
		return
	}

	return
}

//loginUser Creates a Session for the User
func (l *LoginUI) loginUser(localpart, password, homeserverURL string, mainUIs UI) {
	var cli *gomatrix.Client
	var cliErr error
	log.Debugln(homeserverURL)
	if strings.HasPrefix(homeserverURL, "https://") {
		cli, cliErr = getClient(homeserverURL)
	} else if strings.HasPrefix(homeserverURL, "http://") {
		cli, cliErr = getClient(homeserverURL)
	} else {
		cli, cliErr = getClient("https://" + homeserverURL)
	}
	if cliErr != nil {
		log.Errorln(cliErr)
	}

	localpart = strings.Replace(localpart, "@", "", -1)

	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type:                     "m.login.password",
		User:                     localpart,
		Password:                 password,
		InitialDeviceDisplayName: "Morpheusv2 0.1.0-Alpha",
	})
	if err != nil {
		log.Errorln(err)
	}

	cli.SetCredentials(resp.UserID, resp.AccessToken)

	user := &users.User{}
	user.SetCli(cli)
	user.SetMXID(cli.UserID)

	util.User = user

	go func() {
		err = util.DB.SaveCurrentUser(user)
		if err != nil {
			log.Errorln(err)
		}
	}()

	util.E.Raise("setAvatar", nil)
	util.E.Raise("startSync", nil)
	mainUIs.Extra()
}

func (l *LoginUI) Close() {}

func (l *LoginUI) setupDropdown() (err error) {
	// ServerDropdown
	l.serverDropdown = widgets.NewQComboBoxFromPointer(l.widget.FindChild("ServerChooserDropdown", core.Qt__FindChildrenRecursively).Pointer())

	var helloMatrixRespErr error
	if helloMatrixRespV == nil {
		helloMatrixRespV, helloMatrixRespErr = getHelloMatrixList()
		if helloMatrixRespErr != nil {
			log.Errorln(helloMatrixRespErr)
			err = helloMatrixRespErr
			return
		}
	}

	hostnames := convertHelloMatrixRespToNameSlice(helloMatrixRespV)
	l.serverDropdown.AddItems(hostnames)
	l.serverDropdown.ConnectCurrentTextChanged(func(text string) {
		l.server = text
	})
	l.serverDropdown.ConnectCurrentIndexChanged(func(index int) {
		if l.serverDropdown.StyleSheet() == redBorder {
			l.serverDropdown.SetStyleSheet("")
		}
	})

	return
}

func (l *LoginUI) Extra() error {
	return nil
}
