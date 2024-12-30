<template>
    <div class="card card-compact bg-base-100 shadow-xl">
        <div class="card-body">
            <div class="flex flex-col gap-4">
                <label class="input input-bordered flex items-center">
                    <input type="text" class="grow" placeholder="Name of your shader" />
                </label>
                <label class="input input-bordered flex items-center">
                    <input type="text" class="grow" placeholder="Tags, comma separated" />
                </label>
                <label class="flex items-center">
                    <textarea class="grow textarea textarea-bordered" placeholder="Describe your shader"></textarea>
                </label>
                <button class="btn btn-primary">Submit</button>
            </div>
        </div>
    </div>
</template>