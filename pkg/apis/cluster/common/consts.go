/*
Copyright 2018 The Kubernetes Authors.

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

package common

// Constants aren't automatically generated for unversioned packages.
// Instead share the same constant for all versioned packages
type MachineStatusError string

const (
	// Represents that the combination of configuration in the MachineSpec
	// is not supported by this cluster. This is not a transient error, but
	// indicates a state that must be fixed before progress can be made.
	//
	// Example: the ProviderConfig specifies an instance type that doesn't exist,
	InvalidConfigurationMachineError MachineStatusError = "InvalidConfiguration"

	// This indicates that the MachineSpec has been updated in a way that
	// is not supported for reconciliation on this cluster. The spec may be
	// completely valid from a configuration standpoint, but the controller
	// does not support changing the real world state to match the new
	// spec.
	//
	// Example: the responsible controller is not capable of changing the
	// container runtime from docker to rkt.
	UnsupportedChangeMachineError MachineStatusError = "UnsupportedChange"

	// This generally refers to exceeding one's quota in a cloud provider,
	// or running out of physical machines in an on-premise environment.
	InsufficientResourcesMachineError MachineStatusError = "InsufficientResources"

	// There was an error while trying to create a Node to match this
	// Machine. This may indicate a transient problem that will be fixed
	// automatically with time, such as a service outage, or a terminal
	// error during creation that doesn't match a more specific
	// MachineStatusError value.
	//
	// Example: timeout trying to connect to GCE.
	CreateMachineError MachineStatusError = "CreateError"

	// An error was encountered while trying to delete the Node that this
	// Machine represents. This could be a transient or terminal error, but
	// will only be observable if the provider's Machine controller has
	// added a finalizer to the object to more gracefully handle deletions.
	//
	// Example: cannot resolve EC2 IP address.
	DeleteMachineError MachineStatusError = "DeleteError"
)

type ClusterStatusError string

const (
	// InvalidConfigurationClusterError indicates that the cluster
	// configuration is invalid.
	InvalidConfigurationClusterError ClusterStatusError = "InvalidConfiguration"

	// UnsupportedChangeClusterError indicates that the cluster
	// spec has been updated in an unsupported way. That cannot be
	// reconciled.
	UnsupportedChangeClusterError ClusterStatusError = "UnsupportedChange"

	// CreateClusterError indicates that an error was encountered
	// when trying to create the cluster.
	CreateClusterError ClusterStatusError = "CreateError"

	// UpdateClusterError indicates that an error was encountered
	// when trying to update the cluster.
	UpdateClusterError ClusterStatusError = "UpdateError"

	// DeleteClusterError indicates that an error was encountered
	// when trying to delete the cluster.
	DeleteClusterError ClusterStatusError = "DeleteError"
)

type MachineSetStatusError string

const (
	// Represents that the combination of configuration in the MachineTemplateSpec
	// is not supported by this cluster. This is not a transient error, but
	// indicates a state that must be fixed before progress can be made.
	//
	// Example: the ProviderConfig specifies an instance type that doesn't exist.
	InvalidConfigurationMachineSetError MachineSetStatusError = "InvalidConfiguration"
)

type MachineDeploymentStrategyType string

const (
	// Kill all existing machines before creating new ones.
	RecreateMachineDeploymentStrategyType MachineDeploymentStrategyType = "Recreate"

	// Replace the old MachineSet by new one using rolling update
	// i.e gradually scale down the old MachineSet and scale up the new one.
	RollingUpdateMachineDeploymentStrategyType MachineDeploymentStrategyType = "RollingUpdate"
)

type MachineDeploymentConditionType string

// These are valid conditions of a MachineDeployment.
const (
	// Available means the MachineDeployment is available, ie. at least the minimum available
	// replicas required are up and running for at least minReadySeconds.
	MachineDeploymentAvailable MachineDeploymentConditionType = "Available"

	// Progressing means the MachineDeployment is progressing. Progress for a MachineDeployment is
	// considered when a new machine set is created or adopted, and when new machines scale
	// up or old machines scale down. Progress is not estimated for paused MachineDeployments or
	// when progressDeadlineSeconds is not specified.
	MachineDeploymentProgressing MachineDeploymentConditionType = "Progressing"

	// ReplicaFailure is added in a MachineDeployment when one of its machines fails to be created
	// or deleted.
	MachineDeploymentReplicaFailure MachineDeploymentConditionType = "ReplicaFailure"

	// MachineDeploymentFrozen is added in a MachineDeployment when user dont want any actions to be taken for this Machinedeployment
	MachineDeploymentFrozen MachineDeploymentConditionType = "Frozen"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition;
// "ConditionFalse" means a resource is not in the condition; "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// MachinePhase is the ongoing-phase machine is going through at the moment.
type MachinePhase string

// These are the valid values for the MachinePhase.
const (
	// MachinePending should be set when machine is being created but it has not joined the
	// cluster yet.
	MachinePending MachinePhase = "Pending"

	// MachineRunning should be set when machine has joined the cluster and running successfully.
	MachineRunning MachinePhase = "Running"

	// MachineTerminating should be set when machine is being deleted.
	MachineTerminating MachinePhase = "Terminating"

	// MachineUnknown should be set when there is no way for controller to understand the
	// current phase of the machine. For example when is not reachable due to network-issues.
	MachineUnknown MachinePhase = "Unknown"

	// MachineFailed should be set when machine has failed and certain action needs to be taken either by higher-level controller or an admin.
	MachineFailed MachinePhase = "Failed"
)

// MachineOperationType is the type of the last operation performed on the Machine.
type MachineOperationType string

// These are valid values for MachineOperationType.
const (
	// MachineOperationCreate should be set when last operation was Create.
	MachineOperationCreate MachineOperationType = "Create"

	// MachineOperationUpdate should be set when last operation was Update.
	MachineOperationUpdate MachineOperationType = "Update"

	// MachineOperationHealthCheck should be set when last operation was HealthCheck.
	MachineOperationHealthCheck MachineOperationType = "HealthCheck"

	// MachineOperationDelete should be set when last operation was Delete.
	MachineOperationDelete MachineOperationType = "Delete"
)

// MachineOperationState is the current status of the last performed operation.
type MachineOperationState string

// These are the valid values for MachineOperationState.
const (
	// MachineOperationStateProcessing should be set when last operation is ongoing.
	MachineOperationStateProcessing MachineOperationState = "Processing"

	// MachineOperationStateFailed should be set when last operation has failed with certain error.
	MachineOperationStateFailed MachineOperationState = "Failed"

	// MachineOperationStateSuccessful should be set when last operation was successfully completed.
	MachineOperationStateSuccessful MachineOperationState = "Successful"
)
