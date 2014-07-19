package argument

type Action struct {
	Flag
	flags map[string]*Flag
}

func NewAction(
	long_form string,
	short_form string,
	arg_type uint8,
	desc string,
) (*Action, error) {

	self := &Action{
		Flag: Flag{
			long_form:  long_form,
			short_form: short_form,
			arg_type:   arg_type,
			desc:       desc,
		},
		flags: map[string]*Flag{},
	}

	if err := Validator(self).validate(); err != nil {
		return nil, err
	}

	return self, nil
}

// vim:ts=4 sw=4 noet
