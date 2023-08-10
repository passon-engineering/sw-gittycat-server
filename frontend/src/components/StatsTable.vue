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
.stats-table {
  font-family: Arial, sans-serif;
}

.stats-table table {
  width: 100%;
  border-collapse: collapse;
  color: #00ffff; /* Cyan text color */
}

.stats-table td {
  padding: 5px;
  text-align: left;
  border-bottom: 2px solid #ff00ff; /* Magenta border color */
  font-size: 0.9em; /* Smaller font size */
}

.delete-button:hover {
  background-color: #ff7b7b;
  transform: translateY(-2px);
}

/* Specific cell widths */
.cell-title {
  width: 25%;
  font-weight: 600;
}

.cell-file-count,
.cell-directory-count,
.cell-total-size {
  width: 20%;
}

.cell-delete {
  width: 15%;
  padding-right: 5px; /* 5 pixels space from the right */
}

.delete-button {
  float: right; /* This will align the button to the right */
  background-color: #ff4b4b;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}


/* Responsive Design */
@media screen and (max-width: 768px) {
  .stats-table td {
    padding: 8px;
    font-size: 0.8em; /* Adjusted for smaller screens */
  }
}

@media screen and (max-width: 480px) {
  .stats-table td {
    padding: 6px;
    font-size: 0.7em; /* Adjusted for very small screens */
  }
}
</style>
