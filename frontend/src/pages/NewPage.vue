<template>
    <div class="grid grid-rows-2">
        <div>
            <ShaderPlayground :start-code="StartShaderFs">
                <SubmitShaderDetails v-if="isAuthenticated" />
            </ShaderPlayground>
        </div>
        <div class="grid grid-cols-3">

        </div>
    </div>
</template>

<script lang="ts" setup>
import ShaderPlayground from '../components/ShaderPlayground.vue';
import SubmitShaderDetails from '../components/SubmitShaderDetails.vue';
import { StartShaderFs } from '../components/Renderer/Start.shader';
import { useAuth0 } from '@auth0/auth0-vue';

const { isAuthenticated } = useAuth0()

</script>