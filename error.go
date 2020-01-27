/*
 * Copyright 2020 Torben Schinke
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ioutil

import "fmt"

// An IntegerOverflow is always returned, if an Encoder or Decoder recognizes an overflow when performing a conversion.
type IntegerOverflow struct {
	Val interface{} // Val is the current value, which is out of range.
	Max interface{} // Max is the maximum value, which Val should have.
}

// Error reports the current/max message
func (i IntegerOverflow) Error() string {
	return fmt.Sprintf("integer overflow: %d not in [0, %d]", i.Val, i.Max)
}
