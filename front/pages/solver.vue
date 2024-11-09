<template>
  <v-card variant="flat" class="bg-transparent pa-5">
    <div class="mb-5 d-flex align-center justify-space-around">
      <v-btn>Edit grid</v-btn>
    </div>
    <v-card>
      <v-container>
        <v-row
          v-for="(row, j) in grid"
          :class="[j == 3 || j == 6 ? 'row-separator' : '']"
          dense
        >
          <v-col v-for="(value, i) in row">
            <div :class="['cell', i == 3 || i == 6 ? 'cell-separator' : '']">
              <p>{{ value }}</p>
            </div>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
    <div class="mt-5 d-flex align-center justify-space-around">
      <v-btn>View Solution</v-btn>
      <v-btn>Next Solution</v-btn>
    </div>
  </v-card>
</template>

<script lang="ts" setup>
const grid = ref<number[][]>([]);
onMounted(() => {
  let row: number[] = [];
  for (let i = 0; i < 9; i++) {
    row = [];
    for (let j = 1; j <= 9; j++) {
      row.push(j);
    }
    grid.value.push(row);
  }
  console.log(grid);
});
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
</style>
