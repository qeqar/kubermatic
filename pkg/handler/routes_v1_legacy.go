/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package handler

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
	"k8c.io/kubermatic/v2/pkg/handler/middleware"
	"k8c.io/kubermatic/v2/pkg/handler/v1/node"
)

// RegisterV1Legacy declares legacy HTTP paths that can be deleted in the future
// At the time of this writing, there is no clear deprecation policy
func (r Routing) RegisterV1Legacy(mux *mux.Router) {
	//
	// Defines a set of HTTP endpoints for nodes that belong to a cluster
	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}").
		Handler(r.getNodeForClusterLegacy())

	mux.Methods(http.MethodPost).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes").
		Handler(r.createNodeForClusterLegacy())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes").
		Handler(r.listNodesForClusterLegacy())

	mux.Methods(http.MethodDelete).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}").
		Handler(r.deleteNodeForClusterLegacy())
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id} project getNodeForClusterLegacy
//
//     Deprecated:
//     Gets a node that is assigned to the given cluster.
//
//     This endpoint is deprecated, please create a Node Deployment instead.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: Node
//       401: empty
//       403: empty
func (r Routing) getNodeForClusterLegacy() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers, r.userProvider),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.SetPrivilegedClusterProvider(r.clusterProviderGetter, r.seedsGetter),
		)(node.GetNodeForClusterLegacyEndpoint(r.projectProvider, r.privilegedProjectProvider, r.userInfoGetter)),
		node.DecodeGetNodeForClusterLegacy,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes project createNodeForClusterLegacy
//
//     Deprecated:
//     Creates a node that will belong to the given cluster
//
//     This endpoint is deprecated, please create a Node Deployment instead.
//     Use POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       201: Node
//       401: empty
//       403: empty
func (r Routing) createNodeForClusterLegacy() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers, r.userProvider),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
		)(node.CreateNodeForClusterLegacyEndpoint()),
		node.DecodeCreateNodeForClusterLegacy,
		SetStatusCreatedHeader(EncodeJSON),
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes project listNodesForClusterLegacy
//
//     Deprecated:
//     Lists nodes that belong to the given cluster
//
//     This endpoint is deprecated, please create a Node Deployment instead.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: []Node
//       401: empty
//       403: empty
func (r Routing) listNodesForClusterLegacy() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers, r.userProvider),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.SetPrivilegedClusterProvider(r.clusterProviderGetter, r.seedsGetter),
		)(node.ListNodesForClusterLegacyEndpoint(r.projectProvider, r.privilegedProjectProvider, r.userInfoGetter)),
		node.DecodeListNodesForClusterLegacy,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id} project deleteNodeForClusterLegacy
//
//    Deprecated:
//    Deletes the given node that belongs to the cluster.
//
//     This endpoint is deprecated, please create a Node Deployment instead.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: empty
//       401: empty
//       403: empty
func (r Routing) deleteNodeForClusterLegacy() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers, r.userProvider),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.SetPrivilegedClusterProvider(r.clusterProviderGetter, r.seedsGetter),
		)(node.DeleteNodeForClusterLegacyEndpoint(r.projectProvider, r.privilegedProjectProvider, r.userInfoGetter)),
		node.DecodeDeleteNodeForClusterLegacy,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}
