// Copyright 2019-present Open Networking Foundation.
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

package setup

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewIssuer(t *testing.T) {
	issuer := NewIssuer().Build()
	assert.Equal(t, "default", issuer.name)
}

func TestIssuer_SetResource(t *testing.T) {
	resource := NewResource().
		SetName("ca-issuer-test").
		SetResourceType("Issuer").
		Build()
	issuer := NewIssuer().SetResource(resource).Build()
	assert.Equal(t, "ca-issuer-test", issuer.GetName())
}
