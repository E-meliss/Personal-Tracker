<template>
  <div>
    <div v-if="loading" class="subtle">Loading...</div>

    <div v-else-if="items.length === 0" class="empty glass">
      <i class="pi pi-sparkles" style="font-size: 18px;"></i>
      <div>
        <div style="font-weight: 700;">No tasks yet</div>
        <div class="subtle">Add a task to get started.</div>
      </div>
    </div>

    <div v-else class="list">
      <div v-for="t in items" :key="t.id" class="row glass">
        <Checkbox :binary="true" :modelValue="t.isDone" @update:modelValue="() => $emit('toggle', t.id)" />
        <div class="title" :class="{ done: t.isDone }">{{ t.title }}</div>
        <Button icon="pi pi-trash" severity="secondary" text @click="$emit('remove', t.id)" />
      </div>
    </div>
  </div>
</template>

<script setup>
import Button from 'primevue/button';
import Checkbox from 'primevue/checkbox';

defineProps({
  items: { type: Array, required: true },
  loading: { type: Boolean, default: false }
});

defineEmits(['toggle', 'remove']);
</script>

<style scoped>
.list{ display:flex; flex-direction:column; gap:10px; }
.row{
  display:flex;
  align-items:center;
  gap:12px;
  padding: 12px 12px;
}
.title{
  flex:1;
  font-weight: 600;
}
.done{
  opacity:.6;
  text-decoration: line-through;
}
.empty{
  display:flex;
  align-items:center;
  gap:12px;
  padding: 14px;
}
</style>
