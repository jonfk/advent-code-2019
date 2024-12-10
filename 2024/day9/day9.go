package day9

import (
	"log"
	"slices"
	"strconv"
)

const EMPTY_BLOCK = -1

type Input struct {
	blocks []int
}

type ContiguousBlock struct {
	IsEmpty bool
	FileId  int
	Size    int
}

type ContiguousBlocks []ContiguousBlock

func (c ContiguousBlocks) toBlocks() []int {
	var blocks []int

	for _, block := range c {
		fileSize := block.Size
		if block.IsEmpty {
			for range fileSize {
				blocks = append(blocks, EMPTY_BLOCK)
			}
		} else {
			for range fileSize {
				blocks = append(blocks, block.FileId)
			}
		}
	}
	return blocks
}

func Parse(input string) Input {
	var blocks []int
	var fileId int
	for i, ch := range input {
		if ch == '\n' {
			continue
		}

		if i%2 != 0 {
			fileSize, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalf("fileSize conversion error: char = %v", string(ch))
			}
			for range fileSize {
				blocks = append(blocks, EMPTY_BLOCK)
			}
		} else {
			fileSize, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalf("fileSize conversion error: char = %v", string(ch))
			}
			for range fileSize {
				blocks = append(blocks, fileId)
			}
			fileId++
		}

	}
	return Input{blocks}
}

func Parse2(input string) []ContiguousBlock {
	var blocks []ContiguousBlock
	var fileId int
	for i, ch := range input {
		if ch == '\n' {
			continue
		}

		if i%2 != 0 {
			fileSize, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalf("fileSize conversion error: char = %v", string(ch))
			}
			blocks = append(blocks, ContiguousBlock{IsEmpty: true, Size: fileSize})
		} else {
			fileSize, err := strconv.Atoi(string(ch))
			if err != nil {
				log.Fatalf("fileSize conversion error: char = %v", string(ch))
			}
			blocks = append(blocks, ContiguousBlock{Size: fileSize, FileId: fileId})
			fileId++
		}

	}
	return blocks
}

func Compact(inputBlocks []int) []int {
	blocks := make([]int, len(inputBlocks))
	copy(blocks, inputBlocks)
	var compacted bool
	for {
		if compacted {
			break
		}
		var fileBlockIdx int
		for i := range blocks {
			idx := len(blocks) - 1 - i
			if blocks[len(blocks)-1-i] != EMPTY_BLOCK {
				fileBlockIdx = idx
				break
			}
		}

		for i := range blocks {
			if i == fileBlockIdx {
				compacted = true
				break
			}
			if blocks[i] == EMPTY_BLOCK {
				blocks[i], blocks[fileBlockIdx] = blocks[fileBlockIdx], blocks[i]
				break
			}
		}
	}
	return blocks
}

func CompactByFile(blocks []ContiguousBlock) {

	for fIdx := len(blocks) - 1; fIdx >= 0; fIdx-- {
		if !blocks[fIdx].IsEmpty {
			for eIdx := 0; eIdx < len(blocks); eIdx++ {
				if eIdx >= fIdx {
					break
				}
				if blocks[eIdx].IsEmpty && blocks[eIdx].Size >= blocks[fIdx].Size {
					emptyBlocksLeft := blocks[eIdx].Size - blocks[fIdx].Size
					blocks[eIdx].Size = blocks[fIdx].Size
					blocks[fIdx], blocks[eIdx] = blocks[eIdx], blocks[fIdx]

					if emptyBlocksLeft > 0 {
						blocks = slices.Insert(blocks, eIdx+1, ContiguousBlock{IsEmpty: true, Size: emptyBlocksLeft})

						// blocks = Normalize(blocks)
					}
					break
				}
			}
		}

	}
}

func Normalize(blocks []ContiguousBlock) []ContiguousBlock {
	for i := 0; i < len(blocks); i++ {
		if i+1 < len(blocks) && blocks[i].IsEmpty && blocks[i+1].IsEmpty {
			blocks[i].Size = blocks[i].Size + blocks[i+1].Size
			blocks = slices.Delete(blocks, i+1, i+2)
			i -= 1
		}
	}
	return blocks
}

func CalculateChecksumCompacted(blocks []int) int {
	checksum := 0
	for i := range blocks {
		if blocks[i] == EMPTY_BLOCK {
			continue
		}

		checksum += i * blocks[i]
	}
	return checksum
}
