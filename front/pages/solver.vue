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
        <v-container>
          <v-row
            v-for="(row, i) in displayed_grid?.grid"
            :class="[i == 3 || i == 6 ? 'row-separator' : '']"
            dense
          >
            <v-col v-for="(value, j) in row">
              <v-card :elevation="getElevation(i,j)"
                :class="[
                  'cell',getBaseRedBg(i, j),
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
      handleSelectSolutuion(solution)
    }
  }
};
const handleChangeSelectedSolutionIndex = (value: number) => {
  selectedSolutionIndex.value = value;
  selectedSolution.value = solutions.value[selectedSolutionIndex.value];
};
const getBlocFromCase = (row: number, col: number) => {
  return 3 * ((row / 3) >> 0) + ((col / 3) >> 0);
};

const getElevation = (row:number,col:number)=>{
  if(selectedSolution.value && selectedSolution.value.solution.row == row && selectedSolution.value.solution.col == col) return 5
  return 0
}

const getBaseRedBg = (row: number, col: number) => {
  if (
    displayed_grid.value?.grid[row][col] ==
    selectedSolution.value?.solution.value
  ) {
    return "red_bg";
  } else {
    if (displayed_grid.value?.grid[row][col] == 0) {
      if (selectedSolution.value?.solution.type == "bloc") {
        if (
          getBlocFromCase(row, col) ==
          getBlocFromCase(
            selectedSolution.value.solution.row,
            selectedSolution.value.solution.col
          )
        ) {
          return "red_bg";
        }
      }
      if (selectedSolution.value?.solution.type == "row") {
        if (row == selectedSolution.value.solution.row) {
          return "red_bg";
        }
      }
      if (selectedSolution.value?.solution.type == "col") {
        if (col == selectedSolution.value.solution.col) {
          return "red_bg";
        }
      }
    }
  }
  return "";
};
const handleGetAllSolutions = async () => {
  if (displayed_grid.value) {
    const resp = await SolverAPI.getAllSolutions(
      new Grid(displayed_grid.value.grid)
    );

    if (resp) {
      solutions.value = resp;
      
      displayed_grid.value = new Grid(resp[resp.length-1].grid.grid);
      selectedCase.value = {
        row: resp[resp.length-1].solution.row,
        col: resp[resp.length-1].solution.col,
      };
      if (solutions.value.length) handleSelectSolutuion(solutions.value[1]);
    }
  }
};
const handleSetSelectedCase = (row: number, col: number) => {
  selectedCase.value.row = row;
  selectedCase.value.col = col;
};
const handleSelectSolutuion = (sol: Solution) => {
  selectedSolution.value = sol;
  selectedCase.value = { row: sol.solution.row, col: sol.solution.col };
};
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
  border: 2px solid  rgba(68, 0, 255, 0.349);
}

.blue_bg {
  background-color: rgba(46, 116, 247, 0.651);
}
.red_bg {
  background-color: rgba(255, 0, 0, 0.349);
}
</style>
