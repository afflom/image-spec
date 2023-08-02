// Copyright 2016-2022 The Linux Foundation
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

package v1

import (
	"encoding/json"

	"github.com/opencontainers/image-spec/specs-go"
)

// Resources include the taxonomy and logic used to interact with
// a resource. Resources include links to "drivers", which a client
// can dynamically load to interact with a resource, to include
// a resource's protocol information needed for resource retrieval.

// Resource provides `application/vnd.oci.resource.v1+json` mediatype structure when marshalled to JSON.
type Resource struct {
	specs.Versioned

	// MediaType specifies the type of this document data structure e.g. `application/vnd.oci.resource.v1+json`
	MediaType string `json:"mediaType,omitempty"`

	// Schema is the resource's type definition
	Schema json.RawMessage `json:"schema,omitempty"`

	// Interface is the resource's interface information.
	Interface ResourceInterface `json:"interface,omitempty"`
}

// ResourceInterface is a Resource interface definition
//
// ResourceInterface should ultimately be defined similar
// to ImageConfig, but is sparsely populated with base
// fields for testing.
type ResourceInterface struct {
	// InterfaceAddress is the address of the interface
	InterfaceAddress string `json:"interface-address,omitempty"`

	// Action is the  function name that is implemented by
	// by the interface.
	//
	// This is probably analogous to an entrypoint, but
	// serves the need for testing.
	Action string `json:"action,omitempty"`
}
