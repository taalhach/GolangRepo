package types

type Matcher interface {
	Print()
}

func CallIn(m Matcher)  {
	m.Print()
}

