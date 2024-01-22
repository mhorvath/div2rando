<script setup>
import { computed, ref } from "vue";
import ActivityTypeFilter from "./ActivityTypeFilter.vue";
import ActivityFractionFilter from "./ActivityFractionFilter.vue";
import InputNumber from "primevue/inputnumber";
import Button from "primevue/button";
import SelectButton from "primevue/selectbutton";

const model = defineModel();
const multiplePickOptions = ref([
  { label: "SINGLE", value: false },
  { label: "MULTIPLE", value: true },
]);
const NANOSEC_M = 1000 * 1000 * 1000;

const type = computed({
  get() {
    return model.value.ACTIVITY_TYPES;
  },
  set(value) {
    model.value.ACTIVITY_TYPES = value;
  },
});

const duration = computed({
  get() {
    return model.value.DURATION / NANOSEC_M;
  },
  set(value) {
    model.value.DURATION = value * NANOSEC_M;
  },
});

const fraction = computed({
  get() {
    return model.value.FRACTIONS;
  },
  set(value) {
    model.value.FRACTIONS = value;
  },
});

const multiplePick = computed({
  get() {
    return model.value.MULTIPLE_PICK;
  },
  set(value) {
    model.value.MULTIPLE_PICK = value;
  },
});

const durationInNanoSeconds = computed(() => {
  return model.value.DURATION;
});
</script>

<template>
  <div>
    <ActivityTypeFilter v-model="type" />
    <ActivityFractionFilter v-model="fraction" />
    <InputNumber v-model="duration" placeholder="Duration in minutes" />
    <span>{{ $formatNanoseconds(duration * NANOSEC_M) }}</span>
    <SelectButton v-model="multiplePick" :options="multiplePickOptions" optionLabel="label" optionValue="value" />
  </div>
</template>