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

package rooms

import (
	"errors"
	"github.com/Nordgedanken/Morpheusv2/pkg/matrix"
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/matrix-org/gomatrix"
	"log"
	"strings"
)

// Room holds the needed Room data and allows to work with that. It gets normally loaded from the cache
type Room struct {
	id       string
	aliases  []string
	name     string
	avatar   []byte
	topic    string
	messages []matrix.Message
}

// SetRoomID adds the roomID to the current Room
func (r *Room) SetRoomID(id string) {
	r.id = id
}

// SetRoomAliases adds the aliases to the current Room
func (r *Room) SetRoomAliases(aliases []string) {
	r.aliases = aliases
}

// SetName adds the name to the current Room
func (r *Room) SetName(name string) {
	r.name = name
}

// SetAvatar adds the avatar to the current Room
func (r *Room) SetAvatar(avatar []byte) {
	r.avatar = avatar
}

// SetTopic adds the topic to the current Room
func (r *Room) SetTopic(topic string) {
	r.topic = topic
}

// SetMessages adds the messages to the current Room
func (r *Room) SetMessages(messages []matrix.Message) {
	r.messages = messages
}

// GetRoomID returns the room ID from the current Room
func (r *Room) GetRoomID() string {
	return r.id
}

// GetRoomAliases returns the room aliases from the current Room
func (r *Room) GetRoomAliases() []string {
	return r.aliases
}

// GetName returns the name from the current Room
func (r *Room) GetName() (string, error) {
	if r.name == "" {
		type RespRoomName struct {
			Name string `json:"name"`
		}
		resp := &RespRoomName{}
		err := util.User.GetCli().StateEvent(r.id, "m.room.name", "", resp)
		if err != nil && err.(*gomatrix.HTTPError).WrappedError.(*gomatrix.RespError).Err != "M_NOT_FOUND" {
			return "", err
		}
		if err == nil {
			r.name = resp.Name
		} else if err != nil && err.(*gomatrix.HTTPError).WrappedError.(*gomatrix.RespError).Err == "M_NOT_FOUND" {
			r.name = "Name not found"
		}

	}
	return r.name, nil
}

// GetAvatar returns the avatar from the current Room
func (r *Room) GetAvatar() ([]byte, error) {
	log.Println(r.avatar)
	log.Println(len(r.avatar))
	if len(r.avatar) == 0 {
		log.Println("Avatar getting")
		resp := &gomatrix.Event{}
		err := util.User.GetCli().StateEvent(r.id, "m.room.avatar", "", resp)
		if err != nil {
			return nil, err
		}
		var avatar []byte
		value, exists := resp.Content["url"]
		if !exists {
			return nil, errors.New("missing url in avatar state event")
		}
		url, ok := value.(string)
		if !ok {
			return nil, errors.New("value not ok in avatar state event")
		}
		log.Println(url)
		split := strings.Split(strings.TrimPrefix(url, "mxc://"), "/")
		servername := split[0]
		mediaID := split[1]
		mediaURL := util.User.GetCli().HomeserverURL.String() + "/_matrix/media/r0/thumbnail/" + servername + "/" + mediaID + "?width=61&height=61&method=crop"
		log.Println(mediaURL)
		avatar, err = util.User.GetCli().MakeRequest("GET", mediaURL, nil, nil)
		if err != nil {
			return nil, err
		}
		r.avatar = avatar
		log.Println(avatar)
	}
	return r.avatar, nil
}

// GetTopic returns the topic from the current Room
func (r *Room) GetTopic() (string, error) {
	if r.topic == "" {
		resp := &gomatrix.Event{}
		err := util.User.GetCli().StateEvent(r.id, "m.room.topic", "", resp)
		if err != nil {
			return "", err
		}
		value, exists := resp.Content["topic"]
		if !exists {
			return "", errors.New("missing topic in topic state event")
		}
		topic, ok := value.(string)
		if !ok {
			return "", errors.New("value not ok in topic state event")
		}
		r.topic = topic
	}
	return r.topic, nil
}

// GetMessages returns the messages from the current Room
func (r *Room) GetMessages() []matrix.Message {
	// TODO Lazy load from DB
	return r.messages
}
