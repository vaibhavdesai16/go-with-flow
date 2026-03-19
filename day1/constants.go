package day1

func PlayingWithConstants() {
	// basic
	const x = 10
	const y int = 15

	// declaring multiple
	const (
		a = 1
		b = 2
		c = 3
	)

	const (
		A = iota // 0
		B        // 1
		C        // 2
	)

	type Status int
	const (
		Pending Status = iota
		Approved
		Rejected
	)
}
