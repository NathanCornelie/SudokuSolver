//
// Created by nathan on 19/02/24.
//
#include <algorithm>
#include <iterator>
#include <iostream>
#include "verifMethods.h"


bool isGridValid(int grid[9][9]) {
    for (int i = 1; i <= 9; i++) {
        if (!isRowValid(grid, i) || !isColValid(grid, i) || !isBlocValid(grid, i)) return false;
    }
    return true;

}

bool isColValid(const int grid[9][9], const int col) {
    for (int i = 1; i <= 9; i++) {
        for (int j = 0; j < 9; j++) {
            if (i == grid[col - 1][j]) {
                break;
            } else if (j == 8 || grid[col - 1][j] > 9)
                return false;
        }
    }
    return true;

}

bool isRowValid(const int grid[9][9], const int row) {

    for (int i = 1; i <= 9; i++) {
        for (int j = 0; j < 9; j++) {
            if (i == grid[row - 1][j]) {
                break;
            } else if (j == 8 || grid[row - 1][j] > 9)
                return false;
        }
    }
    return true;
}

bool isBlocValid(const int grid[9][9], int bloc) {
    const int lStart = 3 * (bloc / 3);
    const int cStart = 3*(bloc % 3);
    bool find;
    for (int i = 1; i <= 9; i++) {
        find = false;
        for (int j = lStart; j <= lStart + 2; j++) {
            for (int k = cStart; k <= cStart + 2; k++) {
                if (i == grid[j][k]) {
                    find = true;
                    break;
                } else if (k == 2 && j == 2)
                    return false;
            }
            if (find) break;
        }
    }
    return true;

}