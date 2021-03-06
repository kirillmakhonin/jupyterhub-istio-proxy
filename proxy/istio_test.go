/*
Copyright 2020 Splunk Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package proxy

import (
	"testing"
)

func TestVirtualServiceAnnotationNameWithPrefix(t *testing.T) {
	i := IstioClient{vsNamePrefix: "jupyter"}
	actual := i.virtualServiceAnnotationNameWithPrefix()
	expected := "jupyter.splunk.io/proxy-data"
	if expected != actual {
		t.Errorf("expected %q, found %q. Changing this expectation would require manual recycling of exisitng VS in a deployment", expected, actual)
	}
}
