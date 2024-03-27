@256
D=A
@SP
M=D

// func defn
(SimpleFunction.test)
@0
D=A
@SP
AM=M+1
A=A-1
M=D
@0
D=A
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

// push local 1
@LCL
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

// eval not
@SP
AM=M-1
D=M
D=!D
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
