package combinator

import "fmt"

type Toto struct {
	Wow int
	Bla string
}

type OptToto func(*Toto)

//type OptIncreaseWowFuncType func(bonus int)

func (t *Toto) IncreaseWow(bonus int) *Toto {
	t.Wow = t.Wow + bonus
	return t
}

func (t *Toto) AnotherOne(superBla string) *Toto {
	t.Bla = superBla
	return t
}

func OptIncreaseWow(bonus int) OptToto {
	return func(toto *Toto) {
		toto.Wow = toto.Wow + bonus
	}
}

func OptCrazyStuff(myString string, anotherToto Toto, minus int) OptToto {
	return func(toto *Toto) {
		fmt.Printf(myString)
		toto.Bla = anotherToto.Bla
		toto.Wow = toto.Wow - minus
	}
}

func NothingToDoWithOptToto() int {
	return 42
}

func Amend(toto *Toto, optsToto ...OptToto) {
	for _, opt := range optsToto {
		opt(toto)
	}
}

func add(originInt int, ints ...int) int {
	var newInt int
	for _, i := range ints {
		newInt = originInt + i
	}
	return newInt
}
