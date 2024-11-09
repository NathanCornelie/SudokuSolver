cmake_minimum_required(VERSION 3.27)
project(sudokuSolver)

set(CMAKE_CXX_STANDARD 20)

add_executable(sudokuSolver main.cpp
        Tools/verifMethods.cpp
        Tools/verifMethods.h
        Techniques/baseTechniques.cpp
        Techniques/baseTechniques.h
        Tools/tools.cpp
        Tools/tools.h
        solver.cpp
        solver.h
        Techniques/singleNaked.cpp
        Techniques/singleNaked.h)
