package handlers

import (
	"github.com/as14692/amazon-ecs-agent/ecs-agent/api/ecs"
	"github.com/as14692/amazon-ecs-agent/ecs-agent/credentials"
)

type TaskProtectionClientFactoryInterface interface {
	NewTaskProtectionClient(taskRoleCredential credentials.TaskIAMRoleCredentials) ecs.ECSTaskProtectionSDK
}
