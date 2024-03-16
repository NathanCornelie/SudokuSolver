//
// Created by nathan on 29/02/24.
//

#include <vector>
#include <algorithm>
#include <iostream>
#include "singleNaked.h"
#include "../Tools/tools.h"
using namespace std;

void findSingleNaked(int grid[9][9]){
// contient les chiffres manquants dans la ligne i
    vector<int> manquants;


    for (int i = 0; i < 9; ++i) { // <- parcours des lignes
        //initialisations des manquants
        for (int j = 0; j <9; ++j) {
            if(grid[i][j]) manquants.push_back(grid[i][j]);
        }
        for (int j = 0; j < 9; ++j) { // <- parcours des cases de la ligne
            if(!grid[i][j]){
                int indiceBloc = findBlocFromCoordinate(i,j);
                vector<int> possibles;
                 std::copy(manquants.begin(), manquants.end(),std::back_inserter(possibles));
                for (int k = 0; k < possibles.size(); ++k) {
                    if(isNumberInBloc(grid,indiceBloc,possibles[k]) || isNumberInCol(grid,j,possibles[k])){
                        possibles.erase(std::remove(possibles.begin(), possibles.end(), 8), possibles.end());
                    }
                }
                if(possibles.size()==1){
                    placeNumberInGrid(grid,i,j,possibles[0],true);
                    std::cout << "Single Naked \n";
                }

            }
        }
    }
}