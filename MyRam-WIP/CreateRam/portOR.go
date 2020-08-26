package main

import (
	"errors"
	"fmt"
)

const maxInputsPort = 32

const xInputs = 150
const yFirstPort = 300
const yDistanceBwtInputs = 10

const spaceBtwPorts = 100
const xDistanceInputPort = 100
const yDistancePortFromFirstInput = 160
const widthPort = 50

const imenseOrInputNamePrefix = "input"

func generateOrPort(amntInputs int) (string, error) {
	imenseOrPort := newImenseOrPort()

	if amntInputs > maxInputsPort*maxInputsPort {
		return "", errors.New("This port is not able to deal with this much inputs")
	}

	amntInputsRemaining := amntInputs
	yCurPort := yFirstPort
	for amntInputsRemaining > 0 {
		var amntInputsCurPort int
		if amntInputsRemaining > maxInputsPort {
			amntInputsCurPort = maxInputsPort
		} else {
			amntInputsCurPort = amntInputsRemaining
		}

		portLoc := newPair(xInputs+xDistanceInputPort, yCurPort)
		imenseOrPort.addOrPort(portLoc)

		for i := 0; i < amntInputsCurPort; i++ {
			y := yCurPort - yDistancePortFromFirstInput + i*yDistanceBwtInputs
			if i >= maxInputsPort/2 {
				y += yDistanceBwtInputs
			}

			inputPinLoc := newPair(xInputs, y)
			imenseOrPort.addInputPins(inputPinLoc)

			wireFrom := inputPinLoc
			wireTo := newPair(portLoc.x-widthPort, inputPinLoc.y)
			imenseOrPort.addWire(newWire(wireFrom, wireTo))
		}

		yCurPort += spaceBtwPorts + (maxInputsPort+1)*yDistanceBwtInputs
		amntInputsRemaining -= maxInputsPort
	}

	return imenseOrPort.toCircuit(), nil
}

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
///////////////////////////////// PORT //////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

type portType string

const (
	orPort  = "OR Gate"
	andPort = "AND Gate"
	notPort = "NOT Gate"
)

type direction string

const (
	south = "south"
	north = "north"
	east  = "east"
	west  = "west"
)

type port struct {
	loc             pair
	amntInputs      int
	portType        portType
	facingDirection direction
}

func newPort(loc pair, amntInputs int, portType portType, facingDirection direction) port {
	return port{loc: loc, amntInputs: amntInputs, portType: portType, facingDirection: facingDirection}
}

func (p port) toCircuit() string {
	ret := fmt.Sprintf("<comp lib=\"1\" loc=\"(%d,%d)\" name=\"%s\">", p.loc.x, p.loc.y, p.portType)
	ret += newLine + fmt.Sprintf("<a name=\"facing\" val=\"%s\"/>", p.facingDirection)
	ret += newLine + fmt.Sprintf("<a name=\"inputs\" val=\"%d\"/>", p.amntInputs)
	ret += newLine + "</comp>"
	return ret
}

/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
//////////////////////////// IMENSE OR PORT /////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////

type imenseOrPort struct {
	ports     []port
	wires     []wire
	inputPins []pin
	outputPin pin
}

func newImenseOrPort() imenseOrPort {
	return imenseOrPort{}
}

func (p *imenseOrPort) addOrPort(loc pair) {
	port := newPort(loc, maxInputsPort, orPort, east)
	p.ports = append(p.ports, port)
}

func (p *imenseOrPort) addWire(wire wire) {
	p.wires = append(p.wires, wire)
}

func (p *imenseOrPort) addInputPins(loc pair) {
	pinName := fmt.Sprintf("%s%d", imenseOrInputNamePrefix, len(p.inputPins))
	p.inputPins = append(p.inputPins, newPin(loc, pinName))
}

func (p *imenseOrPort) setOutputPins(pin pin) {
	p.outputPin = pin
}

func (p imenseOrPort) toCircuit() string {
	portsStr := ""
	for _, port := range p.ports {
		portsStr += port.toCircuit() + newLine
	}

	wiresStr := ""
	for _, wire := range p.wires {
		wiresStr += wire.toCircuit() + newLine
	}

	inputsStr := ""
	for _, pin := range p.inputPins {
		inputsStr += pin.toCircuit() + newLine
	}

	outputStr := p.outputPin.toCircuit()

	return portsStr + wiresStr + outputStr + inputsStr
}
