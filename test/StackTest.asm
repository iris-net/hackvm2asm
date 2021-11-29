// ===== push Constant 17 =====
@17
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 17 =====
@17
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command EQ =====
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
// EQ
D=M-D
@EQTRUE0
D;JEQ
@SP
A=M
M=0
@EQEND0
0;JMP
(EQTRUE0)
@SP
A=M
M=-1
(EQEND0)
// SP++
@SP
M=M+1


// ===== push Constant 17 =====
@17
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 16 =====
@16
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command EQ =====
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
// EQ
D=M-D
@EQTRUE1
D;JEQ
@SP
A=M
M=0
@EQEND1
0;JMP
(EQTRUE1)
@SP
A=M
M=-1
(EQEND1)
// SP++
@SP
M=M+1


// ===== push Constant 16 =====
@16
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 17 =====
@17
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command EQ =====
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
// EQ
D=M-D
@EQTRUE2
D;JEQ
@SP
A=M
M=0
@EQEND2
0;JMP
(EQTRUE2)
@SP
A=M
M=-1
(EQEND2)
// SP++
@SP
M=M+1


// ===== push Constant 892 =====
@892
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 891 =====
@891
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command LT =====
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
// LT
D=M-D
@LTTRUE3
D;JLT
@SP
A=M
M=0
@LTEND3
0;JMP
(LTTRUE3)
@SP
A=M
M=-1
(LTEND3)
// SP++
@SP
M=M+1


// ===== push Constant 891 =====
@891
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 892 =====
@892
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command LT =====
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
// LT
D=M-D
@LTTRUE4
D;JLT
@SP
A=M
M=0
@LTEND4
0;JMP
(LTTRUE4)
@SP
A=M
M=-1
(LTEND4)
// SP++
@SP
M=M+1


// ===== push Constant 891 =====
@891
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 891 =====
@891
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command LT =====
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
// LT
D=M-D
@LTTRUE5
D;JLT
@SP
A=M
M=0
@LTEND5
0;JMP
(LTTRUE5)
@SP
A=M
M=-1
(LTEND5)
// SP++
@SP
M=M+1


// ===== push Constant 32767 =====
@32767
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 32766 =====
@32766
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command GT =====
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
// GT
D=M-D
@GTTRUE6
D;JGT
@SP
A=M
M=0
@GTEND6
0;JMP
(GTTRUE6)
@SP
A=M
M=-1
(GTEND6)
// SP++
@SP
M=M+1


// ===== push Constant 32766 =====
@32766
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 32767 =====
@32767
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command GT =====
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
// GT
D=M-D
@GTTRUE7
D;JGT
@SP
A=M
M=0
@GTEND7
0;JMP
(GTTRUE7)
@SP
A=M
M=-1
(GTEND7)
// SP++
@SP
M=M+1


// ===== push Constant 32766 =====
@32766
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 32766 =====
@32766
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command GT =====
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
// GT
D=M-D
@GTTRUE8
D;JGT
@SP
A=M
M=0
@GTEND8
0;JMP
(GTTRUE8)
@SP
A=M
M=-1
(GTEND8)
// SP++
@SP
M=M+1


// ===== push Constant 57 =====
@57
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 31 =====
@31
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== push Constant 53 =====
@53
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


// ===== push Constant 112 =====
@112
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


// ===== arithmetic command Neg =====
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Neg
D=0
M=D-M
// SP++
@SP
M=M+1


// ===== arithmetic command And =====
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
// And
M=M&D
// SP++
@SP
M=M+1


// ===== push Constant 82 =====
@82
D=A
@SP
A=M
M=D
// SP++
@SP
M=M+1


// ===== arithmetic command Or =====
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
// Or
M=M|D
// SP++
@SP
M=M+1


// ===== arithmetic command Not =====
// SP--
@SP
M=M-1

// A=Memory[SP]
A=M
// Not
M=!M
// SP++
@SP
M=M+1


