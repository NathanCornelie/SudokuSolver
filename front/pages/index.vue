<template>
  <v-card
    variant="flat"
    class="bg-transparent py-2 py-sm-7 d-flex flex-column align-center justify-space-between w-100 h-100"
  >
    <div style="width: fit-content" class="my-2 my-sm-5">
      <div class="mb-5 d-flex align-center justify-center">
        <div v-if="isEditMode" class="w-50 d-flex justify-space-between">
          <v-btn @click="cancelEdit()" color="red"
            ><MdiIcon icon="mdiWindowClose" :size="width>600?'30':'20'" />Cancel</v-btn
          >
          <v-btn @click="saveEdit()"
            ><MdiIcon icon="mdiContentSave" :size="width>600?'30':'20'" />Save</v-btn
          >
        </div>
        <div v-else class="w-100 d-flex align-center justify-space-around">
          <v-btn @Click="goToEdit()">Edit Mode</v-btn>
          <v-btn @Click="resetGrid()">Reset Grid</v-btn>
          <HelpModal v-if="width<=600" :isEdit="isEditMode" />
        </div>
      </div>
      <v-card class="mb-5 d-flex align-center">
        <div id="solutions-container" class="grid">
          <div
            v-for="(row, i) in displayed_grid?.grid"
            :class="[i == 3 || i == 6 ? 'row-separator' : '', 'd-flex']"
            dense
          >
            <div v-for="(value, j) in row" style="padding: 1px">
              <v-card
                rounded=""
                :elevation="getElevation(i, j)"
                :class="[
                  width > 600 ? 'cell' : 'cell_mobile',
                  value.color ,
                  j == 3 || j == 6 ? 'cell-separator' : '',
                  i == selectedCase.row && j == selectedCase.col
                    ? 'selected_case'
                    : '',

                  
                ]"
                @click="handleSetSelectedCase(i, j)"
              >
                <p v-if="value">{{ value.value }}</p>
              </v-card>
            </div>
          </div>
        </div>
      </v-card>
      <div
        class="mt-5 d-flex align-center justify-space-around"
        v-if="!isEditMode"
      >
        <v-btn @click="handleSolve()">View Solution</v-btn>
        <v-btn @click="handleGetNextSolution">Next Solution</v-btn>
        <v-btn @click="handleGetAllSolutions">All Solutions</v-btn>
      </div>
    </div>

    <v-snackbar
      v-model="snackbar"
      color="error"
      variant="tonal"
      timeout="3000"
      location="top left"
    >
      <p class="font-weight-bold">{{ text_snackbar }}</p>

      <template v-slot:actions>
        <v-btn color="pink" variant="text" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>

    <Solutions
      v-if="!isEditMode"
      :solutions="solutions ? solutions : []"
      :selected-solution-index="selectedSolutionIndex"
      text="hello"
      @select-solution="handleSelectSolutuion"
      @change-selected-solution="handleChangeSelectedSolutionIndex"
    />
    <Keyboard
      v-else
      @handle-click="handleKeyboardClick"
      :selected_key="selectedKey"
    />
  </v-card>

  <div v-if="width>600" id="help" class="help_button">
    <HelpModal :isEdit="isEditMode" />
  </div>
</template>

<script lang="ts" setup>
import SolverAPI from "~/api/SolverAPI";
import {Grid} from "~/models/Grid";
import Solution from "~/models/Solution";

const base_grid = ref<Grid>();
const displayed_grid = ref<Grid>();
const before_edit_grid = ref<Grid>();
const solutions = ref<Solution[]>([]);
const selectedSolution = ref<Solution>();
const selectedSolutionIndex = ref<number>(0);

const isEditMode = ref<boolean>(false);
const displayHelp = ref<boolean>(false);
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
const new_grid = ref<number[][]>([
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
  [0, 0, 0, 0, 0, 0, 0, 0, 0],
]);
const width = ref(0);
let height = 0;

onMounted(async () => {
  await nextTick();
  updateSize();
});
const updateSize = () => {
  if (window) {
    width.value = window.innerWidth;
    height = window.innerHeight;
  }
};

onMounted(() => {
  window.addEventListener("resize", updateSize);
});

onUnmounted(() => {
  window.removeEventListener("resize", updateSize);
});
const snackbar = ref<boolean>(false);
const text_snackbar = ref<String>("");
const selectedKey = ref<number>(-1);
onMounted(() => {
 base_grid.value = new Grid(createGridFromBase(data_grid.value).grid);

 displayed_grid.value = base_grid.value;
  before_edit_grid.value = base_grid.value;
  document.addEventListener("keydown", function (e) {
    if (displayed_grid.value)
      if (e.key in ["1", "2", "3", "4", "5", "6", "7", "8", "9"]) {
        displayed_grid.value.grid[selectedCase.value.row][
          selectedCase.value.col
        ].value = +e.key;
      }
    if (e.key == "ArrowRight") {
      if (selectedCase.value.col < 8) selectedCase.value.col++;
    }
    if (e.key == "ArrowUp") {
      if (selectedCase.value.row > 0) selectedCase.value.row--;
    }
    if (e.key == "ArrowLeft") {
      if (selectedCase.value.col > 0) selectedCase.value.col--;
    }
    if (e.key == "ArrowDown") {
      if (selectedCase.value.row < 8) selectedCase.value.row++;
    }
  });
});
const handleSolve = async () => {
  selectedSolution.value = undefined;
  if (data_grid.value) {
    
    const solution = await SolverAPI.solve(new Grid(createGridFromBase(data_grid.value).grid));

    if (solution) {
      displayed_grid.value = new Grid(solution.grid);
    }
  }
};
const createGridFromBase=(grid:number[][])=>{
  let resp = new Grid()
  for (let i = 0; i < 9; i++) {
    for (let j = 0; j < 9; j++) {
      resp.grid[i][j]={value:grid[i][j],color:""}
    }
  }

  return resp
}
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

const resetGrid = () => {
  if (displayed_grid.value) displayed_grid.value.grid = createGridFromBase(data_grid.value).grid;
  selectedSolution.value = undefined;
  solutions.value = [];
};

const goToEdit = () => {
  isEditMode.value = true;
  if (displayed_grid.value && before_edit_grid.value) {
    before_edit_grid.value = { ...displayed_grid.value };
    displayed_grid.value.grid = createGridFromBase(new_grid.value).grid;
  }
};
const cancelEdit = () => {
  isEditMode.value = false;
  new_grid.value = [
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
    [0, 0, 0, 0, 0, 0, 0, 0, 0],
  ];
  if (before_edit_grid.value) displayed_grid.value = before_edit_grid.value;
};
const saveEdit = async () => {
  if (displayed_grid.value) {
    const resp = await SolverAPI.checkGrid(displayed_grid.value);
    if (resp.verified) {
      isEditMode.value = false;
      solutions.value = [];
    } else {
      snackbar.value = true;
      text_snackbar.value = resp.message;
    }
  }
};


const handleChangeSelectedSolutionIndex = (value: number) => {
  handleSelectSolutuion(solutions.value[value], value);
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


const handleGetAllSolutions = async () => {
  if (displayed_grid.value) {
    const resp = await SolverAPI.getAllSolutions(
      new Grid(displayed_grid.value.grid)
    );

    if (resp?.length) {
      solutions.value = [...solutions.value,...resp];
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
  selectedSolution.value = undefined;
  selectedCase.value.row = row;
  selectedCase.value.col = col;
  if (isEditMode.value && displayed_grid.value && selectedKey.value > -1) {
    displayed_grid.value.grid[selectedCase.value.row][selectedCase.value.col].value =
      selectedKey.value;
  }
};
const handleSelectSolutuion = (sol: Solution, index: number) => {
  selectedSolution.value = sol;
  selectedCase.value = { row: sol.solution.row, col: sol.solution.col };
  displayed_grid.value = new Grid(sol.grid.grid);
  selectedSolutionIndex.value = index;
};
const handleKeyboardClick = (v: number) => {
  if (isEditMode && selectedCase && displayed_grid.value) {
    if (selectedKey.value == v) selectedKey.value = -1;
    else selectedKey.value = v;
  }
};

</script>

<style scoped>
.grid {
  max-width: 90%;
  padding: 15px;
}
.cell {
  border: 2px solid rgb(167, 167, 167);
  width: 50px;
  height: 50px;
  text-align: center;
  font-size: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.col {
  .v-col {
    flex-grow: 0 !important;
  }
}
@media (max-width: 600px) {
  .v-btn {
    font-size: 10px !important;
    padding: 10px 5px;
    margin: 0 5px;
  }
  .grid {
    max-width: 90%;
    padding: 15px;
  }
  .cell-separator {
    margin-left: 5px !important;
  }
  .row-separator {
    margin-top: 5px !important;
  }
}

.cell_mobile {
  border: 2px solid rgb(167, 167, 167);
  width: 35px;
  height: 35px;

  text-align: center;
  font-weight: 300;
  font-size: 26px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cell-separator {
  margin-left: 10px;
}
.row-separator {
  margin-top: 10px;
}

.selected_case {
  background-color: #d1c4e9 !important;
  border: 2px solid rgba(68, 0, 255, 0.349);
}

.blue_bg {
  background-color: rgba(46, 116, 247, 0.651);
}
.red2{
  background-color: rgba(255, 0, 0, 0.349);
}
.red1{
  background-color: rgba(255, 0, 0, 0.589);

}
.help_button {
  position: absolute;

  top: 20px;
  right: 40px;
}
</style>
