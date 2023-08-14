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
          <th></th>
        </tr>
      </thead>
      <tbody>
        <template v-for="(webhook, webhookKey) in webhooks" :key="webhookKey">
          <tr>
            <td class="cell-large">{{ webhook.build_name }}</td>
            <td class="cell-large">{{ webhook.route }}</td>
            <td class="cell-small" :style="{color: webhook.active ? 'green' : 'red'}"><b>{{ webhook.active ? 'Active' : 'Inactive' }}</b></td>
            <td class="cell-medium cell-padding">
              <button class="btn btn-pink btn-right" @click="toggleRowExpansion(webhookKey)">Toggle Details</button>
              <button class="btn btn-red btn-right" 
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

.cell-large {
  width: 35%;
}

.cell-medium {
  width: 20%;
}
.cell-small {
  width: 10%;
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