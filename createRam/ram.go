package main

import (
	"errors"
	"fmt"
	"log"
)

const ramInputNamePrefix = "input"

const xPins = 100
const yFirstPin = 120
const yDistanceBtwPins = 40

const yByteRams = 870
const xFirstByteRAM = 840
const xDistanceBtwByteRams = 250
const xDeltaStartWireFromByteRams = -180
const xDistanceBtwByteRAMInputs = 20
const yDeltaStartWireFromByteRams = -160

func getInputBeforeBits() []string {
	inputs := []string{"is_write", "select_value"}
	for i := 0; i < amntBitsComputer; i++ {
		inputs = append(inputs, fmt.Sprintf("%s%d", ramInputNamePrefix, i))
	}
	return inputs
}

// WIRE FROM BYTE RAM HAS TO START IN BYTE_RAM AND GO UP
// BYTE RAM HAVE TO BE ADDED IN ORDER (from left to right)

func generateRAMCircuit(amntBytes int) string {
	if amntBytes >= getMaxBytesRAM() {
		log.Fatal(errors.New("Computer won't be able to access these bytes"))
	}

	inputs := getInputBeforeBits()
	ram := newRAM(len(inputs))

	for i := 0; i < len(inputs); i++ {
		locPin := newPair(xPins, yFirstPin+i*yDistanceBtwPins)
		curPin := newPin(locPin, inputs[i])
		ram.addPin(curPin)

		startWire := locPin
		ram.addWireFromPin(startWire, i)
	}

	for i := 0; i < amntBytes; i++ {
		locByteRAM := newPair(xFirstByteRAM+i*xDistanceBtwByteRams, yByteRams)
		byteRAM := newRAMByte(locByteRAM)
		ram.addRAMByte(byteRAM)

		for i := 0; i < len(inputs); i++ {
			xWire := (locByteRAM.x + xDeltaStartWireFromByteRams) + i*xDistanceBtwByteRAMInputs
			yWire := byteRAM.loc.y + yDeltaStartWireFromByteRams
			startWire := newPair(xWire, yWire)

			bitIndex := len(inputs) - i - 1
			ram.addWireFromByte(startWire, bitIndex)
		}
	}

	return ram.toCircuit()
}

type ram struct {
	pinWiresStart []pair
	wires         []wire
	inputPins     []pin
	ramBytes      []ramByte
	// TODO:
	// outputPins    []pin
	// orPorts       []port
}

func newRAM(amntPins int) ram {
	return ram{pinWiresStart: make([]pair, amntPins)}
}

func (r *ram) addWireFromPin(pinWireStart pair, i int) {
	r.pinWiresStart[i] = pinWireStart
}

func (r *ram) addWireFromByte(byteWireStart pair, i int) {
	intersectionWires := newPair(byteWireStart.x, r.pinWiresStart[i].y)

	wireFromByte := newWire(byteWireStart, intersectionWires)
	r.wires = append(r.wires, wireFromByte)

	fromPinWire := r.pinWiresStart[i]
	toPinWire := intersectionWires

	wireFromPin := newWire(fromPinWire, toPinWire)
	r.wires = append(r.wires, wireFromPin)
	r.pinWiresStart[i] = toPinWire
}

func (r *ram) addPin(pin pin) {
	r.inputPins = append(r.inputPins, pin)
}

func (r *ram) addRAMByte(ramByte ramByte) {
	r.ramBytes = append(r.ramBytes, ramByte)
}

func (r ram) toCircuit() string {
	wiresStr := ""
	for _, wire := range r.wires {
		wiresStr += wire.toCircuit() + newLine
	}

	pinsStr := ""
	for _, pin := range r.inputPins {
		pinsStr += pin.toCircuit() + newLine
	}

	ramBytesStr := ""
	for _, ramByte := range r.ramBytes {
		ramBytesStr += ramByte.toCircuit() + newLine
	}

	return wiresStr + pinsStr + ramBytesStr
}
