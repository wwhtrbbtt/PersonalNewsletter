<template>
<div class="wrapper">
  <nav class="topnav">
        <router-link v-bind:class="{active: currentRouteName == 'home'}" to="/">Home</router-link>
        <router-link v-bind:class="{active: currentRouteName == 'dashboard'}" to="/dashboard/">Dashboard</router-link>

  </nav>
</div>
</template>

<script>

export default {

computed: {
    currentRouteName() {
        return this.$route.name;
    }
}
}
</script>

<style scoped>

template {
  margin: 0;
  font-family: Arial, Helvetica, sans-serif;
}

.topnav {
  overflow: hidden;
  background-color: #333;
}

.topnav a {
  float: left;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

.topnav a:hover {
  background-color: #ddd;
  color: black;
}

.topnav a.active {
  background-color: #04AA6D;
  color: white;
}

</style>