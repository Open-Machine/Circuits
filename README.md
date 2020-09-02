<h1 align="center">Open Machine's Circuits</i></h1>
<div align="center">

<a href="https://github.com/Open-Machine/Circuits/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Circuits" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Circuits" alt="Forks Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/"><img src="https://img.shields.io/badge/version-0.0.1-blue" alt="Version Badge"/></a>
<a href="https://github.com/Open-Machine/Assembler/commits/"><img src="https://img.shields.io/github/commit-activity/m/Open-Machine/Circuits" alt="commits"/></a>
<a href="https://github.com/Open-Machine/Circuits/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Circuits" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Circuits" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Circuits?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Circuits/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Circuits?color=2b9348" alt="License Badge"/></a>

<img src="https://raw.githubusercontent.com/Open-Machine/README/master/Media/logo-horizontal.png" alt="open-machine"/>

<br/>
<i>The goal is to to design and build a computer from scratch only using logical gates to do so.<i/>
<br/>

<i>This repository is part of a <a href="https://github.com/Open-Machine/">larger project<a/>: developing a computer from scratch with its <a href="https://github.com/Open-Machine/Assembler">assembler<a/> and compiler.

</div>

---

<!-- omit in toc -->
# 🔖 Table of Contents
### 1. [✔ Todo](#-todo)
### 2. [💻 How does a computer work behind the curtains? (WIP)](#-how-does-a-computer-work-behind-the-curtains-wip)
### 3. [🔢 Machine Code](#-machine-code)
### 4. [▶️ Execute the machine](#️-execute-the-machine)
### 5. [📄 Contributing Guidelines](#-contributing-guidelines)

---

# ✔ Todo
- [X] 8 bit core
- [X] 16 bit core
- [X] Signed integer
- [ ] Float
- [ ] Division circuit and instruction
- [ ] Custom registers
- [ ] Custom clock
- [ ] Custom RAM

---

# 💻 How does a computer work behind the curtains? (WIP)
Developers usually have a good understanding of how a computer works. However, for me at least understanding a computer so deeply that you would be able to actually build one yourself seemed like an impossible task. So, in this section I want to break to you the most important pieces of the computer puzzle that you need to know in order to build your own computer from scratch using only circuits.

## Circuit Components

#### Central Processing Unit (CPU)
The CPU is the brain of the computer, where every instruction is processed. It is composed by the sub-components below:
  - *Control Unit (CU)*: it controls the CPU flags, which are used to direct the execution of the instructions;
  - *Arithmetic Logic Unit (ALU)*: it performs sums, subtractions and comparisons between numbers;
  - *Registers*: they are a small memories in the CPU for auxiliary purposes and faster accessed;
  
  	The Open-Machine's Circuit only has two registers: the accumulator (ACC) and the program counter (PC) registers.

#### Random Access Memory (RAM)
RAM is a temporary memory, which means that when the computer is turned off, all of its data is lost. However, it is very fast and serves, among other things, to store variables during the execution of programs.

#### Disk (not implemented in the circuit yet)
Disk is a permanent memory, which means that it is used to store data that shouldn't be deleted after turning it off. However, it is slower than RAM.

The disk can be an Hard Drive (HD) or Solid-State Drive (SSD).

#### Input/Output Devices (not implemented in the circuit yet)
For an actual computer to work, you also need input devices such as keyboard and mouse, and output devices such as screen.

---

# 🔢 Machine Code

A machine code command takes 16 bits in which first 4 bits represent the instruction and the following 12 bits are the parameter. For example, in the command ```0x1202```, the instruction is ```0x1``` and the parameter is ```0x202```.

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
0x1 | [ACC] = [EE] | A value from the memory is copied to the ACC register | variable | Memory address of a variable that will be used in the instruction
0x2 | [EE] = [ACC] | The value from the ACC register is stored into memory | variable | Memory address of a variable that will be used in the instruction
0x3 | [ACC] = [ACC] + [EE] | The sum of the value of the ACC register and a value from the memory is stored in the ACC register | variable | Memory address of a variable that will be used in the instruction
0x4 | [ACC] = [ACC] - [EE] | The difference between the value of the ACC register and a value from the memory is stored in the ACC register | variable | Memory address of a variable that will be used in the instruction
0x7 | [EE] = input value | The input value is copied to the memory | variable | Memory address of a variable that will be used in the instruction
0x8 | Output [EE] | Outputs a value from the memory into the circuit LEDs | variable | Memory address of a variable that will be used in the instruction
0x9 | Finishes program | When this instruction is encountered, the program is finished and no more instructions will be executed | - | No parameter is required
0xa | Jump to EE | Jump to another line of code | instruction | Memory address of a instruction the program will jump to
0xb | Jump to EE if [ACC] > 0 | Jump to another line of code if the value of the ACC register is positive | instruction | Memory address of a instruction the program will jump to if the condition is right
0xd | Jump to EE if [ACC] = 0 | Jump to another line of code if the value of the ACC register is zero | instruction | Memory address of a instruction the program will jump to if the condition is right
0xf | Jump to EE if [ACC] < 0 | Jump to another line of code if the value of the ACC register is negative | instruction | Memory address of a instruction the program will jump to if the condition is right

## Machine Code Example
```sh
01ff # copy value in the address ff in RAM
020a # stores the value of AC in the address 0a
0900 # kills program
```

## Tips
- Remember to add the ```0x9``` instruction at the end of your programs

---

# ▶️ Execute the machine

#### If you want to see it working
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

#### If you only want to see the outputs (only available for the executables generated by the assembler and the compiler)
*⚠️ Warning: Logisim-Evolution is not that stable when it comes to executing it without the graphics!*
1. Run the machine with your program (replace ```executableFile``` with the executable file name generated by the assembler or the compiler):
	```sh
	java -jar logisim-evolution.jar main.circ -load executableFile -tty table
	```
2. The outputs will appear on the console.
   - Ignore the first output
   - The outputs will follow the pattern: ```{16 bits of the main output}     {4 bit ouptut counter}```

---

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/Circuits/blob/master/CONTRIBUTION.md).
