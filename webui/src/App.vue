<script setup>
  import {RouterLink, RouterView} from 'vue-router'
</script>

<script>
export default {
	data() {
		return {
			logged: false,
      identifier: null
		}
	},
	methods: {
    check_login(data) {
      this.logged=data
      this.identifier=localStorage.getItem("id")
      this.$router.replace("/home")
    }
	}
}
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <RouterView 
        @user_login="check_login"
        :identifier="identifier"
      />
    </div>
  </div>
</template>

