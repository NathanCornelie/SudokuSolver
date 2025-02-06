import {Grid} from "./Grid";
export default class Solution {
  solution: SolutionMethod;
  grid: Grid;

  constructor( solution: SolutionMethod,grid: Grid = new Grid()) {
    this.solution = solution
    this.grid = grid;
  }
}

class SolutionMethod {
  value: number;
  row: number;
  col: number;
  method: string;
  type: string;
  constructor(
    value: number,
    row: number,
    col: number,
    method: string,
    type: string
  ) {
    this.value = value;
    this.row = row;
    this.col = col;
    this.method = method;
    this.type = type;
  }
}
