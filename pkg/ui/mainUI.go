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
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/rooms"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix/sync"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// MainUI defines the data for the main ui (that one with the chats)
type MainUI struct {
	widget       *widgets.QWidget
	window       *widgets.QMainWindow
	windowWidth  int
	windowHeight int

	roomList *RoomLayout

	currentRoom matrix.Room
}

// NewMainUI gives you a MainUI struct with prefilled data
func NewMainUI(windowWidth, windowHeight int, window *widgets.QMainWindow) (mainUI *MainUI) {
	mainUI = &MainUI{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		window:       window,
	}
	return
}

// GetWidget returns the QWidget of the MainUI
func (m *MainUI) GetWidget() (widget *widgets.QWidget) {
	return m.widget
}

// NewUI prepares the new UI
func (m *MainUI) NewUI() error {
	m.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2(":/qml/ui/chat.ui")

	file.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(file, m.widget)
	file.Close()

	m.registerSetAvatarEvent()
	m.registerStartSyncEvent()
	m.registerChangeRoomEvent()

	m.widget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
	mainWidget.SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)

	m.widget.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		mainWidget.Resize(event.Size())
		event.Accept()
	})

	// Setup functions and elements
	go m.setupLogout()

	m.window.SetWindowTitle("Morpheus")

	util.E.Raise("setupRoomList", nil)

	return nil
}

func (m *MainUI) registerChangeRoomEvent() {
	roomTitle := widgets.NewQLabelFromPointer(m.widget.FindChild("RoomTitle", core.Qt__FindChildrenRecursively).Pointer())
	roomTopic := widgets.NewQLabelFromPointer(m.widget.FindChild("Topic", core.Qt__FindChildrenRecursively).Pointer())

	util.E.On("changeRoom", func(room interface{}) error {
		switch v := room.(type) {
		case *rooms.Room:
			m.currentRoom = v
			name, err := v.GetName()
			if err != nil {
				return err
			}

			topic, err := v.GetTopic()
			if err != nil {
				return err
			}

			m.window.SetWindowTitle("Morpheus - " + name)

			roomTitle.SetText(name)
			roomTopic.SetText(topic)
		}

		return nil
	})
}

func (m *MainUI) setupRoomList() error {
	m.roomList = NewRoomLayout()
	roomScroll := widgets.NewQScrollAreaFromPointer(m.widget.FindChild("roomScroll", core.Qt__FindChildrenRecursively).Pointer())
	roomScroll.Widget().SetLayout(m.roomList)
	m.roomList.ConnectAddRoom(func(roomID string) {
		err := m.roomList.NewRoom(roomID, roomScroll)
		if err != nil {
			log.Errorln(err)
		}
		if (m.roomList.RoomCount % 5) == 0 {
			util.App.ProcessEvents(core.QEventLoop__AllEvents)
		}
	})
	log.Infoln("Setting up RoomList")
	rooms, err := util.DB.GetRooms()
	if err != nil {
		return err
	}

	m.roomList.Rooms = make(map[string]matrix.Room)

	first := true
	for _, v := range rooms {
		m.roomList.Rooms[v.GetRoomID()] = v
		log.Debugln(v.GetRoomID())
		go m.roomList.AddRoom(v.GetRoomID())
		m.roomList.RoomCount++
		if first {
			util.E.Raise("changeRoom", v)
			util.App.ProcessEvents(core.QEventLoop__AllEvents)
			first = false
		}
	}
	util.App.ProcessEvents(core.QEventLoop__AllEvents)

	m.roomList.RoomCount = 0

	go func() {
		for _, v := range rooms {
			avatar, err := v.GetAvatar()
			if err != nil {
				log.Errorln(err)
			}
			util.E.Raise("setRoomAvatar"+v.GetRoomID(), avatar)
		}
		util.App.ProcessEvents(core.QEventLoop__AllEvents)
	}()

	return nil
}

func (m *MainUI) registerSetAvatarEvent() {
	// userAvatar
	avatarLogo := widgets.NewQLabelFromPointer(m.widget.FindChild("UserAvatar", core.Qt__FindChildrenRecursively).Pointer())
	util.E.On("setAvatar", func(_ interface{}) error {
		image, err := util.User.GetAvatar("")
		if err != nil {
			return err
		}

		avatarLogo.SetPixmap(matrix.ImageToPixmap(image))
		return nil
	})
}

func (m *MainUI) registerStartSyncEvent() {
	util.E.On("startSync", func(_ interface{}) error {
		go sync.NewSync()
		return nil
	})
}

func (m *MainUI) Close() {
	util.E.Remove("setAvatar")
	util.E.Remove("startSync")
	util.E.Remove("setupRoomList")
	util.E.Remove("changeRoom")
	m.roomList.DeleteLater()
	sync.Stop()
}

func (m *MainUI) setupLogout() {
	// Handle LogoutButton
	logoutButton := widgets.NewQPushButtonFromPointer(m.widget.FindChild("LogoutButton", core.Qt__FindChildrenRecursively).Pointer())
	logoutButton.ConnectClicked(func(_ bool) {
		go m.logout()

		m.Close()
		loginUIs := NewLoginUI(m.windowWidth, m.windowHeight, m.window)
		err := SetNewWindow(loginUIs, m.window, m.windowWidth, m.windowHeight)
		if err != nil {
			log.Errorln(err)
		}
	})
	return
}

func (m *MainUI) logout() {
	sync.Stop()
	_, err := util.User.GetCli().Logout()
	if err != nil {
		log.Errorln(err)
	}
	util.User.GetCli().ClearCredentials()

	go func() {
		err = util.DB.RemoveAll()
		if err != nil {
			log.Errorln(err)
		}
	}()
}

func (m *MainUI) Extra() error {
	return m.setupRoomList()
}
