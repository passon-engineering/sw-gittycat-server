<template>
  <div class="stats-table">
    <table>
      <tr>
        <td class="cell-title">{{ title }}</td>
        <td class="cell-file-count">File Count: {{ stats.file_count }}</td>
        <td class="cell-directory-count">Directory Count: {{ stats.directory_count}}</td>
        <td class="cell-total-size">Total Size: {{ formatSize(stats.total_size) }}</td>
        <td class="cell-delete">
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
.stats-table table {
  width: 100%;
  border-collapse: collapse;
  color: #ffffff;
}

.stats-table th,
.stats-table td {
  border: 1px solid #ddd;
  background-color: #2e2e2e;
  padding: 8px;
  text-align: center;
}

.stats-table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #444;
  color: white;
}

/* Specific cell widths */
.cell-title {
  width: 25%;
}

.cell-file-count {
  width: 20%;
}

.cell-directory-count {
  width: 20%;
}

.cell-total-size {
  width: 20%;
}

.cell-delete {
  width: 15%;
}
</style>

