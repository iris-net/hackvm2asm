// ===== push Constant 0 =====
@0
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

// ===== label LOOP_START ======
(LOOP_START)

// ===== push Argument 0 =====
@0
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

// ===== push Argument 0 =====
@0
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


// ===== push Constant 1 =====
@1
D=A
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


// ===== pop Argument 0 ======
@0
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

// ===== push Argument 0 =====
@0
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


// ===== if-goto LOOP_START ======
// SP--
@SP
M=M-1

A=M
D=M
@LOOP_START
D;JGT

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


