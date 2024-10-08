package token

import (
	"github.com/rancher/apiserver/pkg/types"

	mgmtv1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
)

const LabelTokenIsCurrent = "auth.management.llmos.ai/is-current"

func ConvertTokenListToAPIObjectList(tokens []*mgmtv1.Token, session *mgmtv1.Token) []types.APIObject {
	result := make([]types.APIObject, 0, len(tokens))
	for _, tk := range tokens {
		if tk.Name == session.Name {
			tk.Labels[LabelTokenIsCurrent] = "true"
		}

		result = append(result, types.APIObject{
			Type:   tokenSchemaID,
			ID:     tk.Name,
			Object: tk,
		})
	}

	return result
}
