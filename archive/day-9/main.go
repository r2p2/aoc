package main

import (
	"fmt"
	"math"
	"os"
)

func print(mem []uint64) {
	for i := 0; i < len(mem); i++ {
		if mem[i] == math.MaxUint64 {
			fmt.Print(".")
		} else {
			fmt.Print(mem[i])
		}
	}
	fmt.Println()
}

func checksum(mem []uint64) uint64 {
	sum := uint64(0)
	for i := 0; i < len(mem); i++ {
		if mem[i] == math.MaxUint64 {
			continue
		}

		sum += uint64(i) * mem[i]
	}
	return sum
}

type Block struct {
	size    uint64
	value   uint64
	isEmpty bool
}

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(bytes)
	memory_orig := []uint64{}
	n := uint64(0)
	for i := 0; i < len(input)-1; i++ {
		is_data := (i + 1) % 2
		for j := 0; j < int(input[i]-'0'); j++ {
			if is_data == 1 {
				memory_orig = append(memory_orig, n)
			} else {
				memory_orig = append(memory_orig, math.MaxUint64)
			}
		}
		if is_data == 1 {
			n++
		}
	}

	// print(memory)

	{
		var memory []uint64
		memory = append(memory, memory_orig...)

		for s := len(memory); s > 0; s-- {
			if memory[s-1] == math.MaxUint64 {
				continue
			}

			for d := 0; d < s-1; d++ {
				if memory[d] != math.MaxUint64 {
					continue
				}

				memory[d] = memory[s-1]
				memory[s-1] = math.MaxUint64
				// print(memory)
			}
		}

		fmt.Println("part 1:", checksum(memory))
	}

	// print(memory_orig)
	{
		var memory []uint64
		memory = append(memory, memory_orig...)

		for se := len(memory); se > 0; se-- {
			if memory[se-1] == math.MaxUint64 {
				continue
			}

			sw := 1
			for ; se-sw > 0; sw++ {
				if memory[se-sw-1] != memory[se-1] {
					break
				}
			}

			//for ds := 0; ds < len(memory); ds++ {
			for ds := 0; ds < se-sw; ds++ {
				if memory[ds] != math.MaxUint64 {
					continue
				}

				w := 1
				for ; ds+w < len(memory); w++ {
					if memory[ds+w] != memory[ds] {
						break
					}
				}

				if sw <= w {
					// yay copy it

					for x := 0; x < sw; x++ {
						memory[ds+x] = memory[se-x-1]
						memory[se-x-1] = math.MaxUint64
					}

					break
				}

				ds += w
			}

			// print(memory)
			se -= sw - 1
		}
		fmt.Println("part 2:", checksum(memory))
	}

}
