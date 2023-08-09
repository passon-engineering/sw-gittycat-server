<template>
  <div class="stats-table">
    <table>
      <tr class="data-row">
        <td class="data-item title">
          <div class="data-value">{{ title }}</div>
        </td>
        <td class="data-item file-count">
          <div class="data-label">File Count:</div>
          <div class="data-value">{{ stats.file_count }}</div>
        </td>
        <td class="data-item directory-count">
          <div class="data-label">Directory Count:</div>
          <div class="data-value">{{ stats.directory_count }}</div>
        </td>
        <td class="data-item total-size">
          <div class="data-label">Total Size:</div>
          <div class="data-value">{{ formatSize(stats.total_size) }}</div>
        </td>
        <td class="data-item delete">
          <button class="delete-button" @click="deleteAction(stats.actionEndpoint)">Delete</button>
        </td>
      </tr>
    </table>
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
        return 'N/A';
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
.stats-table {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 5px 0; /* Adds vertical spacing between components */
}

table {
  width: 100%;
  border-collapse: collapse;
}

.data-row {
  display: flex;
  background-color: #2e2e2e;
  border: 1px solid #444;
  border-radius: 5px;
  padding: 10px;
  color: #ffffff;
  margin: 3px 0; /* Adds vertical spacing between rows */
}

.data-item {
  display: flex;
  align-items: center;
}

.data-label,
.data-value {
  font-size: 0.9em; /* Makes the font a little smaller */
  white-space: nowrap;
}

.data-label {
  margin-right: 5px;
}

/* Explicit widths for alignment */
.data-item.title {
  width: 25%;
}

.data-item.file-count,
.data-item.directory-count,
.data-item.total-size {
  width: 20%;
}

.data-item.delete {
  width: 15%;
  justify-content: flex-end;
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
