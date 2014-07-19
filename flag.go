package argument

type Flag struct {
	long_form  string
	short_form string
	desc       string
	arg_type   uint8
}

func NewFlag(
	long_form string,
	short_form string,
	arg_type uint8,
	desc string,
) (*Flag, error) {

	self := &Flag{
		long_form:  long_form,
		short_form: short_form,
		arg_type:   arg_type,
		desc:       desc,
	}

	if err := Validator(self).validate(); err != nil {
		return nil, err
	}

	return self, nil
}

// vim:ts=4 sw=4 noet:
