cmake_minimum_required(VERSION 3.27)
project(sudokuSolver)

set(CMAKE_CXX_STANDARD 20)

add_executable(sudokuSolver main.cpp
        verifMethods.cpp
        verifMethods.h
        baseTechniques.cpp
        baseTechniques.h
        tools.cpp
        tools.h
        solver.cpp
        solver.h)
