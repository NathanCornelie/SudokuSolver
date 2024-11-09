//
// Created by nathan on 20/02/24.
//

#ifndef TESTCPP_TOOLS_H
#define TESTCPP_TOOLS_H

int copyGrid(int grid[9][9], int dest[9][9]);

void placeNumberInGrid(int grid[9][9], int row, int col, int num, bool verbose = false);

bool isNumberInRow(int grid[9][9], int row, int number);

bool isNumberInCol(int grid[9][9], int col, int number);

bool isNumberInBloc(int grid[9][9], int bloc, int number);

void fillEmptyPlacesInBloc(int grid[9][9], int bloc);

void fillEmptyPlacesInRow(int grid[9][9], int row);

void fillEmptyPlacesInCol(int grid[9][9], int col);
int findBlocFromCoordinate(int row , int col);

#endif //TESTCPP_TOOLS_H
