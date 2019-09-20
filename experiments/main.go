package main

import (
	"github.com/hervit0/jab/experiments/combinator"
	"log"
)

func main() {
	myToto := combinator.Toto{
		Wow: 123,
		Bla: "abc",
	}

	theOtherToto := combinator.Toto{
		Wow: 0,
		Bla: "qwertyuiop",
	}

	combinator.Amend(
		&myToto,
		combinator.OptIncreaseWow(2),
		combinator.OptIncreaseWow(3),
		combinator.OptCrazyStuff("my log line", theOtherToto, 10),
	)

	// This is not valid
	//amend(&myToto, combinator.NothingToDoWithOptToto())

	theOtherToto.IncreaseWow(23).IncreaseWow(22).AnotherOne("lol")

	log.Printf("%+v", myToto)
	log.Printf("%+v", theOtherToto)
}
