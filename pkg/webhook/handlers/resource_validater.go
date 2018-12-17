package handlers

import (
	"context"
	"log"

	//"net/http"

	//	"strings"

	//	corev1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

type ResouceValidater struct {
	client  client.Client
	decoder types.Decoder
}

// podAnnotator implements admission.Handler.
var _ admission.Handler = &ResouceValidater{}

func (a *ResouceValidater) Handle(ctx context.Context, req types.Request) types.Response {
	log.Printf("Validating Webhook Handle Request %s/ %s/ %s/\n", req.AdmissionRequest.Namespace, req.AdmissionRequest.Kind, req.AdmissionRequest.Operation)
	//resourceObj := &metav1.ObjectMeta{}
	//err := a.decoder.Decode(req, resourceObj)

	//if err != nil {
	//	return admission.ErrorResponse(http.StatusBadRequest, err)
	//}

	//if resourceObj.GetLabels() == nil || resourceObj.GetLabels()["appid"] == "" {
	//	return admission.ValidationResponse(false, "appid is not allowed null")
	//}

	return admission.ValidationResponse(true, "ok")
}

// podAnnotator implements inject.Client.
var _ inject.Client = &ResouceValidater{}

// InjectClient injects the client into the podAnnotator
func (a *ResouceValidater) InjectClient(c client.Client) error {
	a.client = c
	return nil
}

// podAnnotator implements inject.Decoder.
var _ inject.Decoder = &ResouceValidater{}

// InjectDecoder injects the decoder into the podAnnotator
func (a *ResouceValidater) InjectDecoder(d types.Decoder) error {
	a.decoder = d
	return nil
}
