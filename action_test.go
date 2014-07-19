package argument

import (
	`testing`
)

func TestNewAction(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `f`
		arg_type   = TYPE_BOOL
		desc       = `foo desc`
	)

	o, err := NewAction(
		long_form,
		short_form,
		arg_type,
		desc,
	)

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}

	if o.long_form != long_form {
		t.Errorf(`long_form = "%v", want "%v"`, long_form, o.long_form)
	}

	if o.short_form != short_form {
		t.Errorf(`short_form = "%v", want "%v"`, short_form, o.short_form)
	}

	if o.arg_type != arg_type {
		t.Errorf(`arg_type = %v, want %v`, arg_type, o.arg_type)
	}

	if o.desc != desc {
		t.Errorf(`desc = "%v", want "%v"`, desc, o.desc)
	}
}

// vim:ts=4 sw=4 noet:
