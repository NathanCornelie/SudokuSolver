
#include <algorithm>
#include <iostream>
#include <vector>
#include "tools.h"

using namespace std;


vector<tuple<int, int>> findPlacesInBloc(int grid[9][9], int bloc) {
    const int rowStart = 3 * (bloc / 3);
    const int colStart = 3 * (bloc % 3);
    vector<tuple<int, int>> findPositions;
    for (int j = rowStart; j <= rowStart + 2; j++) {
        for (int k = colStart; k <= colStart + 2; k++) {
            if (grid[j][k] == 0)
                findPositions.emplace_back(j, k);
        }
    }
    return findPositions;

}

vector<tuple<int, int>> findPlacesInRow(int grid[9][9], int row) {
    vector<tuple<int, int>> findPositions;
    for (int ind = 0; ind < 9; ++ind) {
        if (row == 2) cout << grid[row][ind] << ',';

        if (grid[row][ind] == 0) findPositions.emplace_back(row, ind);
    }
    return findPositions;
}

vector<tuple<int, int>> findPlacesInCol(int grid[9][9], int col) {
    vector<tuple<int, int>> findPositions;
    for (int ind = 0; ind < 9; ++ind) {
        if (grid[ind][col] == 0) findPositions.emplace_back(ind, col);
    }
    return findPositions;
}

int findSimpleNumber(int grid[9][9], int cible) {
    int copy[9][9];
    copyGrid(grid, copy);
    for (int col = 0; col < 9; col++) {
        for (int row = 0; row < 9; row++) {
            if (copy[row][col] == cible) {
                for (int ind = 0; ind < 9; ind++) {
                    if (!copy[row][ind]) {
                        copy[row][ind] = 10;
                        placeNumberInGrid(copy, row, ind, 10, false);
                        //cout << 1 << "peut pas aller en " << row << "," << ind << ";\n";
                    }
                    if (!copy[ind][col]) {
                        copy[ind][col] = 10;
                        placeNumberInGrid(copy, ind, col, 10, false);
                        //cout << 1 << "peut pas aller en " << ind << "," << col << ";\n";

                    }
                }
            }
        }
    }

    for (int i = 0; i < 9; ++i) {
        for (int j = 0; j < 9; ++j) {
            vector<tuple<int, int> > positions = findPlacesInBloc(copy, j);
            if (positions.size() == 1) {
                cout << "bloc";
                if (!isNumberInBloc(copy, j, cible)) {
                    placeNumberInGrid(copy, get<0>(positions.at(0)), get<1>(positions.at(0)), cible, false);
                    fillEmptyPlacesInBloc(copy, j);
                    fillEmptyPlacesInRow(copy,get<0>(positions.at(0)));
                    fillEmptyPlacesInCol(copy,get<1>(positions.at(0)));

                }
                if (!isNumberInBloc(grid, j, cible))
                    placeNumberInGrid(grid, get<0>(positions.at(0)), get<1>(positions.at(0)), cible,true);
            }
                positions = findPlacesInRow(copy, j);
                if (positions.size() == 1) {
                    cout << "row";
                    if (!isNumberInRow(copy, j, cible)) {
                        placeNumberInGrid(copy, get<0>(positions.at(0)), get<1>(positions.at(0)), cible, false);
                        fillEmptyPlacesInRow(copy, j);
                        fillEmptyPlacesInCol(copy,get<1>(positions.at(0)));
                        fillEmptyPlacesInBloc(copy, findBlocFromCoordinate(j,get<1>(positions.at(0))));

                    }
                    if (!isNumberInRow(grid, j, cible))
                        placeNumberInGrid(grid, get<0>(positions.at(0)), get<1>(positions.at(0)), cible,true);

                }
                positions = findPlacesInCol(copy, j);
                if (positions.size() == 1) {
                    cout << "col";
                    if (!isNumberInCol(copy, j, cible)) {
                        placeNumberInGrid(copy, get<0>(positions.at(0)), get<1>(positions.at(0)), cible, false);
                        fillEmptyPlacesInCol(copy, j);
                        fillEmptyPlacesInRow(copy,get<0>(positions.at(0)));
                        fillEmptyPlacesInBloc(copy, findBlocFromCoordinate(get<0>(positions.at(0)),j));

                    }
                    if (!isNumberInCol(grid, j, cible))
                        placeNumberInGrid(grid, get<0>(positions.at(0)), get<1>(positions.at(0)), cible,true);

                }
            }
        }
    }