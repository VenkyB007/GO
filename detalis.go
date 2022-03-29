package data

type Doctor struct {
	ID        int
	Name      string
	Available string
}

func GetDetails() []*Doctor {
	return doctorslist
}

var doctorslist = []*Doctor{
	&Doctor{
		ID:        1,
		Name:      "venky",
		Available: "from 10AM to 12PM",
	},
	&Doctor{
		ID:        2,
		Name:      "vamsi",
		Available: "from 4PM to 6PM",
	},
}
