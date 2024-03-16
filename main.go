package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "strconv"
)

// CONSTS/STATICS
const POP string = "@SP\nAM=M-1\nD=M\n"
const PUSH string = "@SP\nAM=M+1\nA=A-1\nM=D\n"
var fileName string= "" // for temp vars

// file reader helper
func readLines(path string)[]string{
	f,err := os.ReadFile(path)
	if err!=nil{
		println("error reading",path)
		os.Exit(1)
	}
	data:=string(f)
	data = strings.ReplaceAll(data,"\r\n","\n")
	return strings.Split(strings.Trim(data,"\n"),"\n")
	
}

// convert pop to asm
func translatePop(line []string)string{
	// segment,offset := line[1],line[2]
	segment,offset:= line[1],line[2]
	result := fmt.Sprintf("\n// pop %s %s\n%s",segment,offset,POP)
	
	switch(segment){
		case "static":
		case "pointer":
		case "temp":
		case "local":
		case "argument":
		case "this":
		case "that":
		default:
			println("Syntax Error:",line)
			os.Exit(1)


	}
	return result
}

// convert push to asm
func translatePush(line []string)string{
	segment,offset:= line[1],line[2]
	result := fmt.Sprintf("\n// push %s %s\n",segment,offset)
	intOffset,_ := strconv.Atoi(offset)

	switch(segment){
		case "constant":
			result += fmt.Sprintf("@%s\nD=A\n",offset)

		case "static":
			result += fmt.Sprintf("@%s\n",segment)
			
		case "pointer":
			if offset == "0"{
				result += "@THIS\nD=M\n"
			} else if offset == "1"{
				result += "@THAT\nD=M\n"
			} else{
				println("Syntax Error:",line)
				os.Exit(1)
			}
		case "temp":
			intOffset %= 8
			result += fmt.Sprintf("@R5\nD=A\n@%d\nA=D+A\nD=M\n",intOffset)
		case "local":
			if offset == "0"{
				result += "@LCL\nD=M\n"
			} else{
				result += "@LCL\nD=M\n"
				result += fmt.Sprintf("@%s\nA=D+A\nD=M\n",offset)
			}

		case "argument":
			if offset == "0"{
				result += "@ARG\nD=M\n"
			} else{
				result += "@ARG\nD=M\n"
				result += fmt.Sprintf("@%s\nA=D+A\nD=M\n",offset)
			}
		case "this":
			if offset == "0"{
				result += "@THIS\nD=M\n"
			} else{
				result += "@THIS\nD=M\n"
				result += fmt.Sprintf("@%s\nA=D+A\nD=M\n",offset)
			}

		case "that":
			if offset == "0"{
				result += "@THAT\nD=M\n"
			} else{
				result += "@THAT\nD=M\n"
				result += fmt.Sprintf("@%s\nA=D+A\nD=M\n",offset)
			}

		default:
			println("Syntax Error:",line)
			os.Exit(1)

	}
	result+=PUSH
	return result
}

// convert arithm to asm
func translateBinaryArith(line []string)string{
	return ""
}

// convert other unary arithm to asm i.e not neg
func translateUnaryArith(line []string)string{
	return ""
}

func main(){

	args := os.Args

	if len(args) != 2{
		println("Usage: ./vm_translator <vmfile>")
		os.Exit(1)
	}

	lines:=readLines(args[1]) // read file by lines into array
	fileName = strings.Split(args[1],".")[0] // foo.vm -> foo
	clean:= make([][]string,0)
	fmt.Println(lines)
	
	// remove whitespace and comments un-needed
	for _,v := range lines{
		if len(v)==0{
			continue
		}
		remove,_ :=regexp.MatchString("//.*",v)
		
		if !remove{
			clean = append(clean, strings.Fields(v))
		}
	}
	fmt.Println()
	fmt.Println(clean)

	assembly,s:= "",""

	// go thru data and convert each line into corresponding  hack asm instrs
	for _,v:=range clean{
		if v[0] == "push"{
			s=translatePush(v)
			fmt.Println(s)
		} else if v[0] == "pop"{
			s=translatePop(v)
		} else{
			switch(v[0]){
				case "add":
					s=translateBinaryArith(v)
				case "sub":
					s=translateBinaryArith(v)
				case "and":
					s=translateBinaryArith(v)
				case "or":
					s=translateBinaryArith(v)
				case "lt":
					s=translateBinaryArith(v)
				case "gt":
					s=translateBinaryArith(v)
				case "eq":
					s=translateBinaryArith(v)
				case "not":
					s=translateUnaryArith(v)
				case "neg":
					s=translateUnaryArith(v)
				case "default":
					println("Syntax error:",v)
					os.Exit(1)
			}
		}
		assembly+=s
	}
}