<template>
    <v-dialog v-model="show" :max-width="640">
        <v-card>
            <v-card-title>
                New Task
            </v-card-title>
            <v-card-text>
                <v-container>
                    <v-row>
                        <v-col cols="12">
                            <v-text-field
                                v-model="name"
                                label="Task"
                                prepend-icon="mdi-calendar-check"
                            />
                        </v-col>
                        <v-col cols="12">
                            <v-text-field
                                v-model="desc"
                                label="Description"
                                prepend-icon="mdi-newspaper"
                            />
                        </v-col>
                        <v-col cols="12">
                            <v-menu
                                ref="menu"
                                v-model="menu"
                                :close-on-content-click="false"
                                :return-value.sync="date"
                                transition="scale-transition"
                                offset-y
                                min-width="290px"
                            >
                                <template v-slot:activator="{ on }">
                                    <v-text-field
                                        v-model="date"
                                        label="Due date"
                                        prepend-icon="mdi-calendar"
                                        readonly
                                        v-on="on"
                                    ></v-text-field>
                                </template>
                                <v-date-picker v-model="date" no-title scrollable>
                                    <v-spacer />
                                    <v-btn
                                        text
                                        color="primary"
                                        @click="menu = false"
                                    >
                                        Cancel
                                    </v-btn>
                                    <v-btn
                                        text
                                        color="primary"
                                        @click="$refs.menu.save(date)"
                                    >
                                        OK
                                    </v-btn>
                                </v-date-picker>
                            </v-menu>
                        </v-col>
                    </v-row>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer />
                <v-btn color="green darken-1" dark @click="onOK">OK</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import { Prop, Watch } from "vue-property-decorator";
import restApi from "@/rest-api";
import { Task } from "@/app.d";

@Component({
    components: {},
})
export default class AddTaskDialog extends Vue {
    @Prop({ required: true })
    value!: boolean;

    name = "";
    desc = "";
    date: string = new Date().toDateString();
    menu = false;

    get show(): boolean {
        return this.value;
    }

    set show(show: boolean) {
        this.$emit("input", show);
    }

    onOK() {
        restApi.createTask(this.name, this.desc).then((task: Task) => {
            console.log(JSON.stringify(task));
            this.$store.commit("addTask", task);
        });

        this.show = false;
    }

    @Watch("value")
    onValueChanged(newValue: boolean) {
        if (newValue) {
            this.name = this.desc = "";
            this.date = new Date().toDateString();
        }
    }
}
</script>
