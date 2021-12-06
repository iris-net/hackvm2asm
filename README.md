# hackvm2asm
This is translation tool that translates vm to assembly for HACK.

Since this is for learning, this don't consider any invalid vm cases.

# execute
```
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\BasicTest\BasicTest.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\BasicTest\BasicTest.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\PointerTest\PointerTest.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\PointerTest\PointerTest.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\StaticTest\StaticTest.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\MemoryAccess\StaticTest\StaticTest.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\StackArithmetic\SimpleAdd\SimpleAdd.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\StackArithmetic\SimpleAdd\SimpleAdd.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\StackArithmetic\StackTest\StackTest.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\StackArithmetic\StackTest\StackTest.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\ProgramFlow\BasicLoop\BasicLoop.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\ProgramFlow\BasicLoop\BasicLoop.asm"
go run ./main.go exec --in="{{REPLACE_YOUR_DIRECTORY}}\ProgramFlow\FibonacciSeries\FibonacciSeries.vm" --out="{{REPLACE_YOUR_DIRECTORY}}\ProgramFlow\FibonacciSeries\FibonacciSeries.asm"


```
