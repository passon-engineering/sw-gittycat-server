<template>
  <div id="app" class="container">
    <img src="@/assets/logo.png" alt="Logo" class="logo fadein" @click="showMatrix = !showMatrix" />
    <stats-table :title="'Actions Stats'" :stats="actionsStats" @deleteAction="deleteActions" />
    <stats-table :title="'Repositories Stats'" :stats="repositoriesStats" @deleteAction="deleteRepositories" />
    <stats-table :title="'Artifacts Stats'" :stats="artifactsStats" @deleteAction="deleteArtifacts" />
    <webhook-table class="fadein" :webhooks="webhooks" @toggleActive="toggleActive" />
    <action-table class="fadein" :actions="actions" @rerunAction="rerunAction" />
    <matrix-effect v-if="showMatrix" />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import WebhookTable from './components/WebhookTable.vue'
import ActionTable from './components/ActionTable.vue'
import StatsTable from './components/StatsTable.vue'
import MatrixEffect from './components/MatrixEffect.vue'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    WebhookTable,
    ActionTable,
    StatsTable,
    MatrixEffect
  },
  setup() {
    const webhooks = ref({})
    const actions = ref({})
    const showMatrix = ref(false)
    
    const actionsStats = ref([]);
    const repositoriesStats = ref([]);
    const artifactsStats = ref([]);

    onMounted(() => {
      // Fetch data for webhooks and actions
      fetchWebhooks()
      fetchActions()
      // Fetch stats data for each endpoint
      fetchActionsStats()
      fetchRepositoriesStats()
      fetchArtifactsStats()
    });

    const toggleActive = async (repo_name) => {
      await axios.get(`/webhooks/${repo_name}/toggle_active`)
      await axios.get(`/webhooks/refresh`)
      await fetchWebhooks()
    }

    const rerunAction = async (hash) => {
      await axios.get(`/actions/${hash}/rerun`)
      await fetchActions()
    }

    const deleteActions = async (actionEndpoint) => {
      console.log("deleteActions")
      console.log(actionEndpoint)
    }

    const deleteRepositories = async (actionEndpoint) => {
      console.log("deleteRepositories")
      console.log(actionEndpoint)
    }

    const deleteArtifacts = async (actionEndpoint) => {
      console.log("deleteArtifacts")
      console.log(actionEndpoint)
    }

    const fetchWebhooks = async () => {
      const response = await axios.get('/webhooks')
      webhooks.value = response.data
    }

    const fetchActions = async () => {
      const response = await axios.get('/actions')
      actions.value = response.data
    }

    const fetchActionsStats = async () => {
      // Fetch stats data for actions
      const response = await axios.get('/actions/stats');
      actionsStats.value = response.data;
    }

    const fetchRepositoriesStats = async () => {
      // Fetch stats data for repositories
      const response = await axios.get('/repositories/stats');
      repositoriesStats.value = response.data;
    }

    const fetchArtifactsStats = async () => {
      // Fetch stats data for artifacts
      const response = await axios.get('/artifacts/stats');
      artifactsStats.value = response.data;
    }

    return {
      webhooks,
      actions,
      toggleActive,
      rerunAction,
      showMatrix,
      actionsStats,
      repositoriesStats,
      artifactsStats,
      deleteActions,
      deleteRepositories,
      deleteArtifacts
    }
  }
}
</script>