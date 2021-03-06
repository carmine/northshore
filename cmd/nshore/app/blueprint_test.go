// Copyright 2016 The NorthShore Authors All rights reserved.
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

package cmd

import (
	"reflect"
	"testing"
)

func TestParseBlueprint(t *testing.T) {
	var expected Blueprint = Blueprint{
		Version:     "1",
		Type:        "pipeline",
		Name:        "bp_name",
		Provisioner: "docker",
		Stages: map[string]Stage{
			"gerrit": Stage{
				Image:       "openfrontier/gerrit",
				Description: "42",
				Ports: []map[string]string{{"fromPort": "8080", "toPort": "8080"},
					{"fromPort": "29418", "toPort": "29418"}},
				Variables: map[string]string{"NAME": "value", "NAME2": "value2"},
			},
			"jenkins": Stage{
				Image:       "jenkins",
				Description: "",
				Ports: []map[string]string{{"fromPort": "8080", "toPort": "8088"},
					{"fromPort": "50000", "toPort": "50000"}},
				Variables: map[string]string{"NAME": "v"},
			},
			"artifactory": Stage{
				Image:       "jfrog-docker-reg2.bintray.io/jfrog/artifactory-oss:latest",
				Description: "",
				Ports:       []map[string]string{{"fromPort": "8081", "toPort": "8081"}},
			},
		},
	}
	actual, _ := ParseBlueprint("../../../examples")
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Pipelines are not equal!")
	}
}
