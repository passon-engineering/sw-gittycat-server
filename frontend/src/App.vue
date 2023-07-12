<template>
  <div id="app">
    <h1>Active Webhooks</h1>
    <div v-for="webhook in webhooks" :key="webhook.route" class="webhook">
      <h2>{{ webhook.repoName }}</h2>
      <p><strong>Repo URL:</strong> {{ webhook.repoURL }}</p>
      <p><strong>Route:</strong> {{ webhook.route }}</p>
      <ul>
        <li v-for="command in webhook.commands" :key="command">
          {{ command }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      webhooks: [],
    };
  },
  created() {
    this.fetchWebhooks();
  },
  methods: {
    async fetchWebhooks() {
      try {
        const res = await fetch('/webhooks');
        this.webhooks = await res.json();
      } catch (err) {
        console.error(err);
      }
    },
  },
};
</script>

<style scoped>
.webhook {
  background-color: #f9f9f9;
  margin-bottom: 1em;
  padding: 1em;
  border-radius: 4px;
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1);
}
.webhook h2 {
  margin: 0;
  color: #333;
}
.webhook p {
  margin: 0.5em 0;
  color: #666;
}
.webhook ul {
  margin-top: 0.5em;
  padding-left: 1em;
}
</style>