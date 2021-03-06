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
	"bufio"
	"os"
	"testing"
)

func TestUnmarshalRoute(t *testing.T) {
	var f, err = os.Open("testdata/routes_create.json")
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	var route = "/hsm/test"
	var in = bufio.NewReader(f)
	r, err := unmarshalRoute(route, in)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if route != r.RouteSpec {
		t.Errorf("expected %q, found %q", route, r.RouteSpec)
	}
	if "http://jupyter-hsm:8888" != r.Target {
		t.Errorf("expected %q, found %q", "http://jupyter-hsm:8888", r.Target)
	}
	if !r.Jupyterhub {
		t.Errorf("expected %t, found %t", true, r.Jupyterhub)
	}
	if 2 != len(r.Data) {
		t.Errorf("expected %d, found %d", 2, len(r.Data))
	}
}

func TestEncodeDecodeRoute(t *testing.T) {
	var r = route{}
	r.RouteSpec = "/hsm/test"
	r.Target = "http://jupyter-hsm:8888"
	r.Jupyterhub = true
	var m = make(map[string]interface{})
	m["name"] = "hsm"
	r.Data = m

	s, err := encodeRoute(r)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	var expected = `eyJyb3V0ZV9zcGVjIjoiL2hzbS90ZXN0IiwidGFyZ2V0IjoiaHR0cDovL2p1cHl0ZXItaHNtOjg4ODgiLCJqdXB5dGVyaHViIjp0cnVlLCJkYXRhIjp7Im5hbWUiOiJoc20ifX0=`
	if expected != s {
		t.Errorf("expected %q, found %q", r.Target, s)
	}

	rd, err := decodeRoute(s)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if r.RouteSpec != rd.RouteSpec {
		t.Errorf("expected %q, found %q", r.RouteSpec, rd.RouteSpec)
	}
	if r.Target != rd.Target {
		t.Errorf("expected %q, found %q", r.Target, rd.Target)
	}
	if r.Jupyterhub != rd.Jupyterhub {
		t.Errorf("expected %t, found %t", r.Jupyterhub, rd.Jupyterhub)
	}
	if r.Data["name"] != rd.Data["name"] {
		t.Errorf("expected %q, found %q", r.Data["name"], rd.Data["name"])
	}
}

func TestMarshalRoute(t *testing.T) {
	var r = route{}
	r.RouteSpec = "/hsm/test"
	r.Target = "http://jupyter-hsm:8888"
	r.Jupyterhub = true
	var m = make(map[string]interface{})
	m["name"] = "hsm"
	r.Data = m
	var name, body = marshalRoute(r)
	if r.RouteSpec != name {
		t.Errorf("expected %q, found %q", r.RouteSpec, name)
	}
	if r.Data["name"] != body["name"] {
		t.Errorf("expected %q, found %q", r.Data["name"], body["name"])
	}
	if r.Target != body["target"] {
		t.Errorf("expected %q, found %q", r.Target, body["target"])
	}
	if r.Jupyterhub != body["jupyterhub"] {
		t.Errorf("expected %t, found %t", r.Jupyterhub, body["jupyterhub"])
	}
}
