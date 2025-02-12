<script>
export default {
    data() {
        return {
            searchUser: null,
            foundUsers2: [],
            toAddToGroup: [],
            groupName: ""
        };
    },

    methods: {
        goHome() {
            this.$router.replace("/");
        },

        addToGroup(user) {
            this.toAddToGroup.push(user);
        },

        removeFromGroup(user) {
            this.toAddToGroup = this.toAddToGroup.filter(item => item !== user);
        },

        async search() {
            try {
                let response = await this.$axios.get("/users", {
                    params: {
                        user_id: this.searchUser.trim()
                    }
                });
                this.foundUsers2 = response.data;
                this.searchUser = "";
            } catch (error) {
                console.log(error);
            }
        },

        async createGroup() {
            if (!this.groupName.trim() || this.toAddToGroup.length === 0) {
                alert("Please enter a group name and add at least one user.");
                return;
            }

            try {

                this.toAddToGroup.push(this.identifier)
                let response = await this.$axios.put("/chats", {
                    user_ids: this.toAddToGroup,
                    chat_group: true
                    }
                );

                let chat_id= response.data.chat_id

                response = await this.$axios.put("/chats/"+chat_id, {
                    user_id: this.identifier,
                    first_chat_id: chat_id,
                    chat_name: this.groupName
                    }
                );

                this.toAddToGroup = [];
                this.searchUser = null;
                this.groupName= ""

            }
            catch (error) {
                console.log(error)
            }
            alert("Group created!");
        }
    },

    props: ["identifier"]
};
</script>

<template>
    <div class="container mt-4">
        <div class="d-flex justify-content-between align-items-center mb-3">
            <h2>Create a New Group</h2>
            <button class="btn btn-secondary btn-warning" @click="goHome">Go Back</button>
        </div>

        <div class="mb-3">
            <label for="groupName" class="form-label">Group Name:</label>
            <input type="text" class="form-control" id="groupName" v-model="groupName" placeholder="Enter group name">
        </div>

        <div class="mb-3">
            <label class="form-label">Add Users:</label>
            <div class="d-flex">
                <input class="form-control me-2" type="text" placeholder="@username" v-model="searchUser">
                <button class="btn btn-outline-primary" :disabled="!searchUser || searchUser.trim().length === 0" @click.prevent="search">Search</button>
            </div>
        </div>

        <div v-if="foundUsers2 && foundUsers2.length" class="mt-3">
            <h4>Users Found:</h4>
            <ul class="list-group">
                <li v-for="user in foundUsers2" :key="user.user_id" class="list-group-item d-flex justify-content-between align-items-center">
                    {{ user.user_id }}
                    <button v-if="!toAddToGroup.includes(user.user_id)" class="btn btn-success btn-sm" @click="addToGroup(user.user_id)">Add</button>
                    <button v-else class="btn btn-danger btn-sm" @click="removeFromGroup(user.user_id)">Remove</button>
                </li>
            </ul>
        </div>

        <div v-if="toAddToGroup.length" class="mt-4">
            <h4>Selected Members:</h4>
            <ul class="list-group">
                <li v-for="user in toAddToGroup" :key="user" class="list-group-item">
                    {{ user }}
                </li>
            </ul>
        </div>

        <div class="mt-4">
            <button class="btn btn-primary w-100" :disabled="!groupName.trim() || toAddToGroup.length === 0" @click="createGroup">Create Group</button>
        </div>
    </div>
</template>

<style>
.container {
    max-width: 600px;
    margin: auto;
    background: #f8f9fa;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}
</style>
