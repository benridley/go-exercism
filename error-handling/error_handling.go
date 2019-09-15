package erratum

// Use opens a resource and calls Frob, handling any potential errors.
func Use(o ResourceOpener, input string) (e error) {
	var res Resource

	for res, e = o(); e != nil; res, e = o() {
		if _, ok := e.(TransientError); !ok {
			return e
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
