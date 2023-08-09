<template>
  <div class="data-row">
    <div class="data-title">{{ title }}</div>
    <div class="data-item">
      <div class="data-label">File Count</div>
      <div class="data-value">{{ stats.file_count }}</div>
    </div>
    <div class="data-item">
      <div class="data-label">Directory Count</div>
      <div class="data-value">{{ stats.directory_count }}</div>
    </div>
    <div class="data-item">
      <div class="data-label">Total Size</div>
      <div class="data-value">{{ formatSize(stats.total_size) }}</div>
    </div>
    <button class="delete-button" @click="deleteAction(stat.actionEndpoint)">Delete</button>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'StatsTable',
  props: {
    title: String,
    stats: Array,
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
      this.$emit('deleteAction', actionEndpoint)
    },
  },
})
</script>

<style scoped>
.data-row {
  display: flex;
  justify-content: space-between;
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

.data-item {
  text-align: center;
}

.data-label {
  font-weight: bold;
  margin-bottom: 5px;
  color: #ffffff;
}

.data-value {
  font-size: 1.2em;
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