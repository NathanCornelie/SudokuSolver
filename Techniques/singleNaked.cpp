//
// Created by nathan on 29/02/24.
//

#include <vector>
#include <algorithm>
#include <iostream>
#include "singleNaked.h"
#include <unordered_set>
#include "../Tools/tools.h"

using namespace std;

void findSingleNaked(int grid[9][9]) {
    unordered_set<int> manquants(9);
    unordered_set<int> presents;

    for (int i = 0; i < 9; ++i) {
        manquants.clear();
        presents.clear();
        for (int j = 0; j < 9; ++j) {
            if (grid[i][j]) presents.insert(grid[i][j]);
        }

        for (int j = 0; j < 9; ++j) {
            if(!presents.contains(j))manquants.insert(j);
        }
        for (int j = 0; j < 9; ++j) { // <- parcours des cases de la ligne
            if (!grid[i][j]) {
                int indiceBloc = findBlocFromCoordinate(i, j);
                unordered_set<int> possibles = manquants;
                for (auto itr = possibles.begin();itr!= possibles.end();) {
                    if (i == 6) std::cout << possibles.size() << "\n";
                    if (isNumberInBloc(grid, indiceBloc, *itr) || isNumberInCol(grid, j, *itr))itr = possibles.erase(itr);
                    else++itr;
                }
                if (possibles.size() == 1) {
                    placeNumberInGrid(grid, i, j, *possibles.begin(), true);
                    std::cout << "Single Naked \n";
                }

            }
        }
    }
}