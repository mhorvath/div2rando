<script setup>
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import ColumnGroup from "primevue/columngroup";
import Row from "primevue/row";
import Button from "primevue/button";
import { computed, inject } from "vue";

const data = defineModel();
const $formatNanoseconds = inject("formatNanoseconds");

const emit = defineEmits(["generateRun", "saveRun"]);

const runDuration = computed(() => {
  const totalMilliseconds = data.value.reduce(
    (sum, activity) => sum + activity.duration,
    0
  );
  return $formatNanoseconds(totalMilliseconds);
});

function generateRun() {
  emit("generateRun");
}
function saveRun() {
  emit("saveRun")
}
</script>

<template>
  <DataTable :value="data">
    <template #header>
      <div class="grid">
        <span class="text-xl text-900 font-bold">Run </span>
        <span class="text-300">duration: {{ runDuration }}</span>
        <div>
          <Button icon="pi pi-save" text @click="saveRun" />
          <Button
            @click="generateRun"
            label="RRGeeeezuz"
            icon="pi pi-refresh"
          />
        </div>
      </div>
    </template>
    <Column field="name" header="Name" />
    <Column field="type" header="Type" />
    <Column field="fraction" header="Fraction" />
    <Column field="duration" header="Estimated duration">
      <template #body="slotProps">
        {{ $formatNanoseconds(slotProps.data.duration) }}
      </template>
    </Column>
  </DataTable>
</template>

<style scoped>
.grid {
  display: flex;
}
.grid > *:last-child {
  margin-left: auto;
  margin-right: 0;
}
</style>