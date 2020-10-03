<template>
    <div>
        <v-expansion-panels class="task-list">
            <task-tile v-for="task in tasks" :task="task" :key="task.id" />
        </v-expansion-panels>
        <add-task-dialog v-model="addTaskDialogVisible" />

        <v-fab-transition>
            <v-btn
                class="add-button"
                color="pink"
                dark
                fixed
                top
                right
                fab
                @click="addTaskDialogVisible = true"
            >
                <v-icon>mdi-plus</v-icon>
            </v-btn>
        </v-fab-transition>
    </div>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { Task } from "./app";
import { Prop } from "vue-property-decorator";
import AddTaskDialog from "./add-task-dialog.vue";
import TaskTile from "./task-tile.vue";

@Component({ components: { AddTaskDialog, TaskTile } })
export default class TaskList extends Vue {
    @Prop({ required: true })
    tasks!: Task[];

    addTaskDialogVisible = false;
}
</script>
