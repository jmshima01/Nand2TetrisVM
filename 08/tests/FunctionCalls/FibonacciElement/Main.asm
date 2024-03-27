@256
D=A
@SP
M=D

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push constant 2
@2
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
@if.0
D;JLT
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

// if-goto N_LT_2
@SP
AM=M-1
D=M
@N_LT_2
D;JNE

// goto N_GE_2
@N_GE_2
0;JMP
(N_LT_2)

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D
(N_GE_2)

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push constant 2
@2
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

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push constant 1
@1
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
