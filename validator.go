package argument

import (
	`fmt`
)

type Validator interface {
	validate() error
}

func (self *Flag) validate() error {
	if len(self.long_form) < 1 ||
		len(self.short_form) < 1 {
		return fmt.Errorf(`long_form and short_form must not be empty`)
	}

	if self.long_form[0] == '-' ||
		self.short_form[0] == '-' {
		return fmt.Errorf(`long_form and short_form cannot start with a dash`)
	}

	if self.arg_type != TYPE_INT &&
		self.arg_type != TYPE_UINT &&
		self.arg_type != TYPE_BOOL &&
		self.arg_type != TYPE_STRING {
		return fmt.Errorf(`invalid arg_type`)
	}
	return nil
}

// vim:ts=4 sw=4 noet
