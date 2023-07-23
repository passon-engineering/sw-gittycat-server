<template>
  <div class="action-table">
    <h1 class="title">Webhook Actions</h1>
    <table>
      <thead>
        <tr>
          <th>Repository Name</th>
          <th>Status</th>
          <th>Last Call</th>
          <th>Processing Time</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(action, key) in actions" :key="key">
          <td>{{ action.webhook.repo_name }}</td>
          <td :style="{color: action.success ? 'green' : 'red'}"><b>{{ action.success ? 'Success' : 'Failure' }}</b></td>
          <td>{{ action.last_call }}</td>
          <td>{{ action.processing_time }}</td>
          <td>
            <button class="rerun-btn" @click="rerunAction(action.webhook.repo_name)">
              Re-run
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
  name: 'ActionTable',
  props: ['actions'],
  setup(props, { emit }) {
    const rerunAction = (repo_name) => {
      emit('rerunAction', repo_name)
    }

    return {
      rerunAction
    }
  }
})
</script>

<style scoped>
.title {
  text-align: center;
  color: #ffffff;
  margin-bottom: 20px;
}

.action-table table {
  width: 100%;
  border-collapse: collapse;
  color: #ffffff;
}

.action-table th,
.action-table td {
  border: 1px solid #ddd;
  background-color: #2e2e2e;
  padding: 8px;
  text-align: center;
}

.action-table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #444;
  color: white;
}

.rerun-btn {
  padding: 5px 10px;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 5px;
  background-color: #555;
}

.rerun-btn:hover {
  background-color: #777;
}
</style>
