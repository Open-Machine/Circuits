package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestProgram(t *testing.T) {
	stdout, err := runProgram("program")

	if err != nil {
		os.Exit(1)
	}

	expected := []string{"1335", "EECD", "1234", "f000", "0000", "f000", "0001", "0000", "0001", "8888"}

	scanner := bufio.NewScanner(stdout)
	scanner.Scan() // ignores first print
	var i = 0
	for scanner.Scan() {
		line := scanner.Text()

		binaryStr := strings.ReplaceAll(line, " ", "")[0:16]
		gotNum, errBinary := binaryStringToNumber(binaryStr)

		expectedNum, errHex := hexStringToNumber(expected[i])

		if errBinary != nil || errHex != nil {
			t.Errorf("Parsing error. ErrBinary: '%t', ErrHex: '%t'.", errBinary, errHex)
		} else if gotNum != expectedNum {
			t.Errorf("Got different then expected. Expected: %d, but got: %d.\n", expectedNum, gotNum)
		}

		i++
		if len(expected) == i {
			break
		}
	}

	stdout.Close()
}

func TestBinaryStringToNumber(t *testing.T) {
	tests := []struct {
		param  string
		expect uint64
	}{
		{param: "0000000000000001", expect: 1},
		{param: "1000000000000001", expect: 32769},
		{param: "0001001100110101", expect: 4917},
		{param: "1111111111111111", expect: 65535},
	}

	for i, test := range tests {
		got, err := binaryStringToNumber(test.param)
		if err != nil {
			t.Errorf("[%d] Error: %t", i, err)
		} else if got != test.expect {
			t.Errorf("[%d] Got: %d | Expected: %d", i, got, test.expect)
		}
	}
}

func TestHexStringToNumber(t *testing.T) {
	tests := []struct {
		param  string
		expect uint64
	}{
		{param: "0001", expect: 1},
		{param: "8001", expect: 32769},
		{param: "1335", expect: 4917},
		{param: "ffff", expect: 65535},
	}

	for i, test := range tests {
		got, err := hexStringToNumber(test.param)
		if err != nil {
			t.Errorf("[%d] Error: %t", i, err)
		} else if got != test.expect {
			t.Errorf("[%d] Got: %d | Expected: %d", i, got, test.expect)
		}
	}
}
