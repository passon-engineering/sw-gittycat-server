<template>
  <div id="app" class="container">
    <img src="@/assets/logo.png" alt="Logo" class="logo fadein" @click="showMatrix = !showMatrix" />
    <webhook-table class="fadein" :webhooks="webhooks" @toggleActive="toggleActive" />
    <action-table class="fadein" :actions="actions" @rerunAction="rerunAction" />
    <matrix-effect v-if="showMatrix" />
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import WebhookTable from './components/WebhookTable.vue'
import ActionTable from './components/ActionTable.vue'
import MatrixEffect from './components/MatrixEffect.vue'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    WebhookTable,
    ActionTable,
    MatrixEffect
  },
  setup() {
    const webhooks = ref({})
    const actions = ref({})
    const showMatrix = ref(false)

    const fetchWebhooks = async () => {
      const response = await axios.get('/webhooks')
      webhooks.value = response.data
    }

    const fetchActions = async () => {
      const response = await axios.get('/actions')
      actions.value = response.data
    }

    onMounted(() => {
      fetchWebhooks()
      fetchActions()
    })

    const toggleActive = async (repo_name) => {
      await axios.get(`/webhooks/${repo_name}/toggle_active`)
      await axios.get(`/webhooks/refresh`)
      await fetchWebhooks()
    }

    const rerunAction = async (repo_name) => {
      await axios.get(`/actions/${repo_name}/rerun`)
      await fetchActions()
    }

    return {
      webhooks,
      actions,
      toggleActive,
      rerunAction,
      showMatrix
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

@keyframes fadeIn {
  0% {opacity: 0;}
  100% {opacity: 1;}
}
</style>
