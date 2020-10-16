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

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Parser can be used to read and convert the raw program arguments
type Parser struct {
	flags map[string]Flag
	args  []string
}

// parseOption adds an option to our internal table
func (p *Parser) parseOption(option, next string, prefix, kind string) {
	pieces := strings.Split(option, "=")
	if len(pieces) == 1 {
		if !strings.HasPrefix(next, "-") {
			p.flags[strings.TrimPrefix(pieces[0], prefix)] = Flag{kind, next}
		} else {
			p.flags[strings.TrimPrefix(pieces[0], prefix)] = Flag{kind, ""}
		}
	} else {
		p.flags[strings.TrimPrefix(pieces[0], prefix)] = Flag{kind, pieces[1]}
	}
}

// NewParser does the initial parsing of arguments and returns the resulting Parser
func NewParser(raw []string) (p *Parser, sub string) {
	// Check for subcommand
	if len(raw) < 1 {
		panic("Must use a subcommand")
	}
	sub = raw[0]
	// Init parser
	p = &Parser{make(map[string]Flag), make([]string, 0)}
	// Parse options
	for i, curr := range raw[1:] {
		var next string
		if (i+2) < len(raw) {
			next = raw[i+2]
		}
		switch {
		case strings.HasPrefix(curr, "--"):
			// Parse long option
			p.parseOption(curr, next, "--", Long)
		case strings.HasPrefix(curr, "-"):
			// Parse short option
			p.parseOption(curr, next, "-", Short)
		default:
			// get arguments
			p.args = append(p.args, curr)
		}
	}
	return
}

// setField set a StructField to a value
func setField(field reflect.Value, value string) (extra bool, err error) {
	extra = true
	switch field.Kind() {
	case reflect.Bool:
		// Bools
		field.SetBool(true)
		extra = false
	case reflect.String:
		//String
		field.SetString(value)
	case reflect.Int64:
		// Int64
		i, e := strconv.ParseInt(value, 10, 64)
		if e != nil {
			err = fmt.Errorf("'%s' is not a valid int64", value)
			return
		}
		field.SetInt(i)
	case reflect.Uint64:
		// Uint64
		u, e := strconv.ParseUint(value, 10, 64)
		if e != nil {
			err = fmt.Errorf("'%s' is not a valid uint64", value)
			return
		}
		field.SetUint(u)
	case reflect.Float64:
		// Float64
		f, e := strconv.ParseFloat(value, 64)
		if e != nil {
			err = fmt.Errorf("'%s' is not a valid float64", value)
			return
		}
		field.SetFloat(f)
	default:
		err = fmt.Errorf("[base] Unsupported field type: %s", field.Kind().String())
	}
	return
}

func setSlice(field reflect.Value, value []string) error {
	kind := field.Type().Elem().Kind()
	switch kind {
	case reflect.String:
		field.Set(reflect.ValueOf(value))
	default:
		return fmt.Errorf("[base] Unsupported arg slice type '%s'", kind.String())
	}
	return nil
}

// SetFlags attempts to set the entries in 'flags', using the previously parsed arguments
func (p *Parser) SetFlags(flags interface{}) int {
	// Get the struct element values
	flagsElement := reflect.ValueOf(flags).Elem()
	// Get the struct element types
	flagsType := flagsElement.Type()
	// Iterate over struct fields
	extras := 0
	for i := 0; i < flagsElement.NumField(); i++ {
		tags := flagsType.Field(i).Tag
		element := flagsElement.Field(i)
		if !element.CanSet() {
			continue
		}
		var deletion string
		for k, v := range p.flags {
			if k == tags.Get(v.kind) {
				extra, err := setField(element, v.value)
				if err != nil {
					fmt.Println("Failed to parse flag '" + k + "', reason: " + err.Error())
					os.Exit(1)
				}
				if extra {
					extras++
				}
				deletion = k
				break
			}
		}
		// Remove if a match is found (speed, duplication)
		if deletion != "" {
			delete(p.flags, deletion)
		}
	}
	return extras
}

// UnknownFlags checks for unregistered flags that are set
func (p *Parser) UnknownFlags() {
	if len(p.flags) > 0 {
		for name, flag := range p.flags {
			fmt.Fprintf(os.Stderr, "Unrecognized flag '%s' with argument '%s'\n", name, flag.value)
		}
	}
}

// SetArgs attempts to set the entries in 'args', using the previously parsed arguments
func (p *Parser) SetArgs(args interface{}, extras int) bool {
	argsElement := reflect.ValueOf(args).Elem()
	num := argsElement.NumField()
	if num > 0 {
		if arg := argsElement.Field(num - 1); arg.Kind() == reflect.Slice {
			num--
		}
	}
	p.args = p.args[extras:]
	if len(p.args) < num {
		return false
	}
	for i := 0; i < argsElement.NumField(); i++ {
		arg := argsElement.Field(i)
		if !arg.CanSet() {
			continue
		}
		if arg.Kind() == reflect.Slice {
			if i != (argsElement.NumField() - 1) {
				panic("[base] arg slice must be the last argument")
			}
			if err := setSlice(arg, p.args[i:]); err != nil {
				panic("Failed to parse arg '" + arg.String() + "', reason: " + err.Error())
			}
		} else {
			if _, err := setField(arg, p.args[i]); err != nil {
				panic("Failed to parse arg '" + arg.String() + "', reason: " + err.Error())
			}
		}
	}
	return true
}
