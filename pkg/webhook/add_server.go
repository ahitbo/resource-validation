package webhook

import (
	"resource-validation/pkg/webhook/handlers"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"

	//corev1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/builder"

	apitypes "k8s.io/apimachinery/pkg/types"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, Add)
}

func Add(mgr manager.Manager) error {
	return add(mgr)
}

func add(mgr manager.Manager) error {
	//创建ValidatingAdmissionWebhook
	vd, err := builder.NewWebhookBuilder().
		Validating().
		Rules(
			admissionregistrationv1beta1.RuleWithOperations{
				Operations: []admissionregistrationv1beta1.OperationType{admissionregistrationv1beta1.OperationAll},
				Rule: admissionregistrationv1beta1.Rule{
					APIGroups:   []string{"*"},
					APIVersions: []string{"*"},
					Resources:   []string{"*/*"}}}).
		Handlers(&handlers.ResouceValidater{}).
		WithManager(mgr).
		Build()
	if err != nil {
		return err
	}

	//创建Webhook Server
	svr, err := webhook.NewServer("application-admission-server", mgr, webhook.ServerOptions{
		CertDir: "/tmp/cert",
		BootstrapOptions: &webhook.BootstrapOptions{
			Secret: &apitypes.NamespacedName{
				Namespace: "crdapp-system",
				Name:      "crdapp-webhook-server-secret",
			},

			Service: &webhook.Service{
				Namespace: "crdapp-system",
				Name:      "crdapp-admission-server-service",
				// Selectors should select the pods that runs this webhook server.
				Selectors: map[string]string{
					"control-plane":           "controller-manager",
					"controller-tools.k8s.io": "1.0",
				},
			},
		},
	})

	if err != nil {
		return err
	}

	//注册
	svr.Register(vd)
	return nil
}
