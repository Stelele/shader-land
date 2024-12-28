<template>
  <nav>
    <TopAppBar />
  </nav>
  <RouterView />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router';
import TopAppBar from './components/TopAppBar.vue';
</script>
