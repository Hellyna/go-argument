package argument

const (
	TYPE_INT uint8 = iota
	TYPE_UINT
	TYPE_BOOL
	TYPE_STRING
)

func (self *ArgumentSet) AddFlag(
	long_form string,
	short_form string,
	arg_type uint8,
	desc string,
) error {
	return Adder(self).AddFlag(
		long_form,
		short_form,
		arg_type,
		desc,
	)
}

// vim:ts=4 sw=4 noet:
