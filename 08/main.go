package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"regexp"
)

// -------------------
// Nand2TetrisVM
// takes in single .vm file path or dir containing .vm file(s)
// and produces single hack .asm translation
// @author James Shima
// -------------------


// ========== CONSTS/Globals ========================
const POP string = "@SP\nAM=M-1\nD=M\n" 
const PUSH string = "@SP\nAM=M+1\nA=A-1\nM=D\n"
var ifCounter int = 0
var returnCounter = 0
var currFunc string = ""

//  ============== File IO =======================
func readLines(path string)[]string{
	f,err := os.ReadFile(path)
	if err!=nil{
		panic(err)
	}
	data:=string(f)
	data = strings.ReplaceAll(data,"\r\n","\n")
	return strings.Split(strings.Trim(data,"\n"),"\n")
	
}

func writeToFile(path string,asm string){
	err := os.WriteFile(path,[]byte(asm),0644)
	if err != nil{
		panic(err)
	}
}

func isVMFile(f os.FileInfo) bool{
	if f.IsDir() {return false}
	ext := strings.Split(f.Name(),".")
	if len(ext) == 1 {return false}
	if ext[1] == "vm" {return true} else {return false}
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
				println("Syntax Error pointer POP:",line)
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
				println("Syntax Error temp POP:",line)
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
			println("Syntax Error POP:",line)
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
			// TODO: fix this for multiple files
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
			println("Syntax Error PUSH:",line)
			os.Exit(1)

	}
	result+=PUSH
	return result
}

// convert hack arithm to asm
func translateArith(line []string)string{
	if len(line) != 1{
		fmt.Println("Syntax Error Arith to Long:",line)
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
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D-M\n@"+fmt.Sprintf("if.%d\n",ifCounter) + "D;JLT\n"
			result+= "// false block\n"	+ "D=0\n" + PUSH + fmt.Sprintf("@fi.%d\n",ifCounter) + "0;JMP\n"
			result+= fmt.Sprintf("(if.%d)\n",ifCounter) + "D=-1\n" + PUSH + fmt.Sprintf("(fi.%d)\n",ifCounter)
			ifCounter++

		case "gt":
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D-M\n@"+fmt.Sprintf("if.%d\n",ifCounter) + "D;JGT\n"
			result+= "// false block\n"	+ "D=0\n" + PUSH + fmt.Sprintf("@fi.%d\n",ifCounter) + "0;JMP\n"
			result+= fmt.Sprintf("(if.%d)\n",ifCounter) + "D=-1\n" + PUSH + fmt.Sprintf("(fi.%d)\n",ifCounter)
			ifCounter++

		case "eq": 
			result+=POP + "@R13\nM=D\n" + POP + "@R13\nD=D-M\n@"+fmt.Sprintf("if.%d\n",ifCounter) + "D;JEQ\n"
			result+= "// false block\n"	+ "D=0\n" + PUSH + fmt.Sprintf("@fi.%d\n",ifCounter) + "0;JMP\n"
			result+= fmt.Sprintf("(if.%d)\n",ifCounter) + "D=-1\n" + PUSH + fmt.Sprintf("(fi.%d)\n",ifCounter)
			ifCounter++
		case "not":
			result+=POP +"D=!D\n" + PUSH
		case "neg":
			result+=POP +"D=-D\n" + PUSH
		default:
			println("Syntax Error Arith:",line)
			os.Exit(1)
	}
	return result
}

// =========== Control Flow ================== 
//*if !0 jumpto label else next instr
func translateIfGoto(line []string) string{
	result := "\n// if-goto " + line[1] + "\n"
	result += POP + fmt.Sprintf("@%s\nD;JNE\n",line[1])
	return result
}

// simple force jump to label in asm
func translateGoto(line []string)string{
	return fmt.Sprintf("\n// goto %s\n@%s\n0;JMP\n",line[1],line[1])
}

// makes basic (Label) in asm
func translateLabel(line []string)string{
	return fmt.Sprintf("(%s)\n",line[1])
}

// ============ Functions =================
func translateFunctionHeader(line []string)string{
	result:= "\n// func defn\n"
	result += translateLabel(line)
	localN,_ := strconv.Atoi(line[2])
	for i:=0; i<localN; i++{
		result+= "@0\nD=A\n"+PUSH
	}
	return result
}

func translateReturn()string{
	result:= "// RETURN\n"
	result += "@5\nD=A\n@LCL\nA=M-D\nD=M\n@R13\nM=D\n@SP\nAM=M-1\nD=M\n" // ret addr
	result += "@ARG\nA=M\nM=D\n" // pop *ARG
	result += "D=A\n@SP\nM=D+1\n" // SP = ARG+1

	// get caller frame
	result += "@LCL\nD=M\n@R14\nAM=D-1\nD=M\n@THAT\nM=D\n"
	result += "@R14\nAM=M-1\nD=M\n@THIS\nM=D\n"
	result += "@R14\nAM=M-1\nD=M\n@ARG\nM=D\n"
	result += "@R14\nAM=M-1\nD=M\n@LCL\nM=D\n"
	
	result += "@R13\nA=M\n0;JMP\n" // goto ret addr
	return result
}

func translateCall(line []string)string{
	result := "\n// func call\n"
	argNum,_ := strconv.Atoi(line[2])
	result += fmt.Sprintf("@RET_ADDRESS_CALL_%d\nD=A\n",returnCounter) + PUSH
	result += "@LCL\nD=M\n" + PUSH + "@ARG\nD=M\n" + PUSH + "@THIS\nD=M\n" + PUSH + "@THAT\nD=M\n" + PUSH
	
	argNum+=5

	result += "@SP\nD=M\n" + fmt.Sprintf("@%d\nD=D-A\n@ARG\nM=D\n@SP\nD=M\n@LCL\nM=D\n",argNum)

	result += fmt.Sprintf("@%s\n0;JMP\n",line[1])
	result += fmt.Sprintf("(RET_ADDRESS_CALL_%d)\n",returnCounter)
	
	returnCounter++
	
	return result
}

func bootstrap(filename string)string{
	result := "// " + filename + ".asm\n"	
	result += "@256\nD=A\n@SP\nM=D\n"
	inp := []string{"call","Sys.init","0"}
	result += translateCall(inp)
	return result
}


func handleSingleFile(filepth string)([]string,string,string){
	lines := readLines(filepth)
	filePath := strings.Split(filepth,".")[0] // ../../foo.vm -> ../../foo
	name := strings.Split(filePath,"/")
	fileName := name[len(name)-1] // ../../foo.vm -> foo (for static vars)
	
	return lines,filePath,fileName
}

func handleDir(dirPath string)([][]string,string,[]string){
	dir, err := os.ReadDir(dirPath)
	if err != nil{
		panic(err)
	}

	// fmt.Println(dir)
	noVMfiles := true
	lines := make([][]string,0)
	fileNames := make([]string,0)
	for _,v := range dir{
		info,_ := v.Info()
		if isVMFile(info){
			l,_,_ := handleSingleFile(fmt.Sprintf("%s/%s",dirPath,v.Name()))
			fileNames = append(fileNames,strings.Split(v.Name(),".")[0])
			lines = append(lines,l)
			fmt.Println(v.Name())
			noVMfiles = false
		}
	}
	if noVMfiles{
		println("no vm files found in",dir)
	}
	// fmt.Println(dirPath)
	name := strings.Split(dirPath,"/")
	finalName := name[len(name)-1]
	finalPath := dirPath+"/"+finalName
	return lines,finalPath,fileNames

}


func assemble(lines []string, fileName string)string{
	clean:= make([][]string,0)
	for _,v := range lines{
		if len(v)==0{
			continue
		}
		v = strings.TrimSpace(v) // remove tabs and other annoying whitespace
		remove,_ := regexp.MatchString("^//.*",v)
		if !remove{
			r := strings.Fields(v)
			res := make([]string,0)
			notComment := true
			// remove same line comments
			for _,s := range r{
				if s == "//"{
					notComment = false
				}
				if notComment{
					res = append(res, s)
				}
			} 
			clean = append(clean, res)
		}
	}
	for _,v := range clean{
		fmt.Println(v)
	}

	// init
	// assembly,s := "@256\nD=A\n@SP\nM=D\n",""
	assembly,s := "",""
	// go thru data and convert each line into corresponding hack asm instr(s)
	for _,v := range clean{
		if v[0] == "push"{
			s=translatePush(v,fileName)
		} else if v[0] == "pop"{
			s=translatePop(v,fileName)
		} else if v[0] == "call"{
			s=translateCall(v)
		} else if v[0] == "label"{
			s=translateLabel(v)
		} else if v[0] == "function"{
			s=translateFunctionHeader(v)
		} else if v[0] == "if-goto"{
			s=translateIfGoto(v)	
		} else if v[0] == "goto"{
			s=translateGoto(v)
		} else if v[0] == "return"{
			s=translateReturn()
		} else{
			s=translateArith(v)
		}
		assembly+=s
	}
	fmt.Println(assembly)
	return assembly
}




func main(){
	args := os.Args
	if len(args)!= 2{
		println("error only one arg allowed! \nUSAGE: ./VMTranslator <VM_FILE_PATH | DIR_PATH>")
		os.Exit(1)
	}

	fmt.Println(args)
	
	arginfo,err := os.Stat(args[1])
	if err != nil{
		panic(err)
	}
	lines,filePath, fileName:= [][]string{},"",""
	singleFileLines := []string{}
	res := bootstrap(args[1])
	fileNames := make([]string, 0)

	// =============================
	// GIVEN DIR W/ VM FILE(S) INSIDE 
	// =============================
	
	if arginfo.IsDir(){	
		lines,filePath,fileNames = handleDir(args[1])
		for i,v := range lines{
			res += assemble(v,fileNames[i])
		}
	
	// =============================
	// SINGLE VM FILE GIVEN
	// =============================
	} else{
		singleFileLines,filePath,fileName = handleSingleFile(args[1])
		res += assemble(singleFileLines,fileName)
	}
	fmt.Println(lines,filePath, fileName)
	out := fmt.Sprintf("%s.asm",filePath)
	writeToFile(out,res)
}
