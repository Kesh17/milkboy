package cpu

type CPU struct {
	register Register

	pc uint16
	sp uint16
}
