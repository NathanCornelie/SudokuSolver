#include <iostream>

int copyGrid(int grid[9][9], int dest[9][9]) {
    for (int col = 0; col < 9; col++)
        for (int row = 0; row < 9; row++) {
            dest[row][col] = grid[row][col];
        }
    return 1;
}

void placeNumberInGrid(int grid[9][9], int row, int col, int num, bool verbose = false) {
    grid[row][col] = num;
    if (verbose)std::cout << num << " placé en position" << row << "," << col << ";\n";
}

bool isNumberInRow(int grid[9][9], int row, int number) {
    for (int col = 0; col < 9; col++) {
        if (grid[row][col] == number) return true;
    }
    return false;
}

bool isNumberInCol(int grid[9][9], int col, int number) {
    for (int row = 0; row < 9; row++) {
        if (grid[row][col] == number) return true;
    }
    return false;
}

bool isNumberInBloc(int grid[9][9], int bloc, int number) {
    const int rowStart = 3 * (bloc / 3);
    const int colStart = 3 * (bloc % 3);
    for (int j = rowStart; j <= rowStart + 2; j++) {
        for (int k = colStart; k <= colStart + 2; k++) {
            if (grid[j][k] == number)return true;
        }
    }
    return false;
}

void fillEmptyPlacesInBloc(int grid[9][9], int bloc) {
    const int rowStart = 3 * (bloc / 3);
    const int colStart = 3 * (bloc % 3);
    for (int row = rowStart; row <= rowStart + 2; row++) {
        for (int col = colStart; col <= colStart + 2; col++) {
            if (grid[row][col] == 0)placeNumberInGrid(grid, row, col, 10);
        }
    }

}

void fillEmptyPlacesInRow(int grid[9][9], int row) {
    for (int col = 0; col < 9; col++) {
        if (grid[row][col] == 0) placeNumberInGrid(grid, row, col, 10);
    }

}

void fillEmptyPlacesInCol(int grid[9][9], int col) {
    for (int row = 0; row < 9; row++) {
        if (grid[row][col] == 0) placeNumberInGrid(grid, row, col, 10);
    }

}

int findBlocFromCoordinate(int row , int col){
    return 3*(row / 3)+(col/3);
}
