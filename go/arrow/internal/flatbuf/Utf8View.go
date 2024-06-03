// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

// / Logically the same as Utf8, but the internal representation uses a view
// / struct that contains the string length and either the string's entire data
// / inline (for small strings) or an inlined prefix, an index of another buffer,
// / and an offset pointing to a slice in that buffer (for non-small strings).
// /
// / Since it uses a variable number of data buffers, each Field with this type
// / must have a corresponding entry in `variadicBufferCounts`.
type Utf8View struct {
	_tab flatbuffers.Table
}

func GetRootAsUtf8View(buf []byte, offset flatbuffers.UOffsetT) *Utf8View {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Utf8View{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Utf8View) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Utf8View) Table() flatbuffers.Table {
	return rcv._tab
}

func Utf8ViewStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func Utf8ViewEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
