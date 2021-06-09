<template>
        <!-- <label class="field field_v1">
            <input class="field__input" 
                :placeholder="placeholder" 
                :type="type" 
                :value="value" 
                @input="onInput"    
            /><br><br>
                <span class="field__label-wrap">
                    <span class="field__label">{{ text }}</span>
                </span>
        </label> -->
    <div class="dropdown">
        <link href='https://css.gg/chevron-double-down.css' rel='stylesheet'>
        <button class="dropbtn">{{ title }}
            <svg
  width="17"
  height="17"
  viewBox="0 -3 24 24"
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
>
  <path
    d="M7.75735 5.63605L6.34314 7.05026L12 12.7071L17.6569 7.05029L16.2427 5.63608L12 9.87872L7.75735 5.63605Z"
    fill="currentColor"
  />
  <path
    d="M6.34314 12.7071L7.75735 11.2929L12 15.5356L16.2427 11.2929L17.6569 12.7071L12 18.364L6.34314 12.7071Z"
    fill="currentColor"
  />
</svg>
        </button>
        <div class="dropdown-content">
            <a v-for="(option, key) in options" v-bind:key="key" v-on:click="chosen=option.Key">{{ option.Value }}</a>
        </div>
    </div>
</template>

<script>
export default {
    name: "Dropdown",
    props: {
        title: {
            required: true,
            type: String,
        },
        options: {
            required: true,
            type: Array,
        }
    },
    methods: {
        emit(val) {
            // console.log(event.target.value)
            this.$emit("input", val)
        }
    },
    watch: {
        chosen(val) {
            this.emit(val)
        },
    },
    data: function() {
        return {
            chosen: "",
        }
    }
}
</script>

<style scoped>
.dropbtn {
  color: white;
  padding: 16px;
  font-size: 16px;
  border: 0.2rem solid #ddbea9;
  background-color: #cb997e;
  cursor: pointer;
  min-width: 140px;
}

.dropdown {
  position: relative;
  display: inline-block;
}

.dropdown-content {
  display: none;
  position: absolute;
  background-color: var(--nav1, #fff);
  color: var(--text-dark, #000);
  min-width: 140px;
  cursor: pointer;
}

.dropdown-content a {
  color: black;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
}

/* .dropdown-content a:hover {

} */

.dropdown:hover .dropdown-content {
  display: block;
}

/* .dropdown:hover .dropbtn {
  border: 0.2rem solid #cb997e;
  background-color: #ddbea9;
} */
</style>
