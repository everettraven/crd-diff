package property

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestRequired(t *testing.T) {
	type testcase struct {
		name        string
		oldProperty *apiextensionsv1.JSONSchemaProps
		newProperty *apiextensionsv1.JSONSchemaProps
		err         error
		handled     bool
		required    *Required
	}

	for _, tc := range []testcase{
		{
			name: "no diff, no error, handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			err:      nil,
			handled:  true,
			required: &Required{},
		},
		{
			name:        "new required field, error, handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			err:      errors.New("new required fields [foo] added"),
			handled:  true,
			required: &Required{},
		},
		{
			name:        "new required field, new enforcement set to None, no error, handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			err:     nil,
			handled: true,
			required: &Required{
				NewEnforcement: RequiredValidationNewEnforcementNone,
			},
		},
		{
			name: "required field removed, error, handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo", "bar",
				}},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			err:     errors.New("required fields [bar] removed"),
			handled: true,
			required: &Required{
				NewEnforcement: RequiredValidationRemovalEnforcementStrict,
			},
		},
		{
			name: "required field removed, removal enforcement set to None, no error, handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo", "bar",
				}},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				Required: []string{
					"foo",
				},
			},
			err:     nil,
			handled: true,
			required: &Required{
				NewEnforcement: RequiredValidationNewEnforcementNone,
			},
		},
		{
			name: "different field changed, no error, not handled",
			oldProperty: &apiextensionsv1.JSONSchemaProps{
				ID: "foo",
			},
			newProperty: &apiextensionsv1.JSONSchemaProps{
				ID: "bar",
			},
			err:      nil,
			handled:  false,
			required: &Required{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, handled, err := tc.required.Validate(NewDiff(tc.oldProperty, tc.newProperty))
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.handled, handled)
		})
	}
}
