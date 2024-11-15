<template>
  <v-card class="container my-7 py-5" scrollable>
    <v-virtual-scroll height="360" class="pa-5" :items="[1]">
      <template v-slot:default="item">
        <v-row dense>
          <v-col cols="6" v-for="sol in Props.solutions">
            <v-card style="height: 70px" class="solution px-4 py-2"
              ><h3>{{ sol.solution.method }}</h3>
              <p><span class="mr-2">ligne {{ sol.solution.row + 1 }}</span><span class="mr-2">ligne {{ sol.solution.col+1 }}</span><span>valeur {{ sol.solution.value }}</span> </p>
            </v-card>
          </v-col>
        </v-row>
      </template>
    </v-virtual-scroll>
  </v-card>
</template>

<script lang="ts" setup>
import Solution from "~/models/Solution";

const Props = defineProps<{
  solutions: Solution[] | [];
  text: String;
}>();

onMounted(() => {
  console.log(Props);
});
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
</style>
