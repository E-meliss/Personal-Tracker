<template>
  <div class="container">
    <div class="top glass">
      <div class="left">
        <div class="h1">Personal Tracker</div>
        <div class="subtle">Daily checklist • clean & simple</div>
      </div>

      <div class="right">
        <DatePicker v-model="dateObj" dateFormat="yy-mm-dd" showIcon iconDisplay="input" />
        <Button label="Add" icon="pi pi-plus" @click="showAdd = true" />
      </div>
    </div>

    <div class="grid">
      <div class="card glass">
        <div class="card-header">
          <div>
            <div class="h1">Today</div>
            <div class="subtle">{{ store.day }}</div>
          </div>

          <div class="progress-wrap">
            <Knob :modelValue="store.progress" readonly :size="76" />
            <div class="subtle">{{ store.completed }} / {{ store.total }}</div>
          </div>
        </div>

        <ProgressBar :value="store.progress" :showValue="false" style="height: 10px; border-radius: 999px;" />

        <div style="height: 14px;"></div>

        <TaskList
          :items="store.items"
          :loading="store.loading"
          @toggle="store.toggleTask"
          @remove="store.deleteTask"
        />
      </div>

      <div class="side glass">
        <div class="h1">Tips</div>
        <div class="subtle" style="margin-top: 8px; line-height: 1.6;">
          • 3-6 task ideal<br />
          • “Small wins” mantığı<br />
          • Aynı güne odaklan
        </div>

        <Divider />

        <div class="h1">Quick presets</div>
        <div class="preset">
          <Button label="Skincare" severity="secondary" @click="addPreset('Skincare')" />
          <Button label="Meals" severity="secondary" @click="addPreset('Meals')" />
          <Button label="Workout" severity="secondary" @click="addPreset('Workout')" />
        </div>
      </div>
    </div>

    <AddTaskDialog v-model:visible="showAdd" @add="onAdd" />
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from 'vue';
import { useTasksStore } from '../stores/tasks';

import Button from 'primevue/button';
import DatePicker from 'primevue/datepicker';
import ProgressBar from 'primevue/progressbar';
import Knob from 'primevue/knob';
import Divider from 'primevue/divider';

import TaskList from '../components/TaskList.vue';
import AddTaskDialog from '../components/AddTaskDialog.vue';

const store = useTasksStore();

const showAdd = ref(false);
const dateObj = ref(new Date(store.day));

watch(dateObj, async (val) => {
  const day = val.toISOString().slice(0, 10);
  store.setDay(day);
  await store.fetchTasks();
});

onMounted(async () => {
  await store.fetchTasks();
});

const onAdd = async (title) => {
  await store.addTask(title);
};

const addPreset = async (type) => {
  const presets = {
    Skincare: ['Cleanser', 'Moisturizer', 'Sunscreen'],
    Meals: ['Breakfast', 'Lunch', 'Dinner'],
    Workout: ['Walk 20 min', 'Stretch 10 min', 'Water 2L']
  };

  const list = presets[type] || [];
  for (const item of list) {
    await store.addTask(item);
  }
};
</script>

<style scoped>
.top{
  display:flex;
  align-items:center;
  justify-content:space-between;
  padding:18px 16px;
  gap:14px;
}
.right{
  display:flex;
  gap:10px;
  align-items:center;
}
.grid{
  margin-top: 16px;
  display:grid;
  grid-template-columns: 1.6fr .9fr;
  gap: 14px;
}
.card{
  padding:16px;
}
.side{
  padding:16px;
}
.card-header{
  display:flex;
  align-items:center;
  justify-content:space-between;
  gap:12px;
  margin-bottom: 10px;
}
.progress-wrap{
  display:flex;
  align-items:center;
  gap:10px;
}
.preset{
  margin-top: 10px;
  display:flex;
  flex-wrap: wrap;
  gap: 10px;
}
@media (max-width: 900px){
  .grid{ grid-template-columns: 1fr; }
}
</style>
