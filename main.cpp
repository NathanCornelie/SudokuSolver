

#include "Techniques/baseTechniques.h"
#include "solver.h"

int main() {

    int grid[9][9] = {{8, 2, 3, 4, 5, 6, 7, 8, 9},
                      {4, 0, 6, 4, 5, 6, 7, 1, 9},
                      {7, 0, 9, 4, 5, 6, 7, 8, 9},
                      {4, 2, 3, 4, 5, 6, 7, 8, 9},
                      {5, 2, 3, 4, 5, 6, 7, 8, 9},
                      {6, 2, 3, 4, 5, 6, 7, 8, 9},
                      {7, 2, 3, 4, 5, 6, 7, 8, 9},
                      {8, 2, 3, 4, 5, 6, 7, 8, 9},
                      {9, 2, 3, 4, 5, 6, 7, 8, 9}};

    int correctEasyGrid[9][9] = {{0, 0, 8, 0, 2, 4, 0, 1, 0},
                                 {0, 9, 0, 0, 8, 6, 0, 5, 0},
                                 {3, 0, 2, 0, 1, 9, 6, 0, 0},
                                 {9, 2, 6, 1, 7, 8, 0, 0, 5},
                                 {8, 0, 0, 0, 4, 0, 0, 2, 0},
                                 {0, 7, 0, 6, 0, 0, 0, 0, 1},
                                 {6, 8, 0, 0, 0, 0, 4, 0, 9},
                                 {2, 1, 9, 0, 0, 0, 0, 0, 8},
                                  {0, 3, 0, 8, 9, 5, 1, 0, 2}};

    int singleNakedGrid[9][9] = {
            {0,0,8,0,0,7,9,0,0},
            {0,4,2,0,0,5,0,0,0},
            {0,0,0,6,0,0,0,5,0},
            {0,0,3,0,0,6,8,0,1},
            {0,0,0,0,0,0,0,0,6},
            {9,0,0,0,7,0,0,0,0},
            {0,8,0,1,3,0,4,7,0},
            {0,0,0,0,9,0,0,0,0},
            {0,1,0,0,0,0,0,0,0}
    };

    
    int test[9][9] = {
            {0,0,8,0,0,7,9,0,0},
            {0,4,2,0,0,5,0,0,0},
            {0,0,0,6,0,0,0,5,0},
            {0,0,3,0,0,6,8,0,1},
            {0,0,0,0,0,0,0,0,6},
            {9,0,0,0,7,0,0,0,0},
            {0,8,0,1,3,0,4,7,0},
            {0,0,0,0,9,0,0,0,0},
            {0,1,0,0,0,0,0,0,0}
    };
    solver(singleNakedGrid);
    return 0;
}
