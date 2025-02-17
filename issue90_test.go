/*
 * Copyright 2021 ByteDance Inc.
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
 
package sonic

import (
    "encoding/json"
    "fmt"
    "testing"
)

func TestUnmarshalInfinity(t *testing.T) {
    var v interface{}
    data := []byte("9e370")
    sonicerr := Unmarshal(data, &v)
    stderr := json.Unmarshal(data, &v)
    if sonicerr == nil && stderr != nil {
        t.Errorf("should have unmarshal error like %#v\n", stderr)
    }
    fmt.Println(sonicerr, stderr)
}
