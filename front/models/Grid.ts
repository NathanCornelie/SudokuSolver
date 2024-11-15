class Grid {
  grid: number[][];
  constructor(grid: number[][] | null = null) {
    if (grid) {
      this.grid = grid;
    } else {
      this.grid = Array.from({ length: 9 }, () => Array(9).fill(0));
    }
  }

  setValue(row: number, col: number, value: number): void {
    if (row < 0 || row >= 9 || col < 0 || col >= 9) {
      throw new Error("Row and column indices must be between 0 and 8.");
    }
    if (1 > value || value > 9) {
      throw new Error("grid value must be between 1 and 9");
    }
    this.grid[row][col] = value;
  }
}

export default Grid;
