<template>
    <div>
        <ShaderPlayground :start-code="StartShaderFs" />
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import ShaderPlayground from '../components/ShaderPlayground.vue';
import { StartShaderFs } from '../components/Renderer/Start.shader';

const route = useRoute()
const id = ref(route.params.id)
</script>