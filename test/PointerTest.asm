// ===== push Constant 3030 =====
@3030
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Pointer 0 ======
@THIS
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

// ===== push Constant 3040 =====
@3040
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Pointer 1 ======
@THAT
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

// ===== push Constant 32 =====
@32
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop This 2 ======
@2
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

// ===== push Constant 46 =====
@46
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop That 6 ======
@6
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

// ===== push Pointer 0 =====
@THIS
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Pointer 1 =====
@THAT
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


// ===== push This 2 =====
@2
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


// ===== push That 6 =====
@6
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


