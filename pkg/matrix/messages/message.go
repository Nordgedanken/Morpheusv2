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

// Message holds the needed Message data and allows to work with that. It gets normally loaded from the cache
type Message struct {
	eventID   string
	event     *gomatrix.Event
	authorID  string
	text      string
	timestamp *time.Time
}

// SetEventID adds the eventID to the current Message
func (m *Message) SetEventID(id string) {
	m.eventID = id
}

// SetEvent adds the gomatrix Event to the current Message
func (m *Message) SetEvent(event *gomatrix.Event) {
	m.event = event
}

// SetAuthorMXID adds the author MXID to the current Message
func (m *Message) SetAuthorMXID(mxid string) {
	m.authorID = mxid
}

// SetMessage adds the message Text to the current Message
func (m *Message) SetMessage(message string) {
	m.text = message
}

// SetTimestamp adds the Timestamp to the current Message
func (m *Message) SetTimestamp(ts *time.Time) {
	m.timestamp = ts
}

// GetEventID returns the event ID from the current Message
func (m *Message) GetEventID() (id string) {
	return m.eventID
}

// GetEvent returns the gomatrix event from the current Message
func (m *Message) GetEvent() (event *gomatrix.Event) {
	return m.event
}

// GetAuthorMXID returns the authors mxid from the current Message
func (m *Message) GetAuthorMXID() (mxid string) {
	return m.authorID
}

// GetMessage returns the text from the current Message
func (m *Message) GetMessage() (message string) {
	return m.text
}

// GetTimestamp returns the timestamp from the current Message
func (m *Message) GetTimestamp() (ts *time.Time) {
	return m.timestamp
}

// Show adds the current Message to the UI
// TODO implement it
func (m *Message) Show() error {
	return nil
}
