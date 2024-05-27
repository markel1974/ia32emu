package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"markel/ia32emu/core"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	defaultBaseAddress  = 0x7c00
	defaultStackAddress = 0x7c04
)

func createRamFromKsTool(in string) ([]byte, error) {
	end := strings.LastIndex(in, "]")
	if end < 0 {
		return nil, errors.New("missing ]")
	}
	partial := in[:end]
	start := strings.LastIndex(partial, "[")
	if start < 0 {
		return nil, errors.New("missing [")
	}
	l := partial[start+1:]
	l = strings.TrimSpace(l)
	var ram []byte
	data := strings.Split(l, " ")
	for _, d := range data {
		if len(d) != 2 {
			continue
		}
		n, err := strconv.ParseInt(d, 16, 64)
		if err != nil {
			continue
		}
		ram = append(ram, byte(n))
	}
	return ram, nil
}

func checkPath(filePath string) (string, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return "", fmt.Errorf("no binary file specified or found")
	}
	if info.IsDir() {
		files, err := os.ReadDir(filePath)
		if err != nil {
			return "", fmt.Errorf("no binary file specified or found")
		}
		name := files[0].Name()
		return path.Join(filePath, name), nil
	}
	return filePath, nil
}

func loadRamFile(path string) ([]byte, error) {
	filePath, err := checkPath(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	memSize := uint64(fileInfo.Size())
	ram := make([]byte, memSize)
	f, _ := os.Open(filePath)
	copySize, _ := io.ReadFull(f, ram)
	if int64(copySize) != fileInfo.Size() {
		return nil, errors.New("size not matched")
	}
	return ram, nil
}

func main() {
	var filePath string
	var showHelp bool
	var debugFlag bool
	var windowFlag bool
	var bitMode int
	var baseAddress int
	var stackAddress int
	flag.IntVar(&baseAddress, "b", defaultBaseAddress, "begin address")
	flag.IntVar(&stackAddress, "s", defaultStackAddress, "stack address")
	flag.IntVar(&bitMode, "x", 32, "bit mode")
	flag.BoolVar(&windowFlag, "w", false, "window mode")
	flag.BoolVar(&debugFlag, "d", false, "debug mode")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.StringVar(&filePath, "p", "", "file to run")

	if showHelp {
		flag.Usage()
		return
	}

	log.SetFlags(0)

	sample()

	ram, err := loadRamFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	emu, err := core.NewEmulator(bitMode, uint32(baseAddress), uint32(stackAddress), ram, debugFlag)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := emu.Run(); err != nil {
		log.Println(err.Error())
	}
	emu.Dump()
}

func sample() {
	//input := "mov eax, 0x1;cmp eax, 0x2;jnz not_equal;equal:;jmp 0;not_equal:;mov eax, 0x2;cmp eax, 0x2;jz equal; = [ b8 01 00 00 00 83 f8 02 75 02 eb f4 b8 02 00 00 00 83 f8 02 74 f4 ]"
	//input := "mov eax, 0x60;mov ebx, 0x10;sub eax, 0x10;sub eax, ebx; = [ b8 60 00 00 00 bb 10 00 00 00 83 e8 10 29 d8 ]"
	//input := "mov eax, 0xf1;mov ebx, 0x29;call swap;jmp 0;swap:;mov ecx, ebx;mov ebx, eax;mov eax, ecx;ret; = [ b8 f1 00 00 00 bb 29 00 00 00 e8 02 00 00 00 eb ef 89 d9 89 c3 89 c8 c3 ]"
	input := "start:;mov eax, 0x00f1;mov ebx, 0x0029;call add_routine;jmp 0;add_routine:;mov ebx, eax;mov eax, 0x1011;ret; = [ b8 f1 00 00 00 bb 29 00 00 00 e8 02 00 00 00 eb ef 89 c3 b8 11 10 00 00 c3 ]"
	ram, err := createRamFromKsTool(input)
	if err != nil {
		log.Println(err.Error())
		return
	}
	emu, err := core.NewEmulator(32, defaultBaseAddress, defaultStackAddress, ram, true)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := emu.Run(); err != nil {
		log.Println(err.Error())
	}
	emu.Dump()
	os.Exit(1)
}
