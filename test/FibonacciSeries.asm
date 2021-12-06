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

// ===== push Constant 0 =====
@0
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop That 0 ======
@0
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

// ===== push Constant 1 =====
@1
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== pop That 1 ======
@1
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


// ===== push Constant 2 =====
@2
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

// ===== label MAIN_LOOP_START ======
(MAIN_LOOP_START)

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


// ===== if-goto COMPUTE_ELEMENT ======
// SP--
@SP
M=M-1

A=M
D=M
@COMPUTE_ELEMENT
D;JGT

// ===== goto END_PROGRAM ======
@END_PROGRAM
0;JMP

// ===== label COMPUTE_ELEMENT ======
(COMPUTE_ELEMENT)

// ===== push That 0 =====
@0
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


// ===== push That 1 =====
@1
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

// ===== push Pointer 1 =====
@THAT
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

// ===== goto MAIN_LOOP_START ======
@MAIN_LOOP_START
0;JMP

// ===== label END_PROGRAM ======
(END_PROGRAM)

