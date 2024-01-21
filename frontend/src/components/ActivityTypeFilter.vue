<script setup>
import { onMounted, ref } from "vue";
import { GetActivityTypes } from "../../wailsjs/go/main/App";
import MultiSelect from "primevue/multiselect";

const activityTypes = ref([]);
const selected = defineModel();

onMounted(async () => {
  const r = await GetActivityTypes();
  activityTypes.value = r;
});
</script>

<template>
  <MultiSelect
    v-model="selected"
    :options="activityTypes"
    class="w-full"
  />
</template>