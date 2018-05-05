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

package users

import (
	"github.com/Nordgedanken/Morpheusv2/pkg/util"
	"github.com/matrix-org/gomatrix"
	"strings"
)

// User holds the needed User data and allows to work with that. It gets normally loaded from the cache
type User struct {
	cli                *gomatrix.Client
	mxid               string
	defaultDisplayName string
	displayName        map[string]string
	defaultAvatar      []byte
	avatar             map[string][]byte
}

// SetCli adds the gomatrix.Client to the current User. This only happens if the User is the current User aka the one that logged into the client.
func (u *User) SetCli(cli *gomatrix.Client) {
	u.cli = cli
}

// SetMXID adds the mxid to the current User
func (u *User) SetMXID(id string) {
	u.mxid = id
}

// SetDisplayName adds the displayName to the current User
func (u *User) SetDisplayName(roomID string, name string) {
	if roomID == "" {
		u.defaultDisplayName = name
	} else {
		if u.displayName == nil {
			u.displayName = make(map[string]string)
		}
		u.displayName[roomID] = name
	}
}

// SetAvatar adds the avatar to the current User
func (u *User) SetAvatar(roomID string, avatar []byte) {
	if roomID == "" {
		u.defaultAvatar = avatar
	} else {
		if u.avatar == nil {
			u.avatar = make(map[string][]byte)
		}
		u.avatar[roomID] = avatar
	}
}

// GetMXID returns the mxid from the current User
func (u *User) GetMXID() string {
	return u.mxid
}

// GetDisplayName returns the displayName from the current User
func (u *User) GetDisplayName(roomID string) (string, error) {
	if roomID == "" {
		if u.defaultDisplayName == "" {
			resp, err := util.User.GetCli().GetDisplayName(u.mxid)
			if err != nil {
				return "", err
			}
			u.defaultDisplayName = resp.DisplayName
		}
		return u.defaultDisplayName, nil
	}
	// TODO get Membership Event from Room instead returning default directly
	if u.displayName[roomID] == "" {
		return u.defaultDisplayName, nil
	}
	return u.displayName[roomID], nil
}

// GetAvatar returns the avatar from the current User
func (u *User) GetAvatar(roomID string) ([]byte, error) {
	if roomID == "" {
		if u.defaultAvatar == nil {
			urlPath := util.User.GetCli().BuildURL("profile", u.mxid, "avatar_url")
			s := struct {
				AvatarURL string `json:"avatar_url"`
			}{}

			_, err := util.User.GetCli().MakeRequest("GET", urlPath, nil, &s)
			if err != nil {
				return nil, err
			}

			split := strings.Split(s.AvatarURL, "/")
			servername := strings.TrimPrefix(split[0], "mxc://")
			mediaID := split[1]
			mediaURL := util.User.GetCli().BuildBaseURL("_matrix/media/r0/download", servername, mediaID)
			avatar, err := util.User.GetCli().MakeRequest("GET", mediaURL, nil, nil)
			if err != nil {
				return nil, err
			}
			u.defaultAvatar = avatar
		}
		return u.defaultAvatar, nil
	}
	// TODO get Membership Event from Room instead returning default directly
	if u.avatar[roomID] == nil {
		return u.defaultAvatar, nil
	}
	return u.avatar[roomID], nil
}

// GetCli returns the gomatrix.Client from the current User
func (u *User) GetCli() *gomatrix.Client {
	return u.cli
}
