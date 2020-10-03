<template>
    <v-container>
        <v-row class="text-center">
            <v-col class="mb-4">
                <h1 class="display-2 font-weight-bold mb-3">
                    Tasks for {{ date | date }}
                </h1>
            </v-col>
        </v-row>
        <task-list :tasks="$store.state.tasks" />
    </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import restApi from "@/rest-api";
import { DataPage, Task } from "@/app";
import TaskList from "@/task-list.vue";
import { dateFilter } from "@/filters";

@Component({
    components: { TaskList },
    filters: { date: dateFilter },
})
export default class MainContent extends Vue {
    tasks: Task[] = [];

    date: Date = new Date();

    mounted() {
        restApi.getTasks().then((dataPage: DataPage<Task>) => {
            this.$store.commit("setTasks", dataPage.data);
        });
    }
}
</script>

<style lang="less">
.v-btn.add-button {
    top: 5rem !important;
    right: 3rem !important;
    z-index: 100;
}
</style>
