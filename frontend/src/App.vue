<script setup>
import { GenerateRun } from "../wailsjs/go/main/App";
import { ref } from "vue";
import RunConfiguration from "./components/RunConfiguration.vue";
import ActivityTable from "./components/ActivityTable.vue";
import { SaveRun } from "../wailsjs/go/main/App";
import Toast from "primevue/toast";
import { useToast } from "primevue/usetoast";
const toast = useToast();

const run = ref({ ACTIVITIES: [] });
const runConfig = ref({
  ACTIVITY_FILTER_OPTIONS: [],
});

async function onRunGenerate() {
  const r = await GenerateRun(runConfig.value);
  run.value = r;
}

async function onSaveRun() {
  await SaveRun(run.value);
  toast.add({ severity: 'success', summary: 'Run saved', detail: 'Run saved to disk', life: 3000 });
}
</script>

<template>
  <main>
    <Toast />
    <RunConfiguration v-model="runConfig" />
    <ActivityTable
      v-model="run.ACTIVITIES"
      @generate-run="onRunGenerate"
      @save-run="onSaveRun"
    />
  </main>
</template>
