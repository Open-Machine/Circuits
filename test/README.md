# Test

Code: 16 bits = instruction (4 bits) + parameter (12 bits). For example, in the command ```0x1202```, the instruction is ```0x1``` and the parameter is ```0x202```.

### Ram
```
0000 code 9000 0000x10 memory
```

### Memory
50th = 32 = next
```sh
0000 # aux (32)
0101 # a (33)
1234 # b (34)
0000 # zero (35)
0001 # positive (36)
f000 # negative (37)
ffff # max (38)
8888 # end (39)
```

### Code
```sh
{test sum and subtraction: 0x3 and 0x4}
{switch a and b and print a position}

# a + b (1-2)
1033
3034

# print a + b (3-4)
2032
8032
{expect 1335}

# a - b (5-6)
1033
4034
{expect EECD}

---

{test get and update memory: 0x1 and 0x2}
{switch a and b and print a position}

# aux = a (7-8)
1033
2032

# a = b (9-A)
1034
2033

# b = aux (B-C)
1032
2034

# print b (D)
8033
{expect 1234}

---

{test jumps}

# always jump (E-F)
a010
8038
{doesnt expect print}


# jump > 0 for negative (10-12)
1037
b013
8038
{expect print ffff}

# jump > 0 for zero (13-15)
1035
b016
8038
{expect print ffff}

# jump > 0 for positive (16-18)
1036
b019
8038
{doesnt expect print}


# jump = 0 for negative (19-1B)
1037
d01B
8038
{expect print ffff}

# jump = 0 for positive (1C-1E)
1036
d01F
8038
{expect print ffff}

# jump = 0 for zero (1F-21)
1035
d022
8038
{doesnt expect print}


# jump < 0 for zero (22-24)
1035
f025
8038
{expect print ffff}

# jump < 0 for positive (25-27)
1036
f028
8038
{expect print ffff}

# jump < 0 for negative (28-2A)
1037
f02B
8038
{doesnt expect print}

---

# kill (2B-2C)
8039
{expect 8888}
9000

```
