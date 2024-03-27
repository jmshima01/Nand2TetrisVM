@256
D=A
@SP
M=D

// push constant 0
@0
D=A
@SP
AM=M+1
A=A-1
M=D

// pop local 0
@SP
AM=M-1
D=M
@LCL
A=M
M=D
(LOOP)

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push local 0
@LCL
A=M
D=M
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

// pop local 0
@SP
AM=M-1
D=M
@LCL
A=M
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

// pop argument 0
@SP
AM=M-1
D=M
@ARG
A=M
M=D

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// if-goto LOOP
@SP
AM=M-1
D=M
@LOOP
D;JNE

// push local 0
@LCL
A=M
D=M
@SP
AM=M+1
A=A-1
M=D
