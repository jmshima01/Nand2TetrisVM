#!/bin/env bash

set -e

make
./Virtual 7/StackArithmetic/StackTest/StackTest.vm > /dev/null
./Virtual 7/StackArithmetic/SimpleAdd/SimpleAdd.vm > /dev/null
./Virtual 7/MemoryAccess/BasicTest/BasicTest.vm > /dev/null
./Virtual 7/MemoryAccess/StaticTest/StaticTest.vm > /dev/null
./Virtual 7/MemoryAccess/PointerTest/PointerTest.vm > /dev/null

echo "generated all test .asm files"
