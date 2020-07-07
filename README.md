# Computer-from-Scratch
Designing a computer from scratch. The only components that I didn't made were the ports.

#### Todo
- [ ] Control Unit
- [ ] Signed integer
- [ ] Float
- [ ] Division

## Objective
Understand how a computer works behind the curtains and maybe do some things my way.

## Run
1. Run circuit software
```sh
java -jar logisim-generic-2.7.1.jar
```
2. File -> Open -> Select main.circ
3. Open Main
4. Program - [How do I program this computer?](#code)

## How does a computer work?
### Internal components
- CPU
  - Control Unit (CU)
  - Arithmetic Logic Unit (ALU)
  - Registers: PC, AC, EE
- RAM
- Clock

### Code
The code should be written in RAM and it will be executed from the memmory address in register CP content. Every 4 bytes are considered a line of code.
<br/>Line of code = Instruction (2 bytes) + Memory Address (2 bytes).
ps: The memory address in the lines of code will be called EE - [EE] represents EE value and EE
##### Instructions Table
Code (hexadecimal number) | Description
--- | ---
00 | No operation
01 | [AC] = [EE]
02 | [EE] = [AC]
03 | [AC] = [AC] + [EC]
04 | [AC] = [AC] - [EC]
07 | [EE] = to the input value
08 | Output [EE]
09 | Finish program
10 | Jump to EE
11 | Jump to EE if [AC] > 0
13 | Jump to EE if [AC] = 0
15 | Jump to EE if [AC] < 0
##### Example
```sh
01ff #copy value in the address ff in RAM
```
