<template>
  <v-card
    class="w-50 d-flex pa-3 flex-column align-center justify-space-between container my-7"
  >
    <div>
      <v-row v-for="i in 3" dense>
        <v-col v-for="j in 3" cols="4" class="d-flex justify-center">
          <v-card
            class="d-flex align-center justify-center key"
            :elevation="props.selected_key == (i - 1) * 3 + j ? 5 : 0"
            :style="{
              backgroundColor:
                selected_key == (i - 1) * 3 + j ? 'rgb(118, 182, 238)' : '',
            }"
            @Click="handleClick((i - 1) * 3 + j)"
          >
            {{ (i - 1) * 3 + j }}
          </v-card>
        </v-col>
      </v-row>
    </div>
    <div>
      <v-btn @click="handleClick(0)" :elevation="props.selected_key == 0 ? 5 : 0"
            :style="{
              backgroundColor:
                selected_key == 0 ? 'rgb(118, 182, 238)' : '',
            }">
        <MdiIcon icon="mdiEraser" size="30" />
    </v-btn>
    </div>
  </v-card>
</template>

<script lang="ts" setup>
const emit = defineEmits<{
  handleClick: [value: number];
}>();

const props = defineProps<{
  selected_key: number;
}>();

const handleClick = (v: number) => {
  emit("handleClick", v);
};
</script>

<style>
.container {
  height: 400px;
  background-color: #d1c4e9;
  &div {
    width: 340px;
  }
}

.key {
  height: 100px;
  font-size: 28px;
  width: 100px;
}
.key:hover {
  cursor: pointer;
  background-color: rgb(118, 182, 238);
}
</style>
