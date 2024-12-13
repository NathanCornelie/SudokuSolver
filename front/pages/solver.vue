<template>
  <v-card
    variant="flat"
    class="bg-transparent py-7 d-flex flex-column align-center justify-space-between w-100"
  >
    <div style="width: fit-content" class="my-5">
      <div class="mb-5 d-flex align-center justify-space-around">
        <v-btn>Edit grid</v-btn>
      </div>
      <v-card class="mb-5">
        <v-container id="solutions-container">
          <v-row
            v-for="(row, i) in displayed_grid?.grid"
            :class="[i == 3 || i == 6 ? 'row-separator' : '']"
            dense
          >
            <v-col v-for="(value, j) in row">
              <v-card
                :elevation="getElevation(i, j)"
                :class="[
                  'cell',
                  getBaseRedBg(i, j),
                  j == 3 || j == 6 ? 'cell-separator' : '',
                  i == selectedCase.row && j == selectedCase.col
                    ? 'selected_case'
                    : '',

                  getBaseRedBg(i, j),
                ]"
                @click="handleSetSelectedCase(i, j)"
              >
                <p>{{ value }}</p>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
      <div class="mt-5 d-flex align-center justify-space-around">
        <v-btn @click="handleSolve()">View Solution</v-btn>
        <v-btn @click="handleGetNextSolution">Next Solution</v-btn>
        <v-btn @click="handleGetAllSolutions">All Solutions</v-btn>
      </div>
    </div>

    <Solutions
      :solutions="solutions ? solutions : []"
      :selected-solution-index="selectedSolutionIndex"
      text="hello"
      @select-solution="handleSelectSolutuion"
      @change-selected-solution="handleChangeSelectedSolutionIndex"
    />
  </v-card>
</template>

<script lang="ts" setup>
import SolverAPI from "~/api/SolverAPI";
import Grid from "~/models/Grid";
import Solution from "~/models/Solution";

const base_grid = ref<Grid>();
const displayed_grid = ref<Grid>();
const solutions = ref<Solution[]>([]);
const selectedSolutionIndex = ref<number>(0);
const selectedSolution = ref<Solution>();
const singleNakedColorationPositions = ref<{ row: number; col: number }[]>([]);
const selectedCase = ref<{ row: number; col: number }>({ row: 0, col: 0 });
const data_grid = ref<number[][]>([
  [5, 3, 0, 0, 7, 0, 0, 0, 0],
  [6, 0, 0, 1, 9, 5, 0, 0, 0],
  [0, 9, 8, 0, 0, 0, 0, 6, 0],
  [8, 0, 0, 0, 6, 0, 0, 0, 3],
  [4, 0, 0, 8, 0, 3, 0, 0, 1],
  [7, 0, 0, 0, 2, 0, 0, 0, 6],
  [0, 6, 0, 0, 0, 0, 2, 8, 0],
  [0, 0, 0, 4, 1, 9, 0, 0, 5],
  [0, 0, 0, 0, 8, 0, 0, 7, 9],
]);
onMounted(() => {
  base_grid.value = new Grid(data_grid.value);
  displayed_grid.value = base_grid.value;
});
const handleSolve = async () => {
  if (data_grid.value) {
    const solution = await SolverAPI.solve(new Grid(data_grid.value));

    if (solution) {
      displayed_grid.value = new Grid(solution.grid);
    }
  }
};
const handleGetNextSolution = async () => {
  if (displayed_grid.value) {
    const solution = await SolverAPI.getOneSolution(
      new Grid(displayed_grid.value.grid)
    );
    if (solution) {
      solutions.value?.push(solution);
      displayed_grid.value = new Grid(solution.grid.grid);
      selectedCase.value = {
        row: solution.solution.row,
        col: solution.solution.col,
      };
      handleSelectSolutuion(solution, solutions.value.length - 1);
    }
  }
};
const isValueInCol = (value: number, col: number) => {
  for (let i = 0; i < 9; i++) {
    if (displayed_grid.value?.grid[i][col] == value) {
      return true;
    }
  }
  return false;
};
const isValueInRow = (value: number, row: number) => {
  for (let i = 0; i < 9; i++) {
    if (displayed_grid.value?.grid[row][i] == value) {
      return true;
    }
  }
  return false;
};
const isvalueInBloc = (value: number, bloc: number) => {
  const rowStart = 3 * Math.floor(bloc / 3);
  const colStart = 3 * (bloc % 3);
  for (let i = colStart; i < colStart + 3; i++) {
    for (let j = rowStart; j < rowStart + 3; j++) {
      if (displayed_grid.value?.grid[j][i] == value) return true;
    }
  }
  return false;
};
const handleChangeSelectedSolutionIndex = (value: number) => {
  handleSelectSolutuion(solutions.value[value], value);
};
const getBlocFromCase = (row: number, col: number) => {
  return 3 * ((row / 3) >> 0) + ((col / 3) >> 0);
};

const getElevation = (row: number, col: number) => {
  if (
    selectedSolution.value &&
    selectedSolution.value.solution.row == row &&
    selectedSolution.value.solution.col == col
  )
    return 5;
  return 0;
};

const getBaseRedBg = (row: number, col: number) => {
  if (selectedSolution.value?.solution.method == "base") {
    const color = colorationBase(row, col);
    if (color) return color;
  } else {
    const color = colorationSingleNaked(row, col);
    if (color) return color;
  }

  return "";
};
const colorationBase = (row: number, col: number) => {
  if (selectedSolution.value?.solution.type == "bloc") {
    const bloc = getBlocFromCase(
      selectedSolution.value?.solution.row,
      selectedSolution.value?.solution.col
    );
    const rowStart = 3 * Math.floor(bloc / 3);
    const colStart = 3 * (bloc % 3);
    if (selectedSolution.value.solution.type == "bloc")
      if (
        rowStart <= row &&
        row <= rowStart + 2 &&
        row != selectedSolution.value.solution.row &&
        isValueInRow(selectedSolution.value.solution.value, row)
      ) {
        for (let i = colStart; i < colStart + 3; i++) {
          if (displayed_grid.value?.grid[row][i] == 0) return "red_bg";
        }
      }
    if (
      colStart <= col &&
      col <= colStart + 2 &&
      col != selectedSolution.value.solution.col &&
      isValueInCol(selectedSolution.value.solution.value, col)
    )
      for (let i = rowStart; i < rowStart + 3; i++) {
        if (displayed_grid.value?.grid[i][col] == 0) return "red_bg";
      }

    return "";
  }
  if (selectedSolution.value?.solution.type == "col") {
    const blocStart = selectedSolution.value.solution.row % 3;
    for (let b = blocStart; b <= blocStart + 2; b++) {
      const bloc = 3 * Math.floor(b / 3) + blocStart;
      console.log;
      if (isvalueInBloc(selectedSolution.value.solution.value, bloc)) {
        if (
          getBlocFromCase(
            selectedSolution.value.solution.row,
            selectedSolution.value.solution.col
          ) != bloc
        ) {
          const rowStart = 3 * Math.floor(bloc / 3);
          const colStart = 3 * (bloc % 3);
          let isRowFull = true;
          for (let i = colStart; i <= colStart + 2; i++) {
            if (
              displayed_grid.value?.grid[selectedSolution.value.solution.row][
                i
              ] == 0
            )
              isRowFull = false;
          }
          if (
            !isRowFull &&
            rowStart <= row &&
            row <= rowStart + 2 &&
            colStart <= col &&
            col <= colStart + 2
          )
            return "red_bg";
        }
      } else {
        if (
          row != selectedSolution.value.solution.row &&
          isValueInRow(selectedSolution.value.solution.value, row) &&
          displayed_grid.value?.grid[row][
            selectedSolution.value.solution.col
          ] == 0
        )
          return "red_bg";
      }
    }
  }
  if (selectedSolution.value?.solution.type == "row") {
    const blocStart = 3 * Math.floor(selectedSolution.value.solution.row / 3);
    for (let b = blocStart; b <= blocStart + 2; b++) {
      if (isvalueInBloc(selectedSolution.value.solution.value, b)) {
        if (
          getBlocFromCase(
            selectedSolution.value.solution.row,
            selectedSolution.value.solution.col
          ) != b
        ) {
          const rowStart = 3 * Math.floor(b / 3);
          const colStart = 3 * (b % 3);
          let isRowFull = true;
          for (let i = colStart; i <= colStart + 2; i++) {
            if (
              displayed_grid.value?.grid[selectedSolution.value.solution.row][
                i
              ] == 0
            )
              isRowFull = false;
          }
          if (
            !isRowFull &&
            rowStart <= row &&
            row <= rowStart + 2 &&
            colStart <= col &&
            col <= colStart + 2
          )
            return "red_bg";
        }
      } else {
        if (
          col != selectedSolution.value.solution.col &&
          isValueInCol(selectedSolution.value.solution.value, col) &&
          displayed_grid.value?.grid[selectedSolution.value.solution.row][
            col
          ] == 0
        )
          return "red_bg";
      }
    }
  }
};

const colorationSingleNaked = (row: number, col: number) => {
  let last_col = true;
  let last_row = true;
  let last_bloc = true;
  if (selectedSolution.value?.solution.type == "") {
    if (
      getBlocFromCase(row, col) ==
      getBlocFromCase(
        selectedSolution.value.solution.row,
        selectedSolution.value.solution.col
      )
    ) {
      if (isvalueInBloc(0, getBlocFromCase(row, col))) last_bloc = false;
      if (last_bloc) return "red_bg";
    }
  }
  if (
    selectedSolution.value?.solution.type == "col" &&
    col == selectedSolution.value?.solution.col
  ) {
    for (let i = 0; i < 9; i++) {
      if (displayed_grid.value?.grid[i][col] == 0) last_col = false;
    }
    if (last_col) {
      return "red_bg";
    }
  }
  if (
    selectedSolution.value?.solution.type == "row" &&
    row == selectedSolution.value?.solution.row
  ) {
    for (let i = 0; i < 9; i++) {
      if (displayed_grid.value?.grid[row][i] == 0) last_row = false;
    }
    if (last_row) {
      return "red_bg";
    }
  }

  if (!last_row || !last_col || !last_bloc || selectedSolution.value?.solution.type == "") {
    if (
      displayed_grid.value?.grid[row][col] &&
      singleNakedColorationPositions.value.find(
        (e) => e.col == col && e.row == row
      )
    ) {
      return "red_bg";
    }
  }
};
const handleGetAllSolutions = async () => {
  if (displayed_grid.value) {
    const resp = await SolverAPI.getAllSolutions(
      new Grid(displayed_grid.value.grid)
    );

    if (resp) {
      solutions.value = resp;
      displayed_grid.value = new Grid(resp[resp.length - 1].grid.grid);
      selectedCase.value = {
        row: resp[resp.length - 1].solution.row,
        col: resp[resp.length - 1].solution.col,
      };
      if (solutions.value.length)
        handleSelectSolutuion(
          solutions.value[solutions.value.length - 1],
          resp.length - 1
        );
    }
  }
};
const handleSetSelectedCase = (row: number, col: number) => {
  selectedCase.value.row = row;
  selectedCase.value.col = col;
};
const handleSelectSolutuion = (sol: Solution, index: number) => {
  selectedSolution.value = sol;
  selectedCase.value = { row: sol.solution.row, col: sol.solution.col };
  displayed_grid.value = new Grid(sol.grid.grid);
  selectedSolutionIndex.value = index;
};
watch(
  selectedSolutionIndex,

  (value) => {
    singleNakedColorationPositions.value = [];
    if (selectedSolution.value?.solution.method == "Single Naked") {
      const grid = displayed_grid.value?.grid;
      const row = selectedSolution.value.solution.row;
      const col = selectedSolution.value.solution.col;
      let copy: number[] = [];
      let copyRef: number[] = [];

      for (let i = 1; i < 10; i++) {
        if (i != selectedSolution.value.solution.value) copyRef.push(i);
      }
      if (selectedSolution.value.solution.type == "") {
        const bloc = getBlocFromCase(row, col);
        for (let j = 1; j < 10; j++) {
          if (j != selectedSolution.value.solution.value) {
            copy.push(j);
          }
        }
        for (let i = 0; i < 9; i++) {
          const rowStart = 3 * Math.floor(bloc / 3);
          const colStart = 3 * (bloc % 3);
          let j = rowStart + Math.floor(i / 3);
          let k = colStart + (i % 3);
          if (
            copy.length &&
            grid &&
            grid[j][k] > 0 &&
            copy.find((e) => e == grid[j][k])
          ) {
            singleNakedColorationPositions.value.push({ row: j, col: k });
            copy = copy.filter((e) => e != grid[j][k]);
          }
        }
        for (let i = 0; i < 9; i++) {
          if (copy.length) {
            if (
              grid &&
              grid[row][i] > 0 &&
              copy.find((e) => e == grid[row][i])
            ) {
              singleNakedColorationPositions.value.push({ row, col: i });
              copy = copy.filter((e) => e != grid[row][i]);
            }
          }
        }
        for (let i = 0; i < 9; i++) {
          if (copy.length) {
            if (
              grid &&
              grid[i][col] > 0 &&
              copy.find((e) => e == grid[i][col])
            ) {
              singleNakedColorationPositions.value.push({ row: i, col });
              copy = copy.filter((e) => e != grid[i][col]);
            }
          }
        }
      }
      for (let i = 0; i < 9; i++) {
        if (selectedSolution.value.solution.type == "col") {
          for (let j = 1; j < 10; j++) {
            if (
              j != selectedSolution.value.solution.value &&
              !isValueInRow(j, selectedSolution.value.solution.row)
            ) {
              copy.push(j);
            }
          }
          if (
            grid &&
            grid[i][col] > 0 &&
            copy.indexOf(grid[i][col]) > -1 &&
            copy.find((e) => e == grid[i][col])
          ) {
            singleNakedColorationPositions.value.push({ row: i, col });
            copy = copy.filter((e) => e != grid[i][col]);
          }
        }

        if (selectedSolution.value.solution.type == "row") {
          for (let k = 1; k < 10; k++) {
            if (
              k != selectedSolution.value.solution.value &&
              !isValueInCol(k, selectedSolution.value.solution.col)
            ) {
              copy.push(k);
            }
          }
          if (
            grid &&
            grid[row][i] > 0 &&
            copy.indexOf(grid[row][i]) > -1 &&
            copy.find((e) => e == grid[row][i])
          ) {
            singleNakedColorationPositions.value.push({ row, col: i });
            copy = copy.filter((e) => e != grid[row][i]);
          }
        }
      }
    }
  }
);
</script>

<style>
.cell {
  border: 1px solid black;
  width: 50px;
  height: 50px;
  text-align: center;
  font-weight: 500;
  font-size: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cell-separator {
  margin-left: 10px;
}
.row-separator {
  margin-top: 10px !important;
}

.selected_case {
  background-color: #d1c4e9 !important;
  border: 2px solid rgba(68, 0, 255, 0.349);
}

.blue_bg {
  background-color: rgba(46, 116, 247, 0.651);
}
.red_bg {
  background-color: rgba(255, 0, 0, 0.349);
}
</style>
