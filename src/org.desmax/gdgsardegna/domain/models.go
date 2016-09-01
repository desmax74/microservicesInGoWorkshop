// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package domain

type Message struct {
	ID      int    `json:"ID"`
	CONTENT string `json:"CONTENT"`
}


//Interface
type Messenger interface {
	Send(text []byte) error
}

//applyed only to the type Message
func (m *Message) Send(text []byte) error {
	// ...
	return nil
}
