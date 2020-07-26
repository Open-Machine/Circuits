package main

import "fmt"

type pair struct {
	x int
	y int
}

func newPair(x int, y int) pair {
	return pair{x, y}
}

type wire struct {
	from pair
	to   pair
}

func newWire(from pair, to pair) wire {
	return wire{from, to}
}

func (w wire) toCircuit() string {
	return fmt.Sprintf("<wire from=\"(%d,%d)\" to=\"(%d,%d)\"/>", w.from.x, w.from.y, w.to.x, w.to.y)
}

type pin struct {
	loc  pair
	name string
}

func newPin(loc pair, name string) pin {
	return pin{loc: loc, name: name}
}

func (p pin) toCircuit() string {
	return fmt.Sprintf("<comp lib=\"0\" loc=\"(%d,%d)\" name=\"Pin\"> %s <a name=\"label\" val=\"%s\"/> %s </comp>", p.loc.x, p.loc.y, newLine, p.name, newLine)
}

type ramByte struct {
	loc pair
}

func newRAMByte(loc pair) ramByte {
	return ramByte{loc: loc}
}

func (r ramByte) toCircuit() string {
	return fmt.Sprintf("<comp loc=\"(%d,%d)\" name=\"byte_ram\"> %s <a name=\"facing\" val=\"south\"/> %s </comp>", r.loc.x, r.loc.y, newLine, newLine)
}
