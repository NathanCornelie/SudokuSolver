//
// Created by nathan on 29/02/24.
//

#include <vector>
#include "singleNaked.h"
#include "../Tools/tools.h"
using namespace std;

void findSingleNaked(int grid[9][9]){

    vector<int> manquants;

    // contient les chiffres manquants dans la ligne i
    for (int i = 0; i < 9; ++i) { // <- parcours des lignes
        //initialisations des manquants
        for (int j = 0; j <9; ++j) {
            if(grid[i][j]) manquants.push_back(grid[i][j]);
        }
        for (int j = 0; j < 9; ++j) { // <- parcours des cases de la ligne
            if(!grid[i][j]){
                int indiceBloc = findBlocFromCoordinate(i,j);
                

            }
        }
    }
}