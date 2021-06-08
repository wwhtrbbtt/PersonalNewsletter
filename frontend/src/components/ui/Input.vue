<template>
    <div>
        <label class="field field_v1">
            <input class="field__input" 
                :placeholder="placeholder" 
                :type="type" 
                :value="value" 
                @input="onInput"    
            /><br><br>
                <span class="field__label-wrap">
                    <span class="field__label">{{ text }}</span>
                </span>
        </label>
    </div>
</template>

<script>
export default {
    name: "Input",
    props: {
        value: {
            type: String,
            required: true,
        },
        placeholder: {
            type: String,
            required: false,
        },
        type: {
            type: String,
            required: true,
        },
        text: {
            type: String,
            required: true,
        }
    },
    methods: {
        onInput(event) {
            // console.log(event.target.value)
            this.$emit("input", event.target.value)
        }
    }
}
</script>
