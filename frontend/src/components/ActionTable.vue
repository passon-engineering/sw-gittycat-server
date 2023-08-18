<template>
  <div class="action-table">
    <br>
    <h2 class="title">> Actions</h2>
    <table>
      <thead>
        <tr>
          <th @click="sortBy('webhook.build_name')">Build Name</th>
          <th @click="sortBy('success')">Status</th>
          <th @click="sortBy('last_call')">Calling Timestamp</th>
          <th @click="sortBy('processing_time')">Processing Time</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(action, key) in sortedActions" :key="key">
          <tr>  
            <td>{{ action.webhook.build_name }}</td>
            <td :style="{color: action.success ? 'green' : 'red'}"><b>{{ action.success ? 'Success' : 'Failure' }}</b></td>
            <td>{{ action.last_call }}</td>
            <td>{{ action.processing_time }}</td>
            <td>
              <button class="btn btn-pink btn-right" @click="toggleRowExpansion(key)">Toggle Details</button>
              <button class="btn btn-red btn-right" @click="rerunAction(action.webhook.build_name)">Re-run</button>
            </td>
          </tr>
          <tr v-if="expandedRows[key]">
            <td colspan="4">
              <div>{{ action.output }}</div>
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'ActionTable',
  props: ['actions'],
  data() {
    return {
      sortOrder: -1, // 1 for ascending, -1 for descending
      sortByColumn: 'last_call', // Default sorting column
      expandedRows: {}, // Add this line to store the state of expanded/collapsed rows
    }
  },
  computed: {
    sortedActions() {
      const sorted = Object.keys(this.actions).sort((a, b) => {
        const actionA = this.actions[a][this.sortByColumn];
        const actionB = this.actions[b][this.sortByColumn];

        if (this.sortByColumn) {
          return actionA < actionB ? -this.sortOrder : this.sortOrder;
        }
        return 0;
      });

      return sorted.map(key => this.actions[key]);
    },
  },
  methods: {
    rerunAction(build_name) {
      this.$emit('rerunAction', build_name)
    },
    sortBy(column) {
      if (this.sortByColumn === column) {
        this.sortOrder = -this.sortOrder
      } else {
        this.sortOrder = 1
      }
      this.sortByColumn = column
    },
    toggleRowExpansion(key) {
      // Toggle the state of a given row specified by 'key'
      this.expandedRows[key] = !this.expandedRows[key];
      
      // Assign a new object to force Vue to pick up the change
      this.expandedRows = { ...this.expandedRows };
    }
  },
})
</script>

<style scoped>
.title {
  text-align: center;
  color: #fff;
  margin-bottom: 20px;
}

.action-table {
  font-family: Arial, sans-serif;
}

.action-table table {
  width: 100%;
  border-collapse: collapse;
  color: #fff;
}

.action-table th,
.action-table td {
  padding: 5px;
  text-align: left;
  border-bottom: 2px solid #ffe347;
  font-size: 0.9em; /* Smaller font size */
}

.action-table th {
  font-weight: 600;
}

/* Responsive Design */
@media screen and (max-width: 768px) {
  .action-table td {
    padding: 8px;
    font-size: 0.8em; /* Adjusted for smaller screens */
  }
}

@media screen and (max-width: 480px) {
  .action-table td {
    padding: 6px;
    font-size: 0.7em; /* Adjusted for very small screens */
  }
}
</style>