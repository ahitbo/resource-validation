package handlers

import (
	"context"
	"log"

	//"net/http"

	//	"strings"
	//"github.com/golang/protobuf/proto"
	"encoding/json"

	//corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

type ResouceValidater struct {
	client  client.Client
	decoder types.Decoder
}

// podAnnotator implements admission.Handler.
var _ admission.Handler = &ResouceValidater{}

func (a *ResouceValidater) Handle(ctx context.Context, req types.Request) types.Response {
	log.Printf("Validating Webhook Handle Request Namespace=%s/ Kind=%s/ Operation=%s/\n", req.AdmissionRequest.Namespace, req.AdmissionRequest.Kind, req.AdmissionRequest.Operation)

	r := Resource{}
	err := json.Unmarshal(req.AdmissionRequest.Object.Raw, &r)
	if err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	} else {
		labels := r.GetLabels()
		isPrd := labels["is-prd-deploy"]
		if isPrd != "" && isPrd == "true" {
			appId := labels["appid"]
			changeNo := labels["changeno"]
			if appId == "" || changeNo == "" {
				return admission.ValidationResponse(true, "ok")
			}

		}
	}

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
