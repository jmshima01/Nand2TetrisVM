@256
D=A
@SP
M=D

// push constant 10
@10
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

// push constant 21
@21
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 22
@22
D=A
@SP
AM=M+1
A=A-1
M=D

// pop argument 2
@SP
AM=M-1
D=M
@ARG
A=M+1
A=A+1
M=D

// pop argument 1
@SP
AM=M-1
D=M
@ARG
A=M+1
M=D

// push constant 36
@36
D=A
@SP
AM=M+1
A=A-1
M=D

// pop this 6
@SP
AM=M-1
D=M
@THIS
A=M+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
M=D

// push constant 42
@42
D=A
@SP
AM=M+1
A=A-1
M=D

// push constant 45
@45
D=A
@SP
AM=M+1
A=A-1
M=D

// pop that 5
@SP
AM=M-1
D=M
@THAT
A=M+1
A=A+1
A=A+1
A=A+1
A=A+1
M=D

// pop that 2
@SP
AM=M-1
D=M
@THAT
A=M+1
A=A+1
M=D

// push constant 510
@510
D=A
@SP
AM=M+1
A=A-1
M=D

// pop temp 6
@SP
AM=M-1
D=M
@R11
M=D

// push local 0
@LCL
A=M
D=M
@SP
AM=M+1
A=A-1
M=D

// push that 5
@THAT
D=M
@5
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

// push argument 1
@ARG
A=M+1
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

// push this 6
@THIS
D=M
@6
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D

// push this 6
@THIS
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

// push temp 6
@R11
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
