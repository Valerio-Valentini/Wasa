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
        },

        async uploadFile() {
            let fileInput = document.getElementById("fileUploader")
            console.log(fileInput)
            const file = fileInput.files[0]
            const reader = new FileReader()
            reader.readAsArrayBuffer(file)
            reader.onload = async () => {
                await this.$axios.put("/users/" + this.identifier + "/photo", reader.result, {
                    headers: {
                        "Content-Type": file.type
                    }
                })
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
            <button type="button" class="btn btn-primary" @click="changeUsername">Update Username
            </button>
        </div>

        <div class="row">
            <div class="input-group flex-nowrap">
                <input id="fileUploader" type="file" class="profile-file-upload" accept=".jpg">
                <label class="btn" @click="uploadFile">
                    Update Photo
                </label>
            </div>    
        </div>
    </div>
</template>

<style>
</style>