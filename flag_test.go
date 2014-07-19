package argument

import (
	`testing`
)

func TestNewFlag(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `f`
		arg_type   = TYPE_STRING
		desc       = `foo desc`
	)

	flag, err := NewFlag(
		long_form,
		short_form,
		arg_type,
		desc,
	)

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}

	if flag.long_form != long_form {
		t.Errorf(`flag.long_form = "%v", want "%v"`, flag.long_form, long_form)
	}

	if flag.short_form != short_form {
		t.Errorf(`flag.short_form = "%v", want "%v"`, flag.short_form, short_form)
	}

	if flag.arg_type != arg_type {
		t.Errorf(`flag.arg_type = %v, want %v`, flag.arg_type, arg_type)
	}

	if flag.desc != desc {
		t.Errorf(`flag.desc = "%v", want "%v"`, flag.desc, desc)
	}
}

// vim:ts=4 sw=4 noet:
