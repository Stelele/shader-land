<template>
    <h1>Home Page</h1>
    <div>
        {{ shaders }}
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { ShaderService } from '../services/ShaderService';

onMounted(() => {
    init()
})

const shaders = ref()

async function init() {
    shaders.value = await ShaderService.getShader("test")
}

</script>