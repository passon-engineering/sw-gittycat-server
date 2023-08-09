<template>
  <div class="stats-table">
    <table>
      <tr>
        <td class="title">{{ title }}</td>
        <td class="data-item data-cell">File Count: {{ stats.file_count }}</td>     
        <td class="data-item data-cell">Directory Count: {{ stats.directory_count }}</td>
        <td class="data-item data-cell">Total Size: {{ formatSize(stats.total_size) }}</td>        
        <td class="data-cell"><button class="delete-button" @click="deleteAction(stats.actionEndpoint)">Delete</button></td>
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
.stats-table {
  display: flex;
  justify-content: center;
  align-items: center;
}

table {
  width: 100%;
  border-collapse: collapse;
}

.title {
  font-weight: bold;
  width: 150px;
  padding-right: 10px;
  color: #ffffff;
}

.data-item {
  margin-right: 5px;
  color: #ffffff;
}

.data-cell {
  width: 25%; /* Adjust the width as needed */
  padding: 5px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
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
