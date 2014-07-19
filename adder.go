package argument

type Adder interface {
	AddFlag(
		string,
		string,
		uint8,
		string) error
}

func (self *Action) AddFlag(
	long_form string,
	short_form string,
	arg_type uint8,
	desc string,
) error {

	flag, err := NewFlag(
		long_form,
		short_form,
		arg_type,
		desc,
	)

	if err != nil {
		return err
	}

	self.flags[long_form] = flag
	self.flags[short_form] = flag

	return nil
}

// vim:ts=4 sw=4 noet
