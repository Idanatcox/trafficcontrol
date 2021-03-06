// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tovalidate

import (
	"fmt"
	"reflect"
)

// GreaterThanZero ...
func GreaterThanZero(value interface{}) error {

	ok := false
	var val *int
	var v int
	typ := reflect.TypeOf(value)
	if typ.Kind() == reflect.Ptr {
		val, ok = value.(*int)
		v = *val
	} else {
		v, ok = value.(int)
	}
	if !ok {
		return fmt.Errorf("cannot parse value '%s'", value)
	}
	if v < 1 {
		return fmt.Errorf("must be greater than zero")
	}
	return nil
}

// ToErrors - Flips to an array of errors
func ToErrors(err map[string]error) []error {
	vErrors := []error{}
	for key, value := range err {
		if value != nil {
			errMsg := fmt.Errorf("'%v' %v", key, value)
			vErrors = append(vErrors, errMsg)
		}
	}
	return vErrors
}
