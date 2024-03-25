@256
D=A
@SP
M=D

// push constant 3030
@3030
D=A
@SP
AM=M+1
A=A-1
M=D

// pop pointer 0
@SP
AM=M-1
D=M
@THIS
M=D

// push constant 3040
@3040
D=A
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

// push constant 32
@32
D=A
@SP
AM=M+1
A=A-1
M=D

// pop this 2
@SP
AM=M-1
D=M
@THIS
A=M+1
A=A+1
M=D

// push constant 46
@46
D=A
@SP
AM=M+1
A=A-1
M=D

// pop that 6
@SP
AM=M-1
D=M
@THAT
A=M+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
M=D

// push pointer 0
@THIS
D=M
@SP
AM=M+1
A=A-1
M=D

// push pointer 1
@THAT
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

// push this 2
@THIS
A=M+1
A=A+1
D=M
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

// push that 6
@THAT
D=M
@6
A=D+A
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
