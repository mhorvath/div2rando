<script setup>
import { computed, ref } from "vue";
import ActivityTypeFilter from "./ActivityTypeFilter.vue";
import InputNumber from "primevue/inputnumber";
import Button from "primevue/button";

const model = defineModel();

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

const durationInNanoSeconds = computed(() => {
  return model.value.DURATION
})

</script>

<template>
  <div>
    <ActivityTypeFilter v-model="type" />
    <InputNumber v-model="duration" placeholder="Duration in minutes" />
    <span>{{ $formatNanoseconds(duration*NANOSEC_M) }}</span>
  </div>
</template>