<template>
  <div class="webhook-table">
    <br>
    <h2 class="title">> Webhooks</h2>
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
  color: #00ffff; /* Cyan text color */
  margin-bottom: 20px;
}

.webhook-table {
  font-family: Arial, sans-serif;
}

.webhook-table table {
  width: 100%;
  border-collapse: collapse;
  color: #00ffff; /* Cyan text color */
}

.webhook-table th,
.webhook-table td {
  padding: 5px;
  text-align: left;
  border-bottom: 2px solid #ff00ff; /* Magenta border color */
  font-size: 0.9em; /* Smaller font size */
}

.webhook-table th {
  font-weight: 600;
}

.toggle-btn {
  background-color: #ff4b4b;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

.toggle-btn:hover {
  background-color: #ff7b7b;
  transform: translateY(-2px);
}

/* Responsive Design */
@media screen and (max-width: 768px) {
  .webhook-table td {
    padding: 8px;
    font-size: 0.8em; /* Adjusted for smaller screens */
  }
}

@media screen and (max-width: 480px) {
  .webhook-table td {
    padding: 6px;
    font-size: 0.7em; /* Adjusted for very small screens */
  }
}
</style>