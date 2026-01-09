import { defineStore } from 'pinia';
import { api } from '../services/api';

export const useTasksStore = defineStore('tasks', {
  state: () => ({
    day: new Date().toISOString().slice(0, 10),
    items: [],
    loading: false
  }),
  getters: {
    total: (state) => state.items.length,
    completed: (state) => state.items.filter((t) => t.isDone).length,
    progress: (state) => {
      const total = state.items.length;
      if (total === 0) return 0;
      const done = state.items.filter((t) => t.isDone).length;
      return Math.round((done / total) * 100);
    }
  },
  actions: {
    setDay(dayStr) {
      this.day = dayStr;
    },
    async fetchTasks() {
      this.loading = true;
      try {
        const { data } = await api.get('/tasks', { params: { day: this.day } });
        this.items = data;
      } finally {
        this.loading = false;
      }
    },
    async addTask(title) {
      await api.post('/tasks', { title, day: this.day });
      await this.fetchTasks();
    },
    async toggleTask(id) {
      await api.patch(`/tasks/${id}/toggle`);
      await this.fetchTasks();
    },
    async deleteTask(id) {
      await api.delete(`/tasks/${id}`);
      await this.fetchTasks();
    }
  }
});
