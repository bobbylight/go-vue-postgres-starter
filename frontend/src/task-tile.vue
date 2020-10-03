<template>
    <v-expansion-panel class="task">
        <v-expansion-panel-header
            @mouseover="armed = true"
            @mouseleave="armed = false"
        >
            <template v-slot:default="{ open }">
                <v-row no-gutters>
                    <v-col cols="9" class="headline">
                        <tri-state-checkbox
                            class="task-cb"
                            v-model="state"
                            @input="updateTaskState"
                        />
                        <span
                            v-if="!editingLabel"
                            @click="startEditingLabel"
                            class="editable-title"
                            :class="labelClass"
                        >
                            {{ task.label }}
                        </span>
                        <span v-if="editingLabel">
                            <v-text-field
                                class="d-inline-block headline editable-title"
                                label="Label"
                                v-model="newLabel"
                                @keypress.stop
                                @keydown.esc="cancelUpdateLabel"
                                @keypress.enter="updateLabel"
                                @blur="cancelUpdateLabel"
                                ref="editLabelField"
                            />
                        </span>
                    </v-col>

                    <v-col cols="3">
                        <v-btn-toggle group class="float-right" v-if="armed">
                            <v-btn @click.stop="deleteTask">
                                <v-icon>mdi-delete</v-icon>
                            </v-btn>
                            <v-btn @click.stop="moveUp">
                                <v-icon>mdi-arrow-up</v-icon>
                            </v-btn>
                            <v-btn @click.stop="moveDown">
                                <v-icon>mdi-arrow-down</v-icon>
                            </v-btn>
                        </v-btn-toggle>
                    </v-col>

                    <v-col cols="12" v-if="task.desc">
                        {{ task.desc }}
                    </v-col>
                </v-row>
            </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
            <v-row no-gutters>
                <v-col cols="12" v-if="task.desc">{{ task.desc }}</v-col>
                <v-col cols="6">Started: {{ task.createdAt | date }}</v-col>
                <v-col cols="6">Completed: {{ task.completedAt | date }}</v-col>
            </v-row>
        </v-expansion-panel-content>
    </v-expansion-panel>
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { Task, TaskStatus } from "./app";
import { Prop, Watch } from "vue-property-decorator";
import TriStateCheckbox, { CheckBoxValue } from "@/tri-state-checkbox.vue";
import { dateFilter } from "@/filters";

const TASK_STATE_TO_CHECKBOX_STATE: { [key: string]: CheckBoxValue } = {
    notStarted: "unchecked",
    inProgress: "indeterminate",
    completed: "checked",
};

const CHECKBOX_STATE_TO_TASK_STATE: { [key: string]: TaskStatus } = {
    unchecked: "notStarted",
    indeterminate: "inProgress",
    checked: "completed",
};

type LabelClass = "unchecked" | "indeterminate" | "checked";

@Component({
    components: { TriStateCheckbox },
    filters: { date: dateFilter },
})
export default class TaskTile extends Vue {
    @Prop({ required: true })
    task!: Task;

    state: CheckBoxValue = "unchecked";
    labelClass: LabelClass = "unchecked";
    editingLabel = false;
    newLabel = "";
    armed = false;

    cancelUpdateLabel() {
        this.editingLabel = false;
    }

    deleteTask() {
        this.$store.dispatch("deleteTask", this.task);
    }

    mounted() {
        // Initial value - task never changes so this is OK
        this.state = TASK_STATE_TO_CHECKBOX_STATE[this.task.status];
    }

    moveDown() {
        this.$store.dispatch("reorderTask", {
            task: this.task,
            forward: true,
        });
    }

    moveUp() {
        this.$store.dispatch("reorderTask", {
            task: this.task,
            forward: false,
        });
    }

    @Watch("state")
    onStateChanged(newValue: CheckBoxValue, oldValue: CheckBoxValue) {
        if (newValue && newValue !== oldValue) {
            this.labelClass = newValue;
            //this.state = CHECKBOX_STATE_TO_TASK_STATE[newValue];
        }
    }

    startEditingLabel() {
        this.newLabel = this.task.label;
        this.editingLabel = true;
        this.$nextTick(() => {
            const control = this.$refs.editLabelField as Vue;
            const elem = control.$el as HTMLElement;
            const input: HTMLInputElement | null = elem.querySelector("input");
            if (input) {
                input.focus();
                input.select();
            }
        });
    }

    updateLabel() {
        const newTask: Task = {
            completedAt: this.task.completedAt,
            createdAt: this.task.createdAt,
            desc: this.task.desc,
            id: this.task.id,
            index: -1,
            label: this.newLabel,
            status: this.task.status,
        };

        this.$store.dispatch("updateTask", newTask);
        this.editingLabel = false;
    }

    updateTaskState() {
        const newStatus: TaskStatus = CHECKBOX_STATE_TO_TASK_STATE[this.state];
        const completedAt: string | undefined =
            newStatus === "completed" ? new Date().toISOString() : undefined;
        const newTask: Task = {
            completedAt: completedAt,
            createdAt: this.task.createdAt,
            desc: this.task.desc,
            id: this.task.id,
            index: this.task.index,
            label: this.task.label,
            status: newStatus,
        };

        this.$store.dispatch("updateTaskState", newTask);
        this.editingLabel = false;
    }
}
</script>

<style lang="less">
.task {
    .task-cb {
        display: inline-block;
    }

    .checked {
        color: darkgrey;
        text-decoration: line-through;
    }

    .editable-title {
        display: inline-block;
        min-width: 20rem;
        min-height: 1rem;
    }
}
</style>
