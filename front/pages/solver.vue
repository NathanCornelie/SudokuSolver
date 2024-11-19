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
              <div
                :class="[
                  'cell',
                  j == 3 || j == 6 ? 'cell-separator' : '',
                  i == selectedCase.row && j == selectedCase.col
                    ? 'selected_case'
                    : '',
                ]"

                @click="handleSetSelectedCase(i,j)"
              >
                <p>{{ value }}</p>
              </div>
            </v-col>
          </v-row>
        </v-container>
      </v-card>
      <div class="mt-5 d-flex align-center justify-space-around">
        <v-btn @click="handleSolve()">View Solution</v-btn>
        <v-btn @click="handleGetOneSolution">Next Solution</v-btn>
        <v-btn @click="handleGetAllSolution">All Solutions</v-btn>
      </div>
    </div>

    <Solutions :solutions="solutions ? solutions : []" text="hello" />
  </v-card>
</template>

<script lang="ts" setup>
import SolverAPI from "~/api/SolverAPI";
import Grid from "~/models/Grid";
import Solution from "~/models/Solution";

const base_grid = ref<Grid>();
const displayed_grid = ref<Grid>();
const solutions = ref<Solution[]>([]);
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
const handleGetOneSolution = async () => {
  if (displayed_grid.value) {
    const solution = await SolverAPI.getOneSolution(
      new Grid(displayed_grid.value.grid)
    );
    console.log(solution);
    if (solution) {
      solutions.value?.push(solution);
      console.log(solution.grid);
      displayed_grid.value = new Grid(solution.grid.grid);
      selectedCase.value = {
        row: solution.solution.row,
        col: solution.solution.col,
      };
    }
  }
};const handleGetAllSolution = async () => {
  if (displayed_grid.value) {
    const resp = await SolverAPI.getAllSolutions(
      new Grid(displayed_grid.value.grid)
    );
    
    if (resp) {
      solutions.value=resp; 
      displayed_grid.value = new Grid(resp[-1].grid.grid);
      selectedCase.value = {
        row: resp[-1].solution.row,
        col: resp[-1].solution.col,
      };
    }
  }
};
const handleSetSelectedCase =(row :number,col: number)=>{
  selectedCase.value.row = row
  selectedCase.value.col = col
}
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
  background-color: #d1c4e9;
  border: 2px solid #7e57c2;
}
</style>
