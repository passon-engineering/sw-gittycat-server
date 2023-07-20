<template>
  <div class="webhook-table">
    <h1 class="title">Webhooks</h1>
    <table>
      <thead>
        <tr>
          <th>Repository URL</th>
          <th>Repository Name</th>
          <th>Route</th>
          <th>Status</th>
          <th>Commands</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(webhook, key) in webhooks" :key="key">
          <td>{{ webhook.repo_url }}</td>
          <td>{{ webhook.repo_name }}</td>
          <td>{{ webhook.route }}</td>
          <td :style="{color: webhook.active ? 'green' : 'red'}"><b>{{ webhook.active ? 'Active' : 'Inactive' }}</b></td>
          <td>{{ webhook.commands.join(', ') }}</td>
          <td>
            <button class="toggle-btn" @click="toggleActive(webhook.repo_name)">
              {{ webhook.active ? 'Deactivate' : 'Activate' }}
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
  name: 'WebhookTable',
  props: ['webhooks'],
  setup(props, { emit }) {
    const toggleActive = (repo_name) => {
      emit('toggleActive', repo_name)
    }

    return {
      toggleActive
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

.webhook-table table {
  width: 100%;
  border-collapse: collapse;
  color: #ffffff;
}

.webhook-table th,
.webhook-table td {
  border: 1px solid #ddd;
  background-color: #2e2e2e;
  padding: 8px;
  text-align: center;
}

.webhook-table th {
  padding-top: 12px;
  padding-bottom: 12px;
  text-align: left;
  background-color: #444;
  color: white;
}

.toggle-btn {
  padding: 5px 10px;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 5px;
  background-color: #555;
}

.toggle-btn:hover {
  background-color: #777;
}
</style>