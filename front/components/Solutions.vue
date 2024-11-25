<template>
  <v-card class="container my-7 py-5" scrollable>
    <v-virtual-scroll height="360" class="pa-5" :items="[1]">
      <template v-slot:default>
        <v-row dense>
          <v-col cols="6" v-for="sol in Props.solutions">
            <v-card
              style="height: 70px"
              class="solution px-4 py-2"
              @click="$emit('selectSolution', sol)"
              ><h3>{{ sol.solution.method }}_{{ sol.solution.type }}</h3>
              <p>
                <span class="mr-2">ligne {{ sol.solution.row + 1 }}</span
                ><span class="mr-2">ligne {{ sol.solution.col + 1 }}</span
                ><span>valeur {{ sol.solution.value }}</span>
              </p>
            </v-card>
          </v-col>
        </v-row>
      </template>
    </v-virtual-scroll>
    <div class="w-100 d-flex justify-center ga-3 mt-3">
      <v-btn @click="changeSelectedSolution(-1)"
        ><MdiIcon icon="mdiArrowLeft" size="30" />
      </v-btn>
      <v-btn @click="changeSelectedSolution(1)"
        ><MdiIcon icon="mdiArrowRight" size="30"
      /></v-btn>
    </div>
  </v-card>
</template>

<script lang="ts" setup>
import Solution from "~/models/Solution";
const selectedSolution = ref<Solution>();
const Props = defineProps<{
  solutions: Solution[] | [];
  selectedSolutionIndex: number;
  text: String;
}>();

const emit = defineEmits<{
  (e: "selectSolution", solution: Solution): void;
  (e: "changeSelectedSolution", index: number): void;
}>();
const changeSelectedSolution = (sens: number) => {
  if (sens > 0)
    if (Props.selectedSolutionIndex < Props.solutions.length ) {
      emit("changeSelectedSolution", Props.selectedSolutionIndex + 1);
    } else {
      if (Props.selectedSolutionIndex > 0) {
        emit("changeSelectedSolution", Props.selectedSolutionIndex - 1);
      }
    }
};
</script>

<style lang="scss">
.container {
  min-height: 300px;
  width: 80%;
  background-color: #d1c4e9 !important;
}
.solution {
  border: 1px solid #d1c4e9;
}
.selectedSolution{
   background-color: #d1c4e9 !important;
  
}
</style>
