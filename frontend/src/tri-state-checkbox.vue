<template>
    <v-checkbox
        v-model="checked"
        :indeterminate="indeterminate"
        @click="onClick"
    />
</template>

<script lang="ts">
import Component from "vue-class-component";
import Vue from "vue";
import { Prop, Watch } from "vue-property-decorator";

export type CheckBoxValue = "checked" | "indeterminate" | "unchecked";

@Component
export default class TriStateCheckBox extends Vue {
    @Prop({ required: true })
    value!: CheckBoxValue;

    checked = false;

    get indeterminate(): boolean {
        return this.value === "indeterminate";
    }

    private getNextValue(value: CheckBoxValue): CheckBoxValue {
        console.log("--- " + value);
        switch (value) {
            case "unchecked":
            default:
                return "indeterminate";
            case "indeterminate":
                return "checked";
            case "checked":
                return "unchecked";
        }
    }

    mounted() {
        this.onValueChanged(this.value);
    }

    @Watch("value")
    onValueChanged(newValue: CheckBoxValue) {
        this.checked = newValue === "checked";
    }

    // We need to use onClick rather than onChange to cancel bubbling of the click event,
    // which can toggle expansion panels
    onClick(e: MouseEvent) {
        this.$emit("input", this.getNextValue(this.value));
        e.cancelBubble = true;
    }
}
</script>
