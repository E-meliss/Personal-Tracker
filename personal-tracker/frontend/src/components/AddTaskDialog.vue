<template>
  <Dialog v-model:visible="visibleLocal" modal header="Add Task" :style="{ width: '420px' }">
    <div style="display:flex; flex-direction:column; gap:10px;">
      <InputText v-model="title" placeholder="e.g. Sunscreen" />
      <div class="subtle">Tip: Keep it short and actionable.</div>
    </div>

    <template #footer>
      <Button label="Cancel" severity="secondary" text @click="visibleLocal = false" />
      <Button label="Add" icon="pi pi-check" @click="submit" :disabled="!title.trim()" />
    </template>
  </Dialog>
</template>

<script setup>
import { computed, ref, watch } from 'vue';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';

const props = defineProps({
  visible: { type: Boolean, default: false }
});

const emit = defineEmits(['update:visible', 'add']);

const title = ref('');

const visibleLocal = computed({
  get: () => props.visible,
  set: (v) => emit('update:visible', v)
});

watch(() => props.visible, (v) => {
  if (v) title.value = '';
});

const submit = () => {
  const value = title.value.trim();
  if (!value) return;
  emit('add', value);
  emit('update:visible', false);
};
</script>
