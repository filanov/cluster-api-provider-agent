/*
Copyright 2021.

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

package controllers

import (
	"context"
	"time"

	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	logutil "github.com/openshift/assisted-service/pkg/log"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	capiproviderv1alpha1 "github.com/eranco74/cluster-api-provider-agent/api/v1alpha1"
)

const defaultRequeueAfterOnError = 10 * time.Second

// AgentMachineReconciler reconciles a AgentMachine object
type AgentMachineReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logrus.FieldLogger
}

//+kubebuilder:rbac:groups=capi-provider.agent-install.openshift.io,resources=agentmachines,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=capi-provider.agent-install.openshift.io,resources=agentmachines/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=capi-provider.agent-install.openshift.io,resources=agentmachines/finalizers,verbs=update

func (r *AgentMachineReconciler) Reconcile(originalCtx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx := addRequestIdIfNeeded(originalCtx)
	log := logutil.FromContext(ctx, r.Log).WithFields(
		logrus.Fields{
			"agent_machine":           req.Name,
			"agent_machine_namespace": req.Namespace,
		})

	defer func() {
		log.Info("InfraEnv Reconcile ended")
	}()

	agentMachine := &capiproviderv1alpha1.AgentMachine{}
	if err := r.Get(ctx, req.NamespacedName, agentMachine); err != nil {
		log.WithError(err).Errorf("Failed to get agentMachine %s", req.NamespacedName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// If the AgentMachine is ready, we have nothing to do
	if agentMachine.Status.Ready == true {
		return ctrl.Result{}, nil
	}
	// If I don't have an agent, find one and set 'agentRef'
	if agentMachine.Status.AgentRef == nil {
		return findAgent(ctx, log, agentMachine)
	}

	// If I have an agent, update its IP addresses if necessary

	// If I have an agent, check its conditions and update 'ready'

	// If I have an agent in error, update 'failureReason' and 'failureMessage'

	return ctrl.Result{}, nil
}

func (r *AgentMachineReconciler) findAgent(ctx context.Context, log logrus.FieldLogger, agentMachine *capiproviderv1alpha1.AgentMachine) (ctrl.Result, error) {
	errReply := ctrl.Result{RequeueAfter: defaultRequeueAfterOnError}
	agents := &aiv1beta1.AgentList{}
	if err := r.List(ctx, agents); err != nil {
		return errReply, err
	}
	for i, agent := range agents.Items {
		// Take the first free agent that we find for now
		if agent.Spec.ClusterDeploymentName == nil {
			agentMachine.Status.AgentRef.Name = agent.Name
			agentMachine.Status.AgentRef.Namespace = agent.Namespace
			agentMachine.Spec.ProviderID = "agent://" + agent.Status.Inventory.SystemVendor.SerialNumber
		}
	}
	return ctrl.Result{}, nil
}

func getClusterDeploymentFromAgent(agent)

// SetupWithManager sets up the controller with the Manager.
func (r *AgentMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&capiproviderv1alpha1.AgentMachine{}).
		Complete(r)
}
