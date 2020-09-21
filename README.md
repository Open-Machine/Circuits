<div align="center"> 
<h1>Open Machine's Circuits</h1>
<h4>Circuits of a basic computer.</h4>

<i>This repository is a component of a larger project: <b><a href="https://github.com/Open-Machine/README">Open-Machine</a></b> - an open-source computer developed from scratch.</i>

<a href="https://github.com/Open-Machine/Circuits/stargazers"><img src="https://img.shields.io/github/stars/Open-Machine/Circuits" alt="Stars Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/network/members"><img src="https://img.shields.io/github/forks/Open-Machine/Circuits" alt="Forks Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/"><img src="https://img.shields.io/badge/version-0.0.1-blue" alt="Version Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/commits/"><img src="https://img.shields.io/github/commit-activity/m/Open-Machine/Circuits" alt="commits"/></a>
<a href="https://github.com/Open-Machine/Circuits/pulls"><img src="https://img.shields.io/github/issues-pr/Open-Machine/Circuits" alt="Pull Requests Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/issues"><img src="https://img.shields.io/github/issues/Open-Machine/Circuits" alt="Issues Badge"/></a>
<a href="https://github.com/Open-Machine/Circuits/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Open-Machine/Circuits?color=2b9348"></a>
<a href="https://github.com/Open-Machine/Circuits/blob/master/LICENSE"><img src="https://img.shields.io/github/license/Open-Machine/Circuits?color=2b9348" alt="License Badge"/></a>

<img src="https://raw.githubusercontent.com/Open-Machine/README/stable/Media/logo-horizontal.png" alt="open-machine"/>

<br/>

</div>

<br/>

---

<br/>

<!-- omit in toc -->
# 🔖 Table of Contents
### 1. [💻 How does a computer work behind the curtains?](#-how-does-a-computer-work-behind-the-curtains)
### 2. [🔢 Machine Code](#-machine-code)
### 3. [▶️ Run](#️-run)
### 4. [📄 Contributing Guidelines](#-contributing-guidelines)

<br/>

---

<br/>

# 💻 How does a computer work behind the curtains?

Of course, I don't know and won't tell you everything about how a computer works. However, let's take a closer look to the its components and main functions so you know at least what you are looking at when running the circuit simulation.

## Circuit Components

#### Central Processing Unit (CPU)
The CPU is the brain of the computer, where every instruction is processed. It is composed by the sub-components below:
  - *Control Unit (CU)*: it controls the CPU flags, which are used to direct the execution of the instructions;
  - *Arithmetic Logic Unit (ALU)*: it performs sums, subtractions and comparisons between numbers;
  - *Registers*: they are a small memories in the CPU for auxiliary purposes and faster accessed;
  
  	The Open-Machine's Circuit only has two registers: the accumulator (ACC) and the program counter (PC) registers.

#### Random Access Memory (RAM)
RAM is a temporary memory, which means that when the computer is turned off, all of its data is lost. However, it is very fast and serves, among other things, to store variables during the execution of programs.

#### Disk
Disk is a permanent memory, which means that it is used to store data that shouldn't be deleted after turning it off. However, it is slower than RAM. The disk can be an Hard Drive (HD) or Solid-State Drive (SSD).

#### Input/Output Devices
For an actual computer to work, you also need input devices such as keyboard and mouse, and output devices such as screen.

<br/>

*ps: Since the Open-Computer's is a basic computer, complex input and output devices and disk were chosen not to be implemented. However, the plan is to build it on the future.*

<br/>

---

<br/>

# 🔢 Machine Code

A machine code command takes 16 bits in which first 4 bits represent the instruction and the following 12 bits are the parameter. For example, in the command ```0x1202```, the instruction is ```0x1``` and the parameter is ```0x202```.

## Instructions Table
Let's look at all of the instructions at our disposal.

### Symbols Legend for the Instructions Table
Symbol | Explanation
--- | ---
ACC | The ACC register
EE | Represents a memory index
[ ] | "Value of"

### Instructions Table
Machine Code | Short Instruction Description | Long Instruction Description | Short Param Description | Long Param Description
--- | --- | --- | --- | ---
0x0 | - | This instruction **doesn't perform any action** | - | No parameter is required
0x1 | [ACC] = [EE] | A value from the memory is **copied** to the **ACC** register | variable | Memory address of a variable that will be used in the instruction
0x2 | [EE] = [ACC] | The value from the ACC register is **stored** into **memory** | variable | Memory address of a variable that will be used in the instruction
0x3 | [ACC] = [ACC] + [EE] | The **sum** of the value of the **ACC** register and a value from the memory is stored in the ACC register | variable | Memory address of a variable that will be used in the instruction
0x4 | [ACC] = [ACC] - [EE] | The **difference** between the value of the ACC register and a value from the memory is stored in the ACC register | variable | Memory address of a variable that will be used in the instruction
0x7 | [EE] = input value | The **input** value is copied to the **memory** | variable | Memory address of a variable that will be used in the instruction
0x8 | Output [EE] | Outputs a value from the memory into the circuit LEDs | variable | Memory address of a variable that will be used in the instruction
0x9 | Finishes program | When this instruction is encountered, the **program is finished** and no more instructions will be executed | - | No parameter is required
0xa | Jump to EE | **Jump** to another line of code | instruction | Memory address of a instruction the program will jump to
0xb | Jump to EE if [ACC] > 0 | **Jump** to another line of code **if** the value of the ACC register is **positive** | instruction | Memory address of a instruction the program will jump to if the condition is right
0xd | Jump to EE if [ACC] = 0 | **Jump** to another line of code **if** the value of the ACC register is **zero** | instruction | Memory address of a instruction the program will jump to if the condition is right
0xf | Jump to EE if [ACC] < 0 | **Jump** to another line of code **if** the value of the ACC register is **negative** | instruction | Memory address of a instruction the program will jump to if the condition is right

<br/>

---

<br/>

# 🔀 Code Flow
This section will help you think more in an assembly way.

Because Open-Machine's Circuit only has very simple commands and very few registers, the way to think about your assembly code has to be very different.

### Storage
The circuit has two components that store data: the ACC register and the memory RAM. Both of these are volatile memories, which means that when the circuit is turned the data lost. Let's take a closer look in the memories available and when they should be used:
- **RAM**: is the main memory, it can store thousands of bits and every variable should be stored here.
- **ACC register**: is an auxiliary memory for arithmetic operations that can store only one value. It must not store variables indefinitely. 

	Most CPUs have many registers, so in those cases some registers can be used to store variables indefinitely. However, since Open-Computer's circuit only offers one register, it must be used exclusively as an auxiliary memory for the instructions.

### Operation Flow
Since the circuit has only one register, the flow of the operations will be a little bit different, following a pattern somewhat similar to:

1. Change the value of ACC register
2. Do an instruction
3. Store the value of the ACC register in RAM

**For example**, if you want to sum variables A and B and store the result in C you could use the following instructions:
1. Copy the value of variable A to the ACC register
2. Use the sum instruction to sum the value of ACC register with B and store the result in the ACC register
3. Store the value of the ACC register in C memory address

### IFs, WHILEs, FORs and procedures
If that's your first time programming assembly, it must be very strange to know that there are no ```if```s, ```while```s and ```for```s. However, it's not that hard not having those keywords, because all of those things can be done with the combination arithmetic operations and conditional and unconditional jumps.

Let me show you an example. Imagine you have this code written in C and wanted to translate it to assembly.
```c
	// before
	if (a > b) {
		// ...
	}
	// after
```
One way of doing it would be:
1. **Copy** the value of ```a``` from RAM to the ACC register
2. Update the value of the ACC register with the result of the ```subtraction``` between the value of the ACC register and ```b```
3. **Jump** to **step 5** if the ACC register is greater than zero
4. One or more instructions inside the ```if``` statement
	
	```c
	// ...
	```

5. After instructions

	```c
	// after
	```

### More tips
- Remember to add the **kill** instruction at the **end of your program** to kill the execution

<br/>

---

<br/>

# Machine Code Example
The following assembly code gets two numbers from input and outputs the sum of them. If the sum is greater than zero it will output zero.

*ps: Remember to change the input before starting the clock simulation, because the ```input``` instruction doesn't wait for anything to happen to get the input data.*

```sh
# data inputs
7055
7056

# sum
1055
3056
2057

# output
8057

# if output higher than zero, it will output zero
1057
d00b # if=0
f00b # if<0
80ff # [0xff] is zero since we didn't change it

9000
```

<br/>

---

<br/>

# ▶️ Run
In this section, you will see how to execute the circuit in the GUI.

You can watch [this video](https://www.youtube.com/watch?v=NAITQqdOw7c) as an introduction to Logisim-Evolution, which is the program we will be using to simulate the circuit.

### i. Start the circuit
1. Run the machine:
	```sh
	java -jar logisim-evolution.jar
	```
2. Import the circuit file by navigating the menu:
   *```File -> Open -> Select main.circ from the repository folder```*
3. Open Main file on the left side of Logisim

### ii. Program the circuit
1. Run the circuit. Follow [these](#) steps
2. Paste the executable code into the beginning of the RAM. You may want to change some values of the memory as if you were initializing variables.

### iii. Run the circuit
3. Run the Program by navigating the menu:
   *```Simulate -> Enable 'Ticks Enabled'```*
   - You can change the speed of the program by navigating the menu: 
	*```Simulate -> Tick Frequency -> To get the fastest execution, select the top item```*

<br/>

---

<br/>

# 📄 Contributing Guidelines
Check out the contributing guidelines [here](https://github.com/Open-Machine/Circuits/blob/master/CONTRIBUTING.md).
