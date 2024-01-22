<script setup>
import { nextTick, onMounted, ref } from "vue";
import { GetFractions } from "../../wailsjs/go/main/App";
import MultiSelect from "primevue/multiselect";

const fractions = ref([]);
const selected = defineModel();

onMounted(async () => {
  nextTick(async () => {
    const r = await GetFractions();
    fractions.value = r;
  });
});
</script>

<template>
  <MultiSelect v-model="selected" :options="fractions" class="w-full" />
</template>