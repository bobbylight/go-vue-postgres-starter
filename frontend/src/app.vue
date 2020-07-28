<template>
    <v-app>
        <v-app-bar app color="primary" dark>
            <div class="d-flex align-center">
                <h2>Daily Duties</h2>
            </div>

            <v-spacer />

            <v-btn
                href="https://github.com/vuetifyjs/vuetify/releases/latest"
                target="_blank"
                text
            >
                <span class="mr-2">Latest Release</span>
                <v-icon>mdi-open-in-new</v-icon>
            </v-btn>
        </v-app-bar>

        <v-content>
            <main-content />
        </v-content>
    </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import Component from "vue-class-component";
import MainContent from "@/main-content.vue";

@Component({ components: { MainContent } })
export default class App extends Vue {}
</script>
