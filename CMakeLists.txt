cmake_minimum_required(VERSION 3.27)
project(testCPP)

set(CMAKE_CXX_STANDARD 20)

add_executable(testCPP main.cpp
        verifMethods.cpp
        verifMethods.h
        baseTechniques.cpp
        baseTechniques.h
        tools.cpp
        tools.h)
