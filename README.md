# Nand2TetrisVM
1) James Shima
2) Few days, took me a while to learn the memory segmegmentation before I could code
3) Sorry this is 2 days late midterms been really rough :/ 
What really helped me was looking back a lot at `Pong.asm` and that helped me realize what each instruction should 
look like in an optimal way (kinda a little reverse engineering).

The coolest thing I did was instead of simulating the entire stack and keeping track of vars for the eq
operators I used hardware/asm instead. For example, eq -> (A & B) | (A|B)!