<template>
  <v-card class="d-flex flex-column align-center pa-10 w-66" variant="flat">
    <div class="title"><h1>Sudoku Solver</h1></div>
    <p style="text-decoration: underline">About</p>
    <p class="">
      The purpose of this application is to provide a free interface that allows
      users to obtain one or all the steps for solving a Sudoku grid. It is not
      brute force, but rather a step-by-step resolution using classic Sudoku
      techniques.
    </p>
    <v-btn class="mt-4"> SOLVER </v-btn>
  </v-card>
</template>

<script lang="ts" setup></script>

<style>
.title {
  font-size: 30px;
}
p {
  text-align: center;
}
</style>
