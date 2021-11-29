// ===== push Constant 10 =====
@10
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Local 0 ======
@0
D=A
@LCL
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== push Constant 21 =====
@21
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 22 =====
@22
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Argument 2 ======
@2
D=A
@ARG
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== pop Argument 1 ======
@1
D=A
@ARG
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== push Constant 36 =====
@36
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop This 6 ======
@6
D=A
@THIS
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== push Constant 42 =====
@42
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 45 =====
@45
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop That 5 ======
@5
D=A
@THAT
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== pop That 2 ======
@2
D=A
@THAT
D=M+D
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== push Constant 510 =====
@510
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Temp 6 ======
@R11
D=A
@R13
M=D
// SP--
@SP
M=M-1

A=M
D=M
@R13
A=M
M=D

// ===== push Local 0 =====
@0
D=A
@LCL
A=M+D
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push That 5 =====
@5
D=A
@THAT
A=M+D
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command Add =====
// SP--
@SP
M=M-1

// D=Memory[SP]
@SP
A=M
D=M
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Add
M=M+D
// SP++
@SP
M=M+1


// ===== push Argument 1 =====
@1
D=A
@ARG
A=M+D
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command Sub =====
// SP--
@SP
M=M-1

// D=Memory[SP]
@SP
A=M
D=M
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Sub
M=M-D
// SP++
@SP
M=M+1


// ===== push This 6 =====
@6
D=A
@THIS
A=M+D
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push This 6 =====
@6
D=A
@THIS
A=M+D
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command Add =====
// SP--
@SP
M=M-1

// D=Memory[SP]
@SP
A=M
D=M
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Add
M=M+D
// SP++
@SP
M=M+1


// ===== arithmetic command Sub =====
// SP--
@SP
M=M-1

// D=Memory[SP]
@SP
A=M
D=M
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Sub
M=M-D
// SP++
@SP
M=M+1


// ===== push Temp 6 =====
@R11
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command Add =====
// SP--
@SP
M=M-1

// D=Memory[SP]
@SP
A=M
D=M
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Add
M=M+D
// SP++
@SP
M=M+1


