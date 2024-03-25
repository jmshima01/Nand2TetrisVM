@256
D=A
@SP
M=D

// push constant 17
@17
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 17
@17
D=A
@SP
AM=M+1
A=A-1
M=D

// eval eq
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.0
D;JEQ
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.0
0;JMP
(if.0)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.0)

// push constant 17
@17
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 16
@16
D=A
@SP
AM=M+1
A=A-1
M=D

// eval eq
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.1
D;JEQ
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.1
0;JMP
(if.1)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.1)

// push constant 16
@16
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 17
@17
D=A
@SP
AM=M+1
A=A-1
M=D

// eval eq
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.2
D;JEQ
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.2
0;JMP
(if.2)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.2)

// push constant 892
@892
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 891
@891
D=A
@SP
AM=M+1
A=A-1
M=D

// eval lt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.3
D;JLT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.3
0;JMP
(if.3)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.3)

// push constant 891
@891
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 892
@892
D=A
@SP
AM=M+1
A=A-1
M=D

// eval lt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.4
D;JLT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.4
0;JMP
(if.4)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.4)

// push constant 891
@891
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 891
@891
D=A
@SP
AM=M+1
A=A-1
M=D

// eval lt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.5
D;JLT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.5
0;JMP
(if.5)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.5)

// push constant 32767
@32767
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 32766
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

// eval gt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.6
D;JGT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.6
0;JMP
(if.6)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.6)

// push constant 32766
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 32767
@32767
D=A
@SP
AM=M+1
A=A-1
M=D

// eval gt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.7
D;JGT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.7
0;JMP
(if.7)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.7)

// push constant 32766
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 32766
@32766
D=A
@SP
AM=M+1
A=A-1
M=D

// eval gt
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@if.8
D;JGT
// false block
D=0
@SP
AM=M+1
A=A-1
M=D
@fi.8
0;JMP
(if.8)
D=-1
@SP
AM=M+1
A=A-1
M=D
(fi.8)

// push constant 57
@57
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 31
@31
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 53
@53
D=A
@SP
AM=M+1
A=A-1
M=D

// eval add
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D+M
@SP
AM=M+1
A=A-1
M=D

// push constant 112
@112
D=A
@SP
AM=M+1
A=A-1
M=D

// eval sub
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D-M
@SP
AM=M+1
A=A-1
M=D

// eval neg
@SP
AM=M-1
D=M
D=-D
@SP
AM=M+1
A=A-1
M=D

// eval and
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D&M
@SP
AM=M+1
A=A-1
M=D

// push constant 82
@82
D=A
@SP
AM=M+1
A=A-1
M=D

// eval or
@SP
AM=M-1
D=M
@R13
M=D
@SP
AM=M-1
D=M
@R13
D=D|M
@SP
AM=M+1
A=A-1
M=D

// eval not
@SP
AM=M-1
D=M
D=!D
@SP
AM=M+1
A=A-1
M=D
