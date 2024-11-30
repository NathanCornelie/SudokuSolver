<template>
  <v-card class="container my-7 py-5" style="" scrollable>
    <div
      style="height: 360px"
      class="pa-5 overflow-auto"
      :items="[1]"
      id="solutions-container"
      ref="container-ref"
    >
      <v-row dense>
        <v-col cols="6" v-for="(sol, i) in Props.solutions">
          <v-card
            :ref="(ref) => (cardsRef[i] = ref)"
            :id="`h-${i}`"
            :elevation="selectedSolutionIndex == i ? 5 : 0"
            style="height: 70px"
            :class="[
              'solution',
              ' px-4',
              ' py-2',
              selectedSolutionIndex == i ? 'selectedSolution' : '',
            ]"
            @click="$emit('selectSolution', sol, i)"
            ><h3>{{ sol.solution.method }}_{{ sol.solution.type }}</h3>
            <p>
              <span class="mr-2">ligne {{ sol.solution.row + 1 }}</span
              ><span class="mr-2">ligne {{ sol.solution.col + 1 }}</span
              ><span>valeur {{ sol.solution.value }}</span>
            </p>
          </v-card>
        </v-col>
      </v-row>
    </div>

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
import { useGoTo } from "vuetify";

const selectedSolution = ref<Solution>();
const goTo = useGoTo();
const Props = defineProps<{
  solutions: Solution[] | [];
  selectedSolutionIndex: number;
  text: String;
}>();
const cardsRef: HTMLElement[] = [];
const containerRef = useTemplateRef("container-ref");
const emit = defineEmits<{
  (e: "selectSolution", solution: Solution, index: number): void;
  (e: "changeSelectedSolution", index: number): void;
}>();
const changeSelectedSolution = (sens: number) => {
  if (sens > 0) {
    if (Props.selectedSolutionIndex < Props.solutions.length - 1) {
      emit("changeSelectedSolution", Props.selectedSolutionIndex + 1);
    }
    0;
  } else {
    if (Props.selectedSolutionIndex > 0) {
      emit("changeSelectedSolution", Props.selectedSolutionIndex - 1);
    }
  }
};

watch(
  () => Props.selectedSolutionIndex,
  (value) => {
    if(value == Props.solutions.length-1){
      setTimeout(function(){
      goTo(cardsRef[value]||`#h-${value}`, {
      container: containerRef.value || "#solutions-container",
      offset: -110,
      easing: "easeInQuint",
    });
}, 500);
    }else{
    
    goTo(cardsRef[value]||`#h-${value}`, {
      container: containerRef.value || "#solutions-container",
      offset: -110,
      easing: "easeInQuint",
    });
  }

  }
);
</script>

<style lang="postcss">
.container {
  min-height: 300px;
  width: 80%;
  background-color: #d1c4e9 !important;
}
.solution {
  border: 1px solid #d1c4e9;
}
.selectedSolution {
  background-color: #ffffff71 !important;
}
</style>
