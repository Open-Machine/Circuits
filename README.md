# Circuit
A computer basic circuit using made only using logic ports.

*This repository is part of a bigger project: developing a computer from scratch. Check it out [here](https://github.com/Open-Machine/)!*

#### Todo
- [ ] Signed integer
- [ ] Float
- [ ] Division circuit and command
- [ ] Own Registers
- [ ] Clock
- [ ] Own RAM

// TODO: add images of the Circuit for each

// TODO: it's too basic and if I was talking to a child

---

# 🔖 Table of Contents
- [Circuit](#circuit)
- [🔖 Table of Contents](#-table-of-contents)
- [[WIP] How does a computer work behind the curtains?](#wip-how-does-a-computer-work-behind-the-curtains)

---

# [WIP] How does a computer work behind the curtains?
Developers usually have a good understanding of how a computer works. However, for me at least understanding a computer so deeply that you would be able to actually build one yourself seemed like an impossible task. So, in this section I want to break to you the most important pieces of the computer puzzle that you need to know in order to build your own computer from scratch using only circuits.

## Circuit Components

#### Central Processing Unit (CPU)
The CPU is the brain of the computer, where every instruction is processed. It is composed by the sub-components below:
  - *Control Unit (CU)*: it controls the flags of the entire CPU
  - *Arithmetic Logic Unit (ALU)*: it performs sum, subtracted and comparisons between numbers
  - *Registers*: it's a small memory that stays inside the CPU for faster 

#### Random Access Memory (RAM)
RAM is a memory fast, but temporary

#### Disk
RAM is a fast, but temporary

#### Input/Output Devices
For an actual computer, you also need input and output devices such as 

---

## Code

Line of code = Instruction (2 bytes) + Memory Address (2 bytes).

ps: The memory address in the lines of code will be called EE - [EE] represents EE value and EE


## Instructions Table
### Symbols Legend for the Instructions Table
Symbol | Explanation
--- | ---
ACC | The ACC register
EE | Represents a memory index
[ ] | "Value of"
### Instructions Table
Machine Code | Short Instruction Description | Long Instruction Description | Short Param Description | Long Param Description
--- | --- | --- | --- | ---
0x0 | - | This instruction doesn't perform any action | - | No parameter is required
0x1 | [ACC] = [EE] | A value from the memory is copied to the ACC register | variable | Memory index of a variable that will be used in the instruction
0x2 | [EE] = [ACC] | The value from the ACC register is stored into memory | variable | Memory index of a variable that will be used in the instruction
0x3 | [ACC] = [ACC] + [EE] | The sum of the value of the ACC register and a value from the memory is stored in the ACC register | variable | Memory index of a variable that will be used in the instruction
0x4 | [ACC] = [ACC] - [EE] | The difference between the value of the ACC register and a value from the memory is stored in the ACC register | variable | Memory index of a variable that will be used in the instruction
0x7 | [EE] = to the input value | The input value is copied to the memory | variable | Memory index of a variable that will be used in the instruction
0x8 | Output [EE] | Outputs a value from the memory into the circuit LEDs | variable | Memory index of a variable that will be used in the instruction
0x9 | Finishes program | When this instruction is encountered, the program is finished and no more instructions will be executed | - | No parameter is required
0xa | Jump to EE | Jump to another line of code | instruction | Memory index of a instruction the program will jump to
0xb | Jump to EE if [ACC] > 0 | Jump to another line of code if the value of the ACC register is positive | instruction | Memory index of a instruction the program will jump to if the condition is right
0xd | Jump to EE if [ACC] = 0 | Jump to another line of code if the value of the ACC register is zero | instruction | Memory index of a instruction the program will jump to if the condition is right
0xf | Jump to EE if [ACC] < 0 | Jump to another line of code if the value of the ACC register is negative | instruction | Memory index of a instruction the program will jump to if the condition is right

## Machine Code Example
```sh
01ff # copy value in the address ff in RAM
020a # stores the value of AC in the address 0a
0900 # kills program
```

## Tips
- Remember to add the ```0x9``` instruction at the end of your programs

---

## ▶️ Execute the machine

1. Run the machine:
	```sh
	java -jar logisim-evolution.jar
	```
2. Import the circuit file by navigating the menu:
   *```File -> Open -> Select main.circ from the repository folder```*
3. Open Main file on the left side of Logisim
4. Paste the executable code into the beginning of the RAM. You may want to change some values of the memory as if you were initializing variables.
5. Run the Program by navigating the menu:
   *```Simulate -> Enable 'Ticks Enabled'```*
   - You can change the speed of the program by navigating the menu: 
	*```Simulate -> Tick Frequency -> To get the fastest execution, select the top item```*
