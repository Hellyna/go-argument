package argument

type Parser interface {
	parse() error
}

/*
func (self *_Action) parse_flag(args []string) error {
  for k,v := range args {

    flag,flag_exists := self.flags[
  }
  return nil
}

func (self *ArgumentSet) Parse() error {
  return nil
}

*/
// vim:ts=4 sw=4
