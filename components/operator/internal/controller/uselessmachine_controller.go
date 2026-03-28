package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mecrisv1alpha1 "github.com/yebyen/useless/operator/api/v1alpha1"
	"github.com/yebyen/useless/operator/internal/wasm"
)

// UselessMachineReconciler reconciles a UselessMachine object
type UselessMachineReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	WasmRunner *wasm.Runner
}

// +kubebuilder:rbac:groups=mecris.io,resources=uselessmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mecris.io,resources=uselessmachines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=mecris.io,resources=uselessmachines/finalizers,verbs=update

func (r *UselessMachineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	var machine mecrisv1alpha1.UselessMachine
	if err := r.Get(ctx, req.NamespacedName, &machine); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Prepare current state for WASM
	currentState := wasm.UselessMachineState{
		DailyCount: machine.Status.DailyCount,
		LastPushed: machine.Status.LastPushed,
	}
	nowStr := time.Now().Format(time.RFC3339)

	// Check if action is requested
	if machine.Spec.Action == "push" {
		l.Info("Action 'push' requested, invoking WASM Brain")
		nextState, err := r.WasmRunner.PushButton(ctx, currentState, nowStr)
		if err != nil {
			l.Error(err, "WASM Brain failed during push_button")
			return ctrl.Result{RequeueAfter: 5 * time.Second}, err
		}

		// Update machine status
		machine.Status.DailyCount = nextState.DailyCount
		machine.Status.LastPushed = nextState.LastPushed
		if err := r.Status().Update(ctx, &machine); err != nil {
			return ctrl.Result{}, err
		}
		
		// Clear action once processed
		machine.Spec.Action = ""
		if err := r.Update(ctx, &machine); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Always update nag status
	status, err := r.WasmRunner.GetStatus(ctx, currentState, nowStr)
	if err != nil {
		l.Error(err, "WASM Brain failed during get_status")
	} else {
		if machine.Status.IsNagging != status.IsNagging {
			machine.Status.IsNagging = status.IsNagging
			if err := r.Status().Update(ctx, &machine); err != nil {
				return ctrl.Result{}, err
			}
		}
	}

	// Reconcile periodically to check nag status
	return ctrl.Result{RequeueAfter: 1 * time.Hour}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UselessMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mecrisv1alpha1.UselessMachine{}).
		Complete(r)
}
