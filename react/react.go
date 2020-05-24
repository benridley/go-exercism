package react

type ReactorCell struct {
	value int
}

func (r *ReactorCell) Value() int {
	return r.value
}

type ReactorInputCell struct {
	cell ReactorCell
}

func (r *ReactorInputCell) SetValue(n int) {
	r.cell.value = n
}

func (r *ReactorInputCell) Value() int {
	return r.cell.Value()
}

type ReactorComputeCell struct {
	cell     ReactorCell
	callback *func(int)
}

func (r *ReactorComputeCell) Cancel() {
	r.callback = nil
}

func (r *ReactorComputeCell) AddCallback(f func(int)) Canceler {
	r.callback = &f
	return r
}

type MyReactor struct {
}
