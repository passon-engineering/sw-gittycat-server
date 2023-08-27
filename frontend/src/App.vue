<template>
  <div id="app" class="container">
    <img src="@/assets/logo.png" alt="Logo" class="logo fadein" @click="showMatrix = !showMatrix" />
    <stats-table :title="'Actions'" :stats="actionsStats" @deleteAction="deleteActions" />
    <stats-table :title="'Repositories'" :stats="repositoriesStats" @deleteAction="deleteRepositories" />
    <stats-table :title="'Artifacts'" :stats="artifactsStats" @deleteAction="deleteArtifacts" />
    <webhook-table class="fadein" :webhooks="webhooks" @toggleActive="toggleActive" />
    <action-table class="fadein" :actions="actions" @rerunAction="rerunAction" />
    <div v-if="totalPages > 1" class="pagination">
      <button class="btn" @click="changePage(currentPage - 1)" :disabled="currentPage === 1">Previous</button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
      <button class="btn" @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages">Next</button>
    </div>
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
    const currentPage = ref(1);
    const totalPages = ref(0);
    const actionsStats = ref([]);
    const repositoriesStats = ref([]);
    const artifactsStats = ref([]);

    onMounted(() => {
      fetchWebhooks()
      fetchActions()
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

    const deleteActions = async () => {
      await axios.get(`/actions/delete`)
      await fetchActions()
      await fetchActionsStats()
    }

    const deleteRepositories = async () => {
      await axios.get(`/repositories/delete`)
      await fetchRepositoriesStats()
    }

    const deleteArtifacts = async () => {
      await axios.get(`/artifacts/delete`)
      await fetchArtifactsStats()
    }

    const fetchWebhooks = async () => {
      const response = await axios.get('/webhooks')
      webhooks.value = response.data
    }

    const fetchActions = async (page = 1) => {
      const response = await axios.get(`/actions/${page}`);
      actions.value = response.data.data;
      totalPages.value = response.data.totalPages;
      currentPage.value = page;
    }

    const fetchActionsStats = async () => {
      const response = await axios.get('/actions/stats');
      actionsStats.value = response.data;
    }

    const fetchRepositoriesStats = async () => {
      const response = await axios.get('/repositories/stats');
      repositoriesStats.value = response.data;
    }

    const fetchArtifactsStats = async () => {
      const response = await axios.get('/artifacts/stats');
      artifactsStats.value = response.data;
    }

    const changePage = (page) => {
      if (page >= 1 && page <= totalPages.value) {
          fetchActions(page);
      }
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
      deleteArtifacts,
      changePage,
      currentPage,
      totalPages
    }
  }
}
</script>

<style>
body {
  background-color: #1f1f1f;
  color: #FFFFFF;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 15px;
  box-sizing: border-box;
  font-family: Arial, sans-serif;
  position: relative;
}

.logo {
  display: block;
  width: 200px; 
  height: auto; 
  margin: 0 auto 20px;
}

.fadein {
  animation: fadeIn 2s;
}

.btn {
  font-weight: bold;
  color: #fff;
  border: none;
  border-radius: 20px;
  padding: 5px 10px;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
}

.btn:hover {
  box-shadow: 0 0 5px rgba(255, 255, 255, 0.6);
}

.btn-right{
  float: right;
  margin-right: 5px;
}

.btn-green {background-color: green;}
.btn-pink {background-color: #ff00ff;}
.btn-mustard {background-color: #ffe347;}
.btn-red {background-color: red;}

.pagination {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 0;
}

@keyframes fadeIn {
  0% {opacity: 0;}
  100% {opacity: 1;}
}
</style>