<template>
  <div class="webhook-table">
    <br>
    <h2 class="title">> Webhooks</h2>
    <table>
      <thead>
        <tr>
          <th>Build</th>
          <th>Route</th>
          <th>Status</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(webhook, webhookKey) in webhooks" :key="webhookKey">
          <tr>
            <td>{{ webhook.build_name }}</td>
            <td>{{ webhook.route }}</td>
            <td :style="{color: webhook.active ? 'green' : 'red'}"><b>{{ webhook.active ? 'Active' : 'Inactive' }}</b></td>
            <td>
              <button class="toggle-expansion" @click="toggleRowExpansion(webhookKey)">Toggle Details</button>
              &nbsp;&nbsp;
              <button class="toggle-active" 
                :style="{ backgroundColor: webhook.active ? 'red' : 'green' }"
                @click="toggleActive(webhook.build_name)">
                {{ webhook.active ? 'Deactivate' : 'Activate' }}
              </button>
            </td>
          </tr>
          <tr v-if="expandedRows[webhookKey]">
            <td colspan="4">
              <div><b>Composer Commands:</b></div>
              <div v-for="command in webhook.composer_commands" :key="command">&nbsp;&nbsp;&nbsp;&nbsp;{{ command }}</div>             
              <div v-for="(repo, repoKey) in webhook.repos" :key="repoKey">
                <br>
                <div><b>Repo URL:</b> {{ repo.repo_url }}</div>
                <div><b>Repo Name:</b> {{ repo.repo_name }}</div>
                <div><b>Branch or Commit Hash:</b> {{ repo.branch_or_commit_hash }}</div>
                <div><b>Inner Repo Commands:</b></div>
                <div v-for="command in repo.inner_repo_commands" :key="command">&nbsp;&nbsp;&nbsp;&nbsp;{{ command }}</div>  
              </div>
            </td>
          </tr>
        </template>
      </tbody>
    </table>
  </div>
</template>

<script>
import { ref, defineComponent } from 'vue'

export default defineComponent({
  name: 'WebhookTable',
  props: ['webhooks'],
  setup(props, { emit }) {
    const expandedRows = ref({})

    const toggleActive = (webhookKey) => {
      emit('toggleActive', webhookKey)
    }

    const toggleRowExpansion = (webhookKey) => {
      expandedRows.value[webhookKey] = !expandedRows.value[webhookKey]
    }

    return {
      toggleActive,
      toggleRowExpansion,
      expandedRows
    }
  }
})
</script>

<style scoped>
.title {
  text-align: center;
  color: #fff;
  margin-bottom: 20px;
}

.webhook-table {
  font-family: Arial, sans-serif;
}

.webhook-table table {
  width: 100%;
  border-collapse: collapse;
  color: #fff; 
}

.webhook-table th,
.webhook-table td {
  padding: 5px;
  text-align: left;
  border-bottom: 2px solid #ffe347; /* border color */
  font-size: 0.9em; 
}

.webhook-table th {
  font-weight: 600;
}

.toggle-expansion {
  font-weight: bold;
  background-color: #247ba0;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

.toggle-expansion:hover {
  transform: translateY(-2px);
}

.toggle-active {
  font-weight: bold;
  background-color: #ff4b4b;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

.toggle-active:hover {
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