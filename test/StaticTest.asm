// ===== push Constant 111 =====
@111
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 333 =====
@333
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 888 =====
@888
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop Static 8 ======
@StaticTest.8
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

// ===== pop Static 3 ======
@StaticTest.3
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

// ===== pop Static 1 ======
@StaticTest.1
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

// ===== push Static 3 =====
@StaticTest.3
D=M
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Static 1 =====
@StaticTest.1
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


// ===== push Static 8 =====
@StaticTest.8
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


