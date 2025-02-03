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
    <div class="row position-absolute top-50 start-50 translate-middle">
      <div class="col-12 d-flex justify-content-center">
        <div class="card" style="width: 18rem;" id="crd">
            <div class="card-body">
              <div class="input-group flex-nowrap">
                <span class="input-group-text" id="addon-wrapping">@</span>
                  <input type="text" class="form-control" placeholder="Username" aria-label="Username" aria-describedby="addon-wrapping" v-model="username">
              </div>
              <a href="./src/views/Page.html" class="btn btn-primary" id="btn" @click="login">Accedi</a>
            </div>
        </div>
      </div>
      
    </div>
  </div>
</template>

<style>
  #btn{
    margin-top: 10%;
  }

</style>

