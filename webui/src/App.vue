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
    },

    check_logout() {
      this.identifier=null
      localStorage.removeItem("id");
      this.$router.replace("/login")
    },
	},
  mounted(){
    let id = localStorage.getItem("id")
    if (id) { 
      this.identifier = id; 
    }else{
      this.identifier = null;
    }
  }
}
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <RouterView 
        @user_login="check_login"
        @logout="check_logout"
        :identifier="identifier"
      />
    </div>
  </div>
</template>

