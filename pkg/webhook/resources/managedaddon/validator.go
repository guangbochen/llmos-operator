package managedaddon

import (
	"fmt"

	"github.com/oneblock-ai/webhook/pkg/server/admission"
	admissionregv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/runtime"

	mgmtv1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/constant"
	werror "github.com/llmos-ai/llmos-operator/pkg/webhook/error"
)

type validator struct {
	admission.DefaultValidator
}

var _ admission.Validator = &validator{}

func NewValidator() admission.Validator {
	return &validator{}
}

func (v *validator) Create(_ *admission.Request, obj runtime.Object) error {
	addon := obj.(*mgmtv1.ManagedAddon)
	if addon.Spec.Chart == "" {
		return werror.BadRequest("Chart name can't be empty")
	}
	if addon.Spec.Repo == "" {
		return werror.BadRequest("Repo can't be empty")
	}
	if addon.Spec.Version == "" {
		return werror.BadRequest("Version can't be empty")
	}
	return nil
}

func (v *validator) Update(_ *admission.Request, _ runtime.Object, obj runtime.Object) error {
	addon := obj.(*mgmtv1.ManagedAddon)
	if isSystemAddon(addon) && !addon.Spec.Enabled {
		return werror.MethodNotAllowed(fmt.Sprintf("Can't disable system addon %s", addon.Name))
	}
	return nil
}

func (v *validator) Delete(_ *admission.Request, obj runtime.Object) error {
	addon := obj.(*mgmtv1.ManagedAddon)
	if isSystemAddon(addon) {
		return werror.MethodNotAllowed(fmt.Sprintf("Can't delete system addon %s", addon.Name))
	}
	return nil
}

func isSystemAddon(addon *mgmtv1.ManagedAddon) bool {
	if addon.Labels != nil && addon.Labels[constant.SystemAddonLabel] == "true" {
		return true
	}
	return false
}

func (v *validator) Resource() admission.Resource {
	return admission.Resource{
		Names:      []string{"managedaddons"},
		Scope:      admissionregv1.NamespacedScope,
		APIGroup:   mgmtv1.SchemeGroupVersion.Group,
		APIVersion: mgmtv1.SchemeGroupVersion.Version,
		ObjectType: &mgmtv1.ManagedAddon{},
		OperationTypes: []admissionregv1.OperationType{
			admissionregv1.Create,
			admissionregv1.Delete,
			admissionregv1.Update,
		},
	}
}
