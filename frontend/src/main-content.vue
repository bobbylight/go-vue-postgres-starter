<template>
    <v-container>
        <v-row class="text-center">
            <v-col class="mb-4">
                <h1 class="display-2 font-weight-bold mb-3">
                    Welcome to Your Go/Vuetify/Postgres Application
                </h1>

                <p class="subheading font-weight-regular">
                    Click here to make a REST call
                </p>
                <v-btn @click="getItems()">Click Me</v-btn>
            </v-col>
        </v-row>

        <v-row
            v-for="widget in widgets"
            :key="widget.id"
            class="justify-center"
        >
            <v-col cols="8">
                <v-card>
                    <v-card-title>
                        {{ widget.name }}
                    </v-card-title>
                    <v-card-text>Price: ${{ widget.price }}</v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import restApi from "@/rest-api";
import { Widget } from "@/app";

@Component
export default class MainContent extends Vue {
    widgets: Widget[] = [];

    getItems() {
        restApi.getWidget("111").then((widget: Widget) => {
            this.widgets.push(widget);
        });
    }
}
</script>
