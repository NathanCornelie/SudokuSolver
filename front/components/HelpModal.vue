<template>
  <div class="d-flex flex-column align-end">
    <v-btn @click="display = !display">
      <MdiIcon icon="mdiHelp" size="20" />
    </v-btn>
    <v-card class="mt-2" v-if="display" max-width="300px">
      <v-card-text>
        <div v-if="props.isEdit" class="help_edit">
          <p class="py-2 "><span><MdiIcon icon="mdiHelp" class="mr-1"/></span>Select one of the numbers on the numeric keypad and click on the cell to enter the value.</p>
          <p class="py-2 "><span><MdiIcon icon="mdiHelp" class="mr-1"/></span>To deselect a value, click on it again.</p>
          <p class="py-2 "><span><MdiIcon icon="mdiHelp" class="mr-1"/></span>You can use the keyboard to move through the grid and enter values.</p>
        </div>
        
        <v-list lines="one" v-else>
          <v-list-item
            v-for="(n, i) in content"
            :key="i"
            :title="n.title"
            :subtitle="n.content"
            ><template v-slot:prepend>
              <MdiIcon
                icon="mdiHelp"
                class="mr-3 text-grey-lighten-1"
                size="20"
              />
            </template>
          </v-list-item>
        </v-list>
      </v-card-text>
    </v-card>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps<{ isEdit: boolean }>();
const display = ref<boolean>(false);
const content = ref<{ title: string; content: string }[]>([
  {
    title: "Edit Mode",
    content: "Go into edit mode to enter a custom grid.",
  },
  {
    title: "Reset Grid",
    content: "Return to the default grid.",
  },
  {
    title: "View Solution",
    content: "Display the solution to the grid.",
  },
  {
    title: "Next Solution",
    content: "Retrieve the next move to play.",
  },
  {
    title: "All Solutions",
    content: "Retrieve all moves to complete the grid.",
  },
]);
</script>

<style >
.tooltip {
  background-color: red;
}

</style>
