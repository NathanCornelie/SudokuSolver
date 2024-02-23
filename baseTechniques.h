#include <algorithm>
#include <iostream>
#include <vector>

#ifndef TESTCPP_BASETECHNIQUES_H
#define TESTCPP_BASETECHNIQUES_H
using namespace std;
int findSimpleNumber(int grid[9][9],int cible);

vector<tuple<int, int>> findPlacesInRow(int grid[9][9], int row ) ;
vector<tuple<int, int>> findPlacesInCol(int grid[9][9], int col) ;
vector<tuple<int, int>> findPlacesInBloc(int grid[9][9], int bloc) ;


#endif //TESTCPP_BASETECHNIQUES_H
