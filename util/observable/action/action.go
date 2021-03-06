// Copyright 2018 Shift Devices AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action

// Action describes how the subject of an event is altered by the object.
type Action string

const (
	// Replace replaces the current value of the subject with the object.
	Replace Action = "replace"

	// Prepend prepends the object to the list of values of the subject.
	Prepend Action = "prepend"

	// Append appends the object to the list of values of the subject.
	Append Action = "append"

	// Remove removes the object from the list of values of the subject.
	Remove Action = "remove"

	// Reload tells the listener to reload the state of the subject.
	Reload Action = "reload"
)
