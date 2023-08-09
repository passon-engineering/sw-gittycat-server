<template>
  <div class="data-row">
    <div class="data-title">{{ title }}</div>
    <div class="data-label-value">
      <div class="data-item">
        <span class="data-label">File Count:</span>
        <span class="data-value">{{ stats.file_count }}</span>
      </div>
      <div class="data-item">
        <span class="data-label">Directory Count:</span>
        <span class="data-value">{{ stats.directory_count }}</span>
      </div>
      <div class="data-item">
        <span class="data-label">Total Size:</span>
        <span class="data-value">{{ formatSize(stats.total_size) }}</span>
      </div>
    </div>
    <button class="delete-button" @click="deleteAction(stats.actionEndpoint)">Delete</button>
  </div>
</template>

<script>
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'StatsTable',
  props: {
    title: String,
    stats: Object,
  },
  methods: {
    formatSize(sizeInBytes) {
      if (sizeInBytes === undefined) {
        return 'N/A'; // Return a placeholder or appropriate value for undefined size
      }

      const units = ['B', 'KB', 'MB', 'GB', 'TB'];
      let size = sizeInBytes;
      let unitIndex = 0;

      while (size >= 1024 && unitIndex < units.length - 1) {
        size /= 1024;
        unitIndex++;
      }

      return `${size.toFixed(2)} ${units[unitIndex]}`;
    },
    deleteAction(actionEndpoint) {
      this.$emit('deleteAction', actionEndpoint);
    },
  },
});
</script>

<style scoped>
.data-row {
  display: flex;
  align-items: center;
  background-color: #2e2e2e;
  padding: 10px;
  border: 1px solid #444;
  border-radius: 5px;
  margin-bottom: 10px;
}

.data-title {
  font-weight: bold;
  width: 150px; /* Adjust the width as needed */
  color: #ffffff;
}

.data-label-value {
  display: flex;
  gap: 20px;
  align-items: center;
}

.data-item {
  display: flex;
  align-items: center;
}

.data-label {
  margin-right: 5px;
  color: #ffffff;
}

.data-value {
  margin-right: 5px;
  color: #ffffff;
}

.delete-button {
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 5px;
  padding: 5px 10px;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out;
}

.delete-button:hover {
  background-color: #d32f2f;
}
</style>
