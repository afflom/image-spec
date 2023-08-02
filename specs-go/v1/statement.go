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

// Statement provides `application/vnd.oci.statement.v1+json` mediatype structure when marshalled to JSON.
type Statement struct {
	specs.Versioned

	// MediaType specifies the type of this document data structure e.g. `application/vnd.oci.statement.v1+json`
	MediaType string `json:"mediaType,omitempty"`

	// Subject is the subject of the statement.
	Subject Subject `json:"subject,omitempty"`

	// Predicate is the predicate of the statement.
	Predicate Predicate `json:"predicate,omitempty"`

	// Object is the object of the statement.
	Object Object `json:"object,omitempty"`
}

// Subject describes a triple's subject
type Subject struct {
	// SubjectType is the location of the Resource that
	// explains client interaction with the Subject.
	//
	// PredicateType is a Docker URL.
	SubjectType string `json:"subjectType,omitempty"`
	// Subject is a Resource used to locate the Subject
	// of the Statement.
	//
	// In practical terms, the Subject is the thing being
	// described by the predicate or the target of the
	// relationship of the Object.
	Subject ResourceAddress `json:"subject,omitempty"`
}

// Predicate describes a triple's predicate
type Predicate struct {
	// PredicateType is the location of the Resource that
	// explains client interaction with the Predicate.
	//
	// PredicateType is a Docker URL.
	PredicateType string `json:"predicateType,omitempty"`
	// Predicate is a Resource used to express the Predicate
	// of the Statement.
	//
	// In practical terms, the Predicate is the description
	// of the Subject or the description of the relationship
	// of the Object to the Subject.
	Predicate ResourceAddress `json:"predicate,omitempty"`
}

// Object describes a triple's object
type Object struct {
	// ObjectType is the location of the Resource that
	// explains client interaction with the Object.
	//
	// ObjectType is a Docker URL.
	ObjectType string `json:"objectType,omitempty"`
	// Object is a Resource used to locate the Object
	// of the Statement.
	//
	// In practical terms, the Object is the thing
	// establishing the relationship to the Subject
	// that is described by the Predicate.
	Object ResourceAddress `json:"object,omitempty"`
}

// Resource is the address of a resource. It is a
// json.RawMessage so that the ResourceType can
// specify its address format and retrieval
// instructions.
type ResourceAddress json.RawMessage
