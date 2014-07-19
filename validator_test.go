package argument

import (
	`testing`
)

func Test_validate_valid(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}
}

func Test_validate_valid0(t *testing.T) {
	const (
		long_form  = `foo-bar-bar`
		short_form = `f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}
}

func Test_validate_valid1(t *testing.T) {
	const (
		long_form  = `foo123&&%%$bar-bar`
		short_form = `f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}
}

func Test_validate_long_form_extra_dash(t *testing.T) {
	const (
		long_form  = `-foo`
		short_form = `f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err == nil {
		t.Errorf(`err = nil, want err != nil`)
	}
}

func Test_validate_short_form_extra_dash(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `-f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err == nil {
		t.Errorf(`err = nil, want err != nil`)
	}
}

func Test_validate_both_forms_extra_dash(t *testing.T) {
	const (
		long_form  = `-foo`
		short_form = `-f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err == nil {
		t.Errorf(`err = nil, want err != nil`)
	}
}

func Test_validate_arg_type_out_of_range(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `f`
		arg_type   = 255
		desc       = `foo desc`
	)

	flag := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	err := Validator(flag).validate()

	if err == nil {
		t.Errorf(`err = nil, want err != nil`)
	}
}

// vim:ts=4 sw=4 noet
