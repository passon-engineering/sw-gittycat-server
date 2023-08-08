<template>
  <div class="stats-table">
    <h1 class="title">{{ title }}</h1>
    <table>
      <thead>
        <tr>
          <th>Statistic</th>
          <th>Value</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(stat, key) in stats" :key="key">
          <td>{{ stat.label }}</td>
          <td>{{ stat.value }}</td>
          <td>
            <button class="delete-btn" @click="deleteAction(stat.actionEndpoint)">
              Delete
            </button>
          </td>
        </tr>
      </tbody>
    </table>
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
    deleteAction(actionEndpoint) {
      this.$emit('deleteAction', actionEndpoint)
    },
  },
})
</script>

<style scoped>
.title {
  text-align: center;
  color: #ffffff;
  margin-bottom: 20px;
}

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

.delete-btn {
  padding: 5px 10px;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 5px;
  background-color: #ff4444;
}

.delete-btn:hover {
  background-color: #ff6666;
}
</style>
