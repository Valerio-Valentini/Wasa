<script>
export default {
    data: function () {
        return {
            username : null,
        }
    },

    methods: {
        async changeUsername() {
            try {
                    await this.$axios.put("/users/" + this.identifier, {
                        user_id: this.username.trim()
                        })
                    localStorage.setItem("id", this.username)
                    this.$emit("user_login", true)
                    this.$router.push("/home")
                }
            catch (error) {
                console.log(error)
            }
        }
    },

    props: ["identifier"]
}
</script>

<template>
    <div class="container-fluid">
        <div class="row">
            <div class="input-group flex-nowrap">
                <span class="input-group-text" id="addon-wrapping">@</span>
                <input type="text" class="form-control" placeholder="Username" aria-label="Username" aria-describedby="addon-wrapping"
                v-model="username">
                </div>    
        </div>
        <div class="row">
            <button type="button" class="btn btn-primary" @click="changeUsername">Primary</button>
        </div>
    </div>
</template>

<style>
</style>