package domain

type Address struct {
	Id     string
	Street string
	Number string
	CP     string
	State  string
}

func (a *Address) IsValid() bool {
	if a.Id == "" {
		return false
	}
	if a.Street == "" {
		return false
	}
	if a.Number == "" {
		return false
	}
	if a.CP == "" {
		return false
	}
	if a.State == "" {
		return false
	}

	return true
}
