// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package engine contains code for interacting with container-running backends and handling events from them.
// It supports Docker as the sole task engine type.
package engine

import (
	"github.com/as14692/amazon-ecs-agent/agent/config"
	"github.com/as14692/amazon-ecs-agent/agent/containermetadata"
	"github.com/as14692/amazon-ecs-agent/agent/dockerclient/dockerapi"
	dm "github.com/as14692/amazon-ecs-agent/agent/engine/daemonmanager"
	"github.com/as14692/amazon-ecs-agent/agent/engine/dockerstate"
	"github.com/as14692/amazon-ecs-agent/agent/engine/execcmd"
	"github.com/as14692/amazon-ecs-agent/agent/engine/serviceconnect"
	"github.com/as14692/amazon-ecs-agent/agent/taskresource"
	"github.com/as14692/amazon-ecs-agent/ecs-agent/api/ecs/model/ecs"
	"github.com/as14692/amazon-ecs-agent/ecs-agent/credentials"
	"github.com/as14692/amazon-ecs-agent/ecs-agent/eventstream"
)

// NewTaskEngine returns a default TaskEngine
func NewTaskEngine(cfg *config.Config, client dockerapi.DockerClient,
	credentialsManager credentials.Manager,
	containerChangeEventStream *eventstream.EventStream,
	imageManager ImageManager, hostResources map[string]*ecs.Resource, state dockerstate.TaskEngineState,
	metadataManager containermetadata.Manager,
	resourceFields *taskresource.ResourceFields,
	execCmdMgr execcmd.Manager,
	serviceConnectManager serviceconnect.Manager,
	daemonManagers map[string]dm.DaemonManager) TaskEngine {

	hostResourceManager := NewHostResourceManager(hostResources)
	taskEngine := NewDockerTaskEngine(cfg, client, credentialsManager,
		containerChangeEventStream, imageManager, &hostResourceManager,
		state, metadataManager, resourceFields, execCmdMgr,
		serviceConnectManager, daemonManagers)

	return taskEngine
}
