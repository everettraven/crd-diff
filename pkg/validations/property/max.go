package property

import (
	"cmp"
	"fmt"

	"github.com/everettraven/crd-diff/pkg/validations/results"
)

func maxVerification[T cmp.Ordered](older *T, newer *T) error {
	var err error
	switch {
	case older == nil && newer != nil:
		err = fmt.Errorf("constraint %v added when there were no restrictions previously", *newer)
	case older != nil && newer != nil && *newer < *older:
		err = fmt.Errorf("constraint decreased from %v to %v", *older, *newer)
	}
	return err
}

type Maximum struct{}

func (m *Maximum) Name() string {
	return "Maximum"
}

func (m *Maximum) Validate(diff Diff) (bool, *results.Result) {
	reset := func(diff Diff) Diff {
		oldProperty := diff.Old()
		newProperty := diff.New()
		oldProperty.Maximum = nil
		newProperty.Maximum = nil
		return NewDiff(oldProperty, newProperty)
	}

	err := maxVerification(diff.Old().Maximum, diff.New().Maximum)
	if err != nil {
		err = fmt.Errorf("maximum: %s", err.Error())
	}

	return IsHandled(diff, reset), &results.Result{
		Error:      err,
		Subresults: []*results.Result{},
	}
}

type MaxItems struct{}

func (m *MaxItems) Name() string {
	return "MaxItems"
}

func (m *MaxItems) Validate(diff Diff) (bool, *results.Result) {
	reset := func(diff Diff) Diff {
		oldProperty := diff.Old()
		newProperty := diff.New()
		oldProperty.MaxItems = nil
		newProperty.MaxItems = nil
		return NewDiff(oldProperty, newProperty)
	}

	err := maxVerification(diff.Old().MaxItems, diff.New().MaxItems)
	if err != nil {
		err = fmt.Errorf("maxItems: %s", err.Error())
	}

	return IsHandled(diff, reset), &results.Result{
		Error:      err,
		Subresults: []*results.Result{},
	}
}

type MaxLength struct{}

func (m *MaxLength) Name() string {
	return "MaxLength"
}

func (m *MaxLength) Validate(diff Diff) (bool, *results.Result) {
	reset := func(diff Diff) Diff {
		oldProperty := diff.Old()
		newProperty := diff.New()
		oldProperty.MaxLength = nil
		newProperty.MaxLength = nil
		return NewDiff(oldProperty, newProperty)
	}

	err := maxVerification(diff.Old().MaxLength, diff.New().MaxLength)
	if err != nil {
		err = fmt.Errorf("maxLength: %s", err.Error())
	}

	return IsHandled(diff, reset), &results.Result{
		Error:      err,
		Subresults: []*results.Result{},
	}
}

type MaxProperties struct{}

func (m *MaxProperties) Name() string {
	return "MaxProperties"
}

func (m *MaxProperties) Validate(diff Diff) (bool, *results.Result) {
	reset := func(diff Diff) Diff {
		oldProperty := diff.Old()
		newProperty := diff.New()
		oldProperty.MaxProperties = nil
		newProperty.MaxProperties = nil
		return NewDiff(oldProperty, newProperty)
	}

	err := maxVerification(diff.Old().MaxProperties, diff.New().MaxProperties)
	if err != nil {
		err = fmt.Errorf("maxProperties: %s", err.Error())
	}

	return IsHandled(diff, reset), &results.Result{
		Error:      err,
		Subresults: []*results.Result{},
	}
}
