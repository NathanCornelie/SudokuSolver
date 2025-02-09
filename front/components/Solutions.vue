<template>
  <v-card class="container my-sm-7 py-sm-2 d-flex flex-column justify-end" style="" scrollable>
    <div
     
      class="pa-2 overflow-auto"
      :items="[1]"
      id="solutions-container"
      ref="container-ref"
    >
      <v-row dense>
        <v-col :cols="width>600?6:12" v-for="(sol, i) in Props.solutions">
          <v-card
            :ref="(ref) => (cardsRef[i] = ref)"
            :id="`h-${i}`"
            :elevation="selectedSolutionIndex == i ? 5 : 0"
           
            :class="[
              'solution',
              ' px-4',
              ' py-2',
              selectedSolutionIndex == i ? 'selectedSolution' : '',
            ]"
            @click="$emit('selectSolution', sol, i)"
            ><h3>{{ sol.solution.method }} <span class="text-grey"> {{ sol.solution.type ? `(${sol.solution.type})`:'' }}</span></h3>
            <p class="font-weight-bold" >
              <span class="mr-2" >row <span class="font-weight-regular">{{ sol.solution.row + 1 }}</span></span
              ><span class="mr-2">col <span class="font-weight-regular">{{ sol.solution.col + 1 }}</span></span
              ><span>value <span class="font-weight-regular">{{ sol.solution.value }}</span></span>
            </p>
          </v-card>
        </v-col>
      </v-row>
    </div>

    <div v-if="width>600" class="w-100 d-flex justify-center ga-3 mt-3">
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

<style scopped >


@media (max-width: 600px) {
h3{
  font-size:16px;
}
}

.container {
  height: 100%;
  width: 80%;
  margin-bottom: 20px;
  background-color: #d1c4e9 !important;
}
.solution {
  border: 1px solid #d1c4e9;
}
.selectedSolution {
  background-color: #ffffff71 !important;
}
</style>
