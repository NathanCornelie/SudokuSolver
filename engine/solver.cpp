//
// Created by nathan on 24/02/24.
//
#include "Tools/verifMethods.h"
#include "solver.h"
#include "Techniques/baseTechniques.h"
#include "Techniques/singleNaked.h"


void initManquants(int grid[9][9], int manquant[9]) {
    int finds[9];
    for (int i = 0; i < 9; ++i) {
        for (int j = 0; j < 9; ++j) {
            if (grid[i][j]) {
                finds[grid[i][j] - 1]++;
            }
        }
    }
    for (int i = 0; i < 9; ++i) {
        manquant[i] = 9 - finds[i];
    }
}

void solver(int grid[9][9]) {
    int nbrChiffreManquant[9];
    initManquants(grid, nbrChiffreManquant);
    bool numberPlaced = false;
    while (!isGridValid(grid)) {
        for (int i = 0; i < 9; ++i) {

            if(nbrChiffreManquant[i]) {
                findSimpleNumber(grid, i + 1, nbrChiffreManquant,numberPlaced);
            }
            if(!numberPlaced){
            findSingleNaked(grid,numberPlaced);
        }
        }
        
    }
}
