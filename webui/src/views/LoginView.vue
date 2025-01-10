<script>
export default {
	data: function() {
		return {
			username: null
		}
	},
	methods: {
        async login () {
            try {
                let response = await this.$axios.post("/session", {
                    user_id: this.username.trim()
                })

                console.log(response)
                
            } catch (error) {
                console.log(error)
                
            }
        }
		
	
	}
}
</script>

<template>
	<div class="container-fluid">
        <div class="row">
            <div class="col">
                <h1 class="text-center">
                    Login Page
                </h1>
            </div>
        </div>
        <div class="row">
            <div class="d-flex justify-content-center">
                <div class="col-5">
                    <div class="input-group mb-3">
                        <input type="text" class="form-control" placeholder="Insert Username" v-model="username">
                            <button class="btn btn-outline-secondary" type="button" id="button-addon2" @click="login">Login</button >
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
</style>