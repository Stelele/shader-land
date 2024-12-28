<template>
    <LogIn v-if="isLogin" />
    <SignUp v-else />
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import LogIn from '../components/LogIn.vue';
import SignUp from '../components/SignUp.vue';

const isLogin = ref(true)
</script>