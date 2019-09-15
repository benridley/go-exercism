package erratum

// Use opens a resource and calls Frob, handling any potential errors.
func Use(o ResourceOpener, input string) (e error) {
	var res Resource
L:
	for {
		tres, err := o()
		switch err.(type) {
		case nil:
			res = tres
			break L
		case TransientError:
			continue L
		case error:
			return err
		default:
			continue L
		}
	}
	defer res.Close()
	defer func() {
		r := recover()
		if fe, ok := r.(FrobError); ok {
			res.Defrob(fe.defrobTag)
			e = fe
		} else if err, ok := r.(error); ok {
			e = err
		}
	}()
	res.Frob(input)
	return nil
}
