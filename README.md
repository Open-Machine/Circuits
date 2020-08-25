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

###### Instructions Table
Machine Code | Description | Requires Param
--- | --- | ---
0x0 | No operation | No
0x1 | [AC] = [EE] | Yes
0x2 | [EE] = [AC] | Yes
0x3 | [AC] = [AC] + [EC] | Yes
0x4 | [AC] = [AC] - [EC] | Yes
0x7 | [EE] = to the input value | Yes
0x8 | Output [EE] | Yes
0x9 | Finish program | No
0xa | Jump to EE | Yes
0xb | Jump to EE if [AC] > 0 | Yes
0xd | Jump to EE if [AC] = 0 | Yes
0xf | Jump to EE if [AC] < 0 | Yes
### Machine Code Example
```sh
01ff # copy value in the address ff in RAM
020a # stores the value of AC in the address 0a
0900 # kills program
```

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
