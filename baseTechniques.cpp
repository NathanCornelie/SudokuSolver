
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

void findSimpleNumber(int grid[9][9], int cible , int manquants[9]) {
    // trouve et place tout les occurence de la cible plaçable dans la grille
    int copy[9][9];
    copyGrid(grid, copy);
    // * on remplis la grid avec le chiffre 10 partout ou la cible ne peut pas aller
    for (int col = 0; col < 9; col++) {
        for (int row = 0; row < 9; row++) {
            if (copy[row][col] == cible) {
                for (int ind = 0; ind < 9; ind++) {
                    if (copy[row][ind] == 0) {
                        placeNumberInGrid(copy, row, ind, 10 );
                    }
                    if (!copy[ind][col]) {
                        placeNumberInGrid(copy, ind, col, 10);

                    }
                }
            }
        }
    }

    // si on trouve une ligne, une colonne ou dans un bloc on y ajoute le chiffre et on met à jours les 10
    for (int i = 0; i < 9; ++i) {
        for (int j = 0; j < 9; ++j) {
            vector<tuple<int, int> > positionsDisponibles = findPlacesInBloc(copy, j);
            if (positionsDisponibles.size() == 1) {

                if (!isNumberInBloc(copy, j, cible)) {
                    placeNumberInGrid(copy, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)), cible, false);
                    fillEmptyPlacesInBloc(copy, j);
                    fillEmptyPlacesInRow(copy,get<0>(positionsDisponibles.at(0)));
                    fillEmptyPlacesInCol(copy,get<1>(positionsDisponibles.at(0)));

                }
                if (!isNumberInBloc(grid, j, cible)) {
                    placeNumberInGrid(grid, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)),
                                      cible, true);
                    manquants[cible - 1]--;
                } cout << " (bloc) \n";
            }
            positionsDisponibles = findPlacesInRow(copy, j);
                if (positionsDisponibles.size() == 1) {

                    if (!isNumberInRow(copy, j, cible)) {
                        placeNumberInGrid(copy, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)), cible, false);
                        fillEmptyPlacesInRow(copy, j);
                        fillEmptyPlacesInCol(copy,get<1>(positionsDisponibles.at(0)));
                        fillEmptyPlacesInBloc(copy, findBlocFromCoordinate(j,get<1>(positionsDisponibles.at(0))));

                    }
                    if (!isNumberInRow(grid, j, cible)) {
                        placeNumberInGrid(grid, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)),
                                          cible, true);
                         manquants[cible - 1]--;
                    } cout << " (row)\n";
                }
            positionsDisponibles = findPlacesInCol(copy, j);
                if (positionsDisponibles.size() == 1) {

                    if (!isNumberInCol(copy, j, cible)) {
                        placeNumberInGrid(copy, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)), cible, false);
                        fillEmptyPlacesInCol(copy, j);
                        fillEmptyPlacesInRow(copy,get<0>(positionsDisponibles.at(0)));
                        fillEmptyPlacesInBloc(copy, findBlocFromCoordinate(get<0>(positionsDisponibles.at(0)), j));

                    }
                    if (!isNumberInCol(grid, j, cible)) {
                        placeNumberInGrid(grid, get<0>(positionsDisponibles.at(0)), get<1>(positionsDisponibles.at(0)),
                                          cible, true);
                        manquants[cible - 1]--;
                    }
                    cout << " (col)\n";
                }
            }
        }
    }