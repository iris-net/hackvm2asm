// ===== push Constant 7 =====
@7
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 8 =====
@8
D=A
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


