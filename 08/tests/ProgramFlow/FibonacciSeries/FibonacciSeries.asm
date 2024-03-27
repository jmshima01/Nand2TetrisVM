@256
D=A
@SP
M=D

// push argument 1
@ARG
A=M+1
D=M
@SP
AM=M+1
A=A-1
M=D

// pop pointer 1
@SP
AM=M-1
D=M
@THAT
M=D

// push constant 0
@0
D=A
@SP
AM=M+1
A=A-1
M=D

// pop that 0
@SP
AM=M-1
D=M
@THAT
A=M
M=D

// push constant 1
@1
D=A
@SP
AM=M+1
A=A-1
M=D

// pop that 1
@SP
AM=M-1
D=M
@THAT
A=M+1
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
(LOOP)

// push argument 0
@ARG
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// if-goto COMPUTE_ELEMENT
@SP
AM=M-1
D=M
@COMPUTE_ELEMENT
D;JNE

// goto END
@END
0;JMP
(COMPUTE_ELEMENT)

// push that 0
@THAT
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push that 1
@THAT
A=M+1
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

// pop that 2
@SP
AM=M-1
D=M
@THAT
A=M+1
A=A+1
M=D

// push pointer 1
@THAT
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

// pop pointer 1
@SP
AM=M-1
D=M
@THAT
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

// goto LOOP
@LOOP
0;JMP
(END)
