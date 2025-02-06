<template>
    <div class="prose">
        <h1>Home</h1>
    </div>
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
    shaders.value = await ShaderService.getShaders()
}

</script>