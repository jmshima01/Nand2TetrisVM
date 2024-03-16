package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ========== CONSTS/Globals ========================
const POP string = "@SP\nAM=M-1\nD=M\n" 
const PUSH string = "@SP\nAM=M+1\nA=A-1\nM=D\n"
// var STACK []int = make([]int, 0)

//  ============== File IO =======================
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

func writeToFile(path string,asm string){
	err := os.WriteFile(path,[]byte(asm),0644)
	if err != nil{
		println("error reading",path)
		os.Exit(1)
	}
}


// ============ ASM CONVERSIONS ==================

// convert pop to asm  |SET M to D from pop for the given seg offset|
func translatePop(line []string, fileName string)string{
	segment,v:= line[1],line[2]
	offset,_ := strconv.Atoi(v)
	result := fmt.Sprintf("\n// pop %s %d\n",segment,offset)
	
	switch(segment){
		case "static":
			result += POP + fmt.Sprintf("@%s.%d\nM=D\n",fileName,offset)
			
		case "pointer":
			if offset == 0{
				result += POP + "@THIS\nM=D\n"
			} else if offset == 1{
				result += POP + "@THAT\nM=D\n"
			} else{
				println("Syntax Error:",line)
				os.Exit(1)
			}
		case "temp":
			
			switch(offset){
				case 0:
					result+= POP + "@R5\nM=D\n"
				case 1:
					result+= POP + "@R6\nM=D\n"
				case 2:
					result+= POP + "@R7\nM=D\n"
				case 3:
					result+= POP + "@R8\nM=D\n"
				case 4:
					result+= POP + "@R9\nM=D\n"
				case 5:
					result+= POP + "@R10\nM=D\n"
				case 6:
					result+= POP + "@R11\nM=D\n"
				case 7:
					result+= POP + "@R12\nM=D\n"
			default:
				println("Syntax Error:",line)
				os.Exit(1)
			}
			
		case "local":
			if offset == 0{
				result += POP + "@LCL\nA=M\nM=D\n"
			} else if offset == 1{
				result += POP + "@LCL\nA=M+1\nM=D\n"
								
			} else if offset < 7{
				result += POP + "@LCL\nA=M+1\n"
				for i:=0; i<offset-1; i++{
					result += "A=A+1\n"	
				}
				result += "M=D\n"	
					
			} else{
				result+= fmt.Sprintf("@LCL\nD=M\n@%d\nD=D+A\n@R13\nM=D\n%s@R13\nA=M\nM=D\n",offset,POP)
			}

		case "argument":
			if offset == 0{
				result += POP + "@ARG\nA=M\nM=D\n"
			} else if offset == 1{
				result += POP + "@ARG\nA=M+1\nM=D\n"
								
			} else if offset < 7{
				result += POP + "@ARG\nA=M+1\n"
				for i:=0; i<offset-1; i++{
					result += "A=A+1\n"	
				}
				result += "M=D\n"	
					
			} else{
				result+= fmt.Sprintf("@ARG\nD=M\n@%d\nD=D+A\n@R13\nM=D\n%s@R13\nA=M\nM=D\n",offset,POP)
			}
		
		case "this":
			if offset == 0{
				result += POP + "@THIS\nA=M\nM=D\n"
			} else if offset == 1{
				result += POP + "@THIS\nA=M+1\nM=D\n"
								
			} else if offset < 7{
				result += POP + "@THIS\nA=M+1\n"
				for i:=0; i<offset-1; i++{
					result += "A=A+1\n"	
				}
				result += "M=D\n"	
					
			} else{
				result+= fmt.Sprintf("@THIS\nD=M\n@%d\nD=D+A\n@R13\nM=D\n%s@R13\nA=M\nM=D\n",offset,POP)
			}
		
		case "that":
			if offset == 0{
				result += POP + "@THAT\nA=M\nM=D\n"
			} else if offset == 1{
				result += POP + "@THAT\nA=M+1\nM=D\n"
								
			} else if offset < 7{
				result += POP + "@THAT\nA=M+1\n"
				for i:=0; i<offset-1; i++{
					result += "A=A+1\n"	
				}
				result += "M=D\n"	
					
			} else{
				result+= fmt.Sprintf("@THAT\nD=M\n@%d\nD=D+A\n@R13\nM=D\n%s@R13\nA=M\nM=D\n",offset,POP)
			}
		default:
			println("Syntax Error:",line)
			os.Exit(1)

	}
	return result
}

// convert push to asm |SET D to RIGHT VALUE for pushing|!
func translatePush(line []string, fileName string)string{
	segment,v:= line[1],line[2]
	offset,_ := strconv.Atoi(v)
	result := fmt.Sprintf("\n// push %s %d\n",segment,offset)

	switch(segment){
		case "constant":
			result += fmt.Sprintf("@%d\nD=A\n",offset)

		case "static":
			
			result += fmt.Sprintf("@%s.%d\nD=M\n",fileName,offset)
			
		case "pointer":
			if offset == 0{
				result += "@THIS\nD=M\n"
			} else if offset == 1{
				result += "@THAT\nD=M\n"
			} else{
				println("Syntax Error:",line)
				os.Exit(1)
			}
		case "temp":
			switch(offset){
				case 0:
					result+="@R5\nD=M\n"
				case 1:
					result+="@R6\nD=M\n"
				case 2:
					result+="@R7\nD=M\n"
				case 3:
					result+="@R8\nD=M\n"
				case 4:
					result+="@R9\nD=M\n"
				case 5:
					result+="@R10\nD=M\n"
				case 6:
					result+="@R11\nD=M\n"
				case 7:
					result+="@R12\nD=M\n"
				default:
					println("Syntax Error:",line)
					os.Exit(1)
			}
		
		case "local":
			if offset == 0{
				result += "@LCL\nA=M\nD=M\n"
			} else if offset == 1{
				result += "@LCL\nA=M+1\nD=M\n"
								
			} else if offset == 2{
				result += "@LCL\nA=M+1\nA=A+1\nD=M\n"
					
			} else{
				result += fmt.Sprintf("@LCL\nD=M\n@%d\nA=D+A\nD=M\n",offset)
			}

		case "argument":
			if offset == 0{
				result += "@ARG\nA=M\nD=M\n"

			} else if offset == 1{
				result += "@ARG\nA=M+1\nD=M\n"
								
			} else if offset == 2{
				result += "@ARG\nA=M+1\nA=A+1\nD=M\n"
					
			} else{
				result += fmt.Sprintf("@ARG\nD=M\n@%d\nA=D+A\nD=M\n",offset)
			}
		
		case "this":
			if offset == 0{
				result += "@THIS\nD=M\n"
			} else if offset == 1{
				result += "@THIS\nA=M+1\nD=M\n"
								
			} else if offset == 2{
				result += "@THIS\nA=M+1\nA=A+1\nD=M\n"
					
			} else{
				result += fmt.Sprintf("@THIS\nD=M\n@%d\nA=D+A\nD=M\n",offset)
			}
		
		case "that":
			if offset == 0{
				result += "@THAT\nA=M\nD=M\n"
			} else if offset == 1{
				result += "@THAT\nA=M+1\nD=M\n"	
								
			} else if offset == 2{
				result += "@THAT\nA=M+1\nA=A+1\nD=M\n"
					
			} else{
				result += fmt.Sprintf("@THAT\nD=M\n@%d\nA=D+A\nD=M\n",offset)
			}
	

		default:
			println("Syntax Error:",line)
			os.Exit(1)

	}
	result+=PUSH
	return result
}

// convert arithm to asm
func translateArith(line []string)string{
	if len(line) != 1{
		println("Syntax Error:",line)
		os.Exit(1)
	}
	op:= line[0]
	result := fmt.Sprintf("\n// eval %s\n",op)
	switch(op){
		case "add":
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D+M\n" + PUSH
		case "sub":
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D-M\n" + PUSH
		case "and":
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D&M\n" + PUSH
		case "or":
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D|M\n" + PUSH
		case "lt":
		case "gt":

		case "eq":
			result+= POP + "@R13\nM=D\n"
			
		case "not":
			result+=POP +"D=!D\n" + PUSH
		case "neg":
			result+=POP +"D=-D\n" + PUSH
		default:
			println("Syntax Error:",line)
			os.Exit(1)
	}

	return result
}

func main(){

	args := os.Args

	if len(args) != 2{
		println("Usage: ./vm_translator <vmfile>")
		os.Exit(1)
	}

	lines:=readLines(args[1]) // read file by lines into array
	filePath := strings.Split(args[1],".")[0] // ../../foo.vm -> ../../foo
	name := strings.Split(filePath,"/")
	fileName := name[len(name)-1] // ../../foo.vm -> foo (for static vars)
	clean:= make([][]string,0)
	// fmt.Println(lines)
	
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

	assembly,s := "@256\nD=A\n@SP\nM=D\n",""

	// go thru data and convert each line into corresponding hack asm instrs
	for _,v := range clean{
		if v[0] == "push"{
			s=translatePush(v,fileName)
		} else if v[0] == "pop"{
			s=translatePop(v,fileName)
		} else{
			s=translateArith(v)
		}
		assembly+=s
	}

	fmt.Println(assembly)
	writeToFile(fmt.Sprintf("%s.asm",filePath),assembly)
	
}