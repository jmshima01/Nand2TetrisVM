1) James Shima
2) Few days, took me a while to learn the memory segmegmentation before I could code
3) Sorry this is 2 days late midterms been really rough :/ 
I used `go` again and what really helped me was looking back a lot at `Pong.asm` and that helped me realize what each instruction should 
look like in an optimal way (kinda a little reverse engineering).

The coolest thing I did was instead of simulating the entire stack and keeping track of vars for the equality
operators I used asm instead taking advantage of the JMP instructions and labels. (Why would I simulate the CPU thats 
memory intensive).

4) 
run `make` with latest version of go installed (see `LANGINFO` if this is an issue).
A executable bin named `Virtual` will be made.
Then pass `Virtual` a .vm path and a .asm file with be written to that same path (as per textbook instrs)

`./Virtual path/to/vmfile`