package argument

import (
	`testing`
)

func TestAddFlag(t *testing.T) {
	const (
		long_form  = `foo`
		short_form = `f`
		arg_type   = TYPE_INT
		desc       = `foo desc`
	)

	action, _ := NewAction(
		`bar`,
		`b`,
		TYPE_UINT,
		`bar desc`,
	)

	err := Adder(action).AddFlag(
		long_form,
		short_form,
		arg_type,
		desc,
	)

	if err != nil {
		t.Errorf(`err = "%v", want nil`, err)
	}

	if v, exists := action.flags[long_form]; !exists {
		t.Errorf(`action.flags["%v"] does not exist`, long_form)
	} else {
		if v.desc != desc {
			t.Errorf(`action.flags["%v"].desc = "%v", want "%v"`,
				desc,
				v.desc,
				desc,
			)
		}
	}

	if v, exists := action.flags[short_form]; !exists {
		t.Errorf(`action.flags["%v"] does not exist`, short_form)
	} else {
		if v.desc != desc {
			t.Errorf(`action.flags["%v"].desc = "%v", want "%v"`,
				short_form,
				v.desc,
				desc,
			)
		}
	}
}

// vim:ts=4 sw=4 noet:
