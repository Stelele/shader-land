<template>
    <div>
        <ShaderPlayground :start-code="StartShaderFs" />
    </div>
</template>

<script lang="ts" setup>
import ShaderPlayground from '../components/ShaderPlayground.vue';
import { StartShaderFs } from '../components/Renderer/Start.shader';

</script>