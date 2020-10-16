//
// Copyright 2017-2020 Bryan T. Meyers <root@datadrake.com>
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
//

package options

const (
	// Short flags start with a single hyphen (ie. -flag)
	Short = "short"
	// Long flags start with two hyphens (ie. --flag)
	Long = "long"
)

// Flag is an CLI switch that may affect execution
type Flag struct {
	kind  string
	value string
}
