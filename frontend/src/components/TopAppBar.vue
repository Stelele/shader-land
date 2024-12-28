<template>
    <div class="navbar bg-base-100">
        <div class="flex-1 gap-4">
            <RouterLink to="/" class="btn btn-ghost">Shader Land</RouterLink>
            <div class="form-control">
                <input type="text" placeholder="Search" class="input input-bordered w-24 md:w-auto" />
            </div>
        </div>

        <div class="flex-none">
            <ul class="menu menu-horizontal px-1">
                <li>
                    <RouterLink to="/view/1">Browse</RouterLink>
                </li>
                <li>
                    <RouterLink to="/new">New</RouterLink>
                </li>
                <li>
                    <RouterLink to="/login">Sign In</RouterLink>
                </li>
            </ul>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { RouterLink } from 'vue-router';
</script>