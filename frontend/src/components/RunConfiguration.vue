<script setup>
import { onMounted, ref } from "vue";
import ActivityFilter from "./ActivityFilter.vue";
import Button from "primevue/button";

import { GenerateRun } from "../../wailsjs/go/main/App";

const config = defineModel();

function addActivityFilter() {
  config.value.ACTIVITY_FILTER_OPTIONS.push({
    ACTIVITY_TYPES: [],
    FRACTIONS: [],
    DURATION: 0,
    MULTIPLE_PICK: true,
    MAX_DISTANCE: 0,
  });
}

function removeActivityFilter(index) {
  config.value.ACTIVITY_FILTER_OPTIONS.splice(index, 1);
}

onMounted(() => {
  if (config.value.ACTIVITY_FILTER_OPTIONS.length === 0) {
    addActivityFilter();
  }
});
</script>

<template>
  <div>
    <h1>RunConfigurator</h1>
    <div class="grid column">
      <h2>
        ActivityFilter
        <Button icon="pi pi-plus" text @click="addActivityFilter" />
      </h2>
      <div
        v-for="(c, i) in config.ACTIVITY_FILTER_OPTIONS"
        :key="i"
        class="grid"
      >
        <Button icon="pi pi-minus" text @click="removeActivityFilter(i)" />
        <ActivityFilter v-model="config.ACTIVITY_FILTER_OPTIONS[i]" />
      </div>
    </div>
    <div class="grid">
      <h2>SortActivities</h2>
    </div>
  </div>
</template>

<style scoped>
.grid {
  display: flex;
}
.column {
  flex-direction: column;
}
</style>