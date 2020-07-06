# Basic-Computer-with-Circuits
Creating a basic computer using only the following components: clock and AND/OR/XOR ports.

# About computer
### Internal components
- CPU
  - Control Unit (CU)
  - Arithmetic Logic Unit (ALU)
  - Registers:
    - PC
    - AC
    - EE
- RAM
- Clock

### How does it work?
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

## Run
```sh
java -jar logisim-generic-2.7.1.jar
```
