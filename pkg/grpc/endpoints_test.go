/*
Copyright 2021 The Dapr Authors
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

//nolint:nosnakecase
package grpc

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/dapr/dapr/pkg/config"
	runtimev1pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
)

func TestEndpointCompleteness(t *testing.T) {
	// Get the list of endpoints in the runtime
	runtimeEndpoints := []string{}
	prefix := "/" + runtimev1pb.Dapr_ServiceDesc.ServiceName + "/"
	for _, m := range runtimev1pb.Dapr_ServiceDesc.Methods {
		runtimeEndpoints = append(runtimeEndpoints, prefix+m.MethodName)
	}
	for _, m := range runtimev1pb.Dapr_ServiceDesc.Streams {
		runtimeEndpoints = append(runtimeEndpoints, prefix+m.StreamName)
	}
	sort.Strings(runtimeEndpoints)

	// Get the list of endpoints in this package (regardless of group)
	packageEndpoints := []string{}
	for _, g := range endpoints {
		packageEndpoints = append(packageEndpoints, g...)
	}
	sort.Strings(packageEndpoints)

	assert.Equal(t, runtimeEndpoints, packageEndpoints, "the list of endpoints defined in this package does not match the endpoints defined in the %s gRPC service", runtimev1pb.Dapr_ServiceDesc.ServiceName)
}

func hUnary(ctx context.Context, req any) (any, error) {
	return nil, nil
}

func hStream(srv any, stream grpc.ServerStream) error {
	return nil
}

func testMiddleware(u grpc.UnaryServerInterceptor, s grpc.StreamServerInterceptor) func(t *testing.T, method string, expectErr bool) {
	return func(t *testing.T, method string, expectErr bool) {
		t.Helper()

		_, err := u(nil, nil, &grpc.UnaryServerInfo{
			FullMethod: method,
		}, hUnary)
		if expectErr {
			assert.Error(t, err)
			assert.ErrorContains(t, err, "Unimplemented")
		} else {
			assert.NoError(t, err)
		}

		err = s(nil, nil, &grpc.StreamServerInfo{
			FullMethod: method,
		}, hStream)
		if expectErr {
			assert.Error(t, err)
			assert.ErrorContains(t, err, "Unimplemented")
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestSetAPIEndpointsMiddleware(t *testing.T) {
	t.Run("state.v1 endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["state.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "state.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("state.v1alpha1 endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1alpha1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["state.v1alpha1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "state.v1alpha1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("publish endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "publish",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["publish.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "publish.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("actors endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "actors",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["actors.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "actors.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("bindings endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "bindings",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["bindings.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "bindings.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("secrets endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "secrets",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["secrets.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "secrets.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("metadata endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "metadata",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["metadata.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "metadata.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("shutdown endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "shutdown",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["shutdown.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "shutdown.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("invoke endpoints allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "invoke",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		for _, e := range endpoints["invoke.v1"] {
			tm(t, e, false)
		}

		for k, v := range endpoints {
			if k != "invoke.v1" {
				for _, e := range v {
					tm(t, e, true)
				}
			}
		}
	})

	t.Run("non-Dapr runtime APIs are always allowed", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "invoke",
				Version:  "v1",
				Protocol: "grpc",
			},
		}

		tm := testMiddleware(setAPIEndpointsMiddlewares(a))

		tm(t, "/myservice/method", false)
	})

	t.Run("no rules, middlewares are nil", func(t *testing.T) {
		u, s := setAPIEndpointsMiddlewares(nil)
		assert.Nil(t, u)
		assert.Nil(t, s)
	})

	t.Run("protocol mismatch, middlewares are nil", func(t *testing.T) {
		a := []config.APIAccessRule{
			{
				Name:     "state",
				Version:  "v1",
				Protocol: "http",
			},
		}

		u, s := setAPIEndpointsMiddlewares(a)
		assert.Nil(t, u)
		assert.Nil(t, s)
	})
}
