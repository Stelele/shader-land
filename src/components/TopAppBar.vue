<template>
    <div class="navbar bg-base-100">
        <div class="flex-1">
            <RouterLink to="/" class="btn btn-ghost text-xl">daisyUI</RouterLink>
        </div>
        <div class="flex-none">
            <ul class="menu menu-horizontal px-1">
                <li>
                    <RouterLink to="/view/1">Link</RouterLink>
                </li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
</script>