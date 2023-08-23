/*
 * Copyright 2019 the go-netty project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package frame

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/alopt/go-netty/utils"

	"github.com/alopt/go-netty"
)

func TestFixedLengthCodec(t *testing.T) {

	var cases = []struct {
		length int
		input  []byte
		output interface{}
	}{
		{length: 5, output: []byte("12345")},
		{length: 5, output: strings.NewReader("12345")},
	}

	for index, c := range cases {
		codec := FixedLengthCodec(c.length)
		t.Run(fmt.Sprint(codec.CodecName(), "#", index), func(t *testing.T) {
			ctx := MockHandlerContext{
				MockHandleRead: func(message netty.Message) {
					if dst := utils.MustToBytes(message); !bytes.Equal(dst, c.input) {
						t.Fatal(dst, "!=", c.input)
					}
				},

				MockHandleWrite: func(message netty.Message) {
					c.input = utils.MustToBytes(message)
				},
			}
			codec.HandleWrite(ctx, c.output)
			codec.HandleRead(ctx, c.input)
		})
	}
}
