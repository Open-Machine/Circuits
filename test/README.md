# Test

## Run Test
```sh
go test
```

---

## Code Explanation
Code: 16 bits = instruction (4 bits) + parameter (12 bits). For example, in the command ```0x1202```, the instruction is ```0x1``` and the parameter is ```0x202```.

### Ram
```
code 9000 0000x10 memory
```

### Memory
50th = 0x32 = next
```sh
0000 # aux (0x32)
0101 # a (0x33)
1234 # b (0x34)
0000 # zero (0x35)
0001 # positive (0x36)
f000 # negative (0x37)
ffff # max (0x38)
8888 # end (0x39)
```

### Code
```sh

{test execution of first command}

8038
{expect print 0xffff=65535}

---

{test sum and subtraction: 0x3 and 0x4}
{switch a and b and print a position}

# a + b (1-2)
1033
3034

# print a + b (3-4)
2032
8032
{expect print 0x1335=4917}

# a - b (5-6)
1033
4034

# print a - b (7-8)
2032
8032
{expect print 0xEECD=61133}

---

{test get and update memory: 0x1 and 0x2}
{switch a and b and print a position}

# aux = a (9-A)
1033
2032

# a = b (B-C)
1034
2033

# b = aux (D-E)
1032
2034

# print b (F)
8033
{expect print 0x1234=4660}

---

{test jumps}

# always jump (10-11)
a012
8038
{doesnt expect print}


# jump > 0 for negative (12-14)
1037
b015
8037
{expect print 0xf000=61440}

# jump > 0 for zero (15-17)
1035
b018
8035
{expect print 0x0000=0}

# jump > 0 for positive (18-1A)
1036
b01B
8038
{doesnt expect print}


# jump = 0 for negative (1B-1D)
1037
d01E
8037
{expect print 0xf000=61440}

# jump = 0 for positive (1E-20)
1036
d021
8036
{expect print 0x0001=1}

# jump = 0 for zero (21-23)
1035
d024
8038
{doesnt expect print}


# jump < 0 for zero (24-26)
1035
f027
8035
{expect print 0x0000=0}

# jump < 0 for positive (27-29)
1036
f02A
8036
{expect print 0x0001=1}

# jump < 0 for negative (2A-2C)
1037
f02D
8038
{doesnt expect print}

---

# kill (2D-2E)
8039
{expect print 0x8888=34952}
9000

```
