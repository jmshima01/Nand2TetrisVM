1) James Shima
2) few days
3) `Pong.asm` helped me a lot again and just understanding how the stack frames are ordered was needed.

4) To run my project you should do the following
- first make sure you have golang installed (see `LANGINFO`)
- then run `make` and this should create a executable called `VMTranslator`
- then you can run the bin (`chmod +x` ran by Makefile) using...
`./VMTranslator <path/to/dir>`
or
`./VMTranslator <path/to/single/vm/file>`

example...
`./VMTranslator tests/FunctionCalls/StaticsTest` will create `StaticsTest.asm` in the same dir as the path provided.
for issues email: jamesshima@mines.edu