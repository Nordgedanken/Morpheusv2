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

package messages

import (
	"github.com/matrix-org/gomatrix"
	"time"
)

type Message struct {
}

func (m *Message) SetEventID(id string) {

}

func (m *Message) SetEvent(event *gomatrix.Event) {

}

func (m *Message) SetAuthorMXID(mxid string) {

}

func (m *Message) SetMessage(message string) {

}

func (m *Message) SetTimestamp(ts *time.Time) {

}

func (m *Message) GetEventID() (id string) {
	return ""
}

func (m *Message) GetEvent() (event *gomatrix.Event) {
	return nil
}

func (m *Message) GetAuthorMXID() (mxid string) {
	return ""
}

func (m *Message) GetMessage() (message string) {
	return ""
}

func (m *Message) GetTimestamp() (ts *time.Time) {
	return nil
}

func (m *Message) Show() error {
	return nil
}
